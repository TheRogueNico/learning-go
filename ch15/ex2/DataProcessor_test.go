package main

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync"
	"testing"
)

func TestParser(t *testing.T) {
	testcases := []struct {
		name    string
		data    []byte
		want    Input
		wantErr bool
	}{
		{name: "valid", data: []byte("id1\n+\n3\n4"), want: Input{Id: "id1", Op: "+", Val1: 3, Val2: 4}},
		{name: "too few lines", data: []byte("id2\n+\n3"), wantErr: true},
		{name: "non-numeric val1", data: []byte("id3\n+\nabc\n4"), wantErr: true},
		{name: "non-numeric val2", data: []byte("id4\n+\n3\nabc"), wantErr: true},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := parser(tc.data)
			if (err != nil) != tc.wantErr {
				t.Fatalf("error = %v, wantErr %v", err, tc.wantErr)
			}
			if tc.wantErr {
				return
			}
			if got != tc.want {
				t.Errorf("got %+v, want %+v", got, tc.want)
			}
		})
	}
}

func TestDataProcessor(t *testing.T) {
	testcases := []struct {
		name    string
		data    []byte
		want    Result
		wantOut bool
	}{
		{name: "add", data: []byte("id1\n+\n3\n4"), want: Result{Id: "id1", Value: 7}, wantOut: true},
		{name: "sub", data: []byte("id2\n-\n3\n4"), want: Result{Id: "id2", Value: -1}, wantOut: true},
		{name: "mul", data: []byte("id3\n*\n3\n4"), want: Result{Id: "id3", Value: 12}, wantOut: true},
		{name: "div", data: []byte("id4\n/\n8\n4"), want: Result{Id: "id4", Value: 2}, wantOut: true},
		{name: "division by zero", data: []byte("id5\n/\n12\n0")},
		{name: "unknown operator", data: []byte("id6\n?\n1\n2")},
		{name: "invalid input", data: []byte("Some random data!")},
	}

	in := make(chan []byte, len(testcases))
	out := make(chan Result, len(testcases))
	for _, tc := range testcases {
		in <- tc.data
	}
	close(in)
	DataProcessor(in, out)

	got := make(map[string]Result, len(out))
	for r := range out {
		got[r.Id] = r
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			r, ok := got[tc.want.Id]
			if ok != tc.wantOut {
				t.Fatalf("got result %v, wantOut %v", ok, tc.wantOut)
			}
			if tc.wantOut && r != tc.want {
				t.Errorf("got %+v, want %+v", r, tc.want)
			}
		})
	}
}

func TestWriteData(t *testing.T) {
	in := make(chan Result, 2)
	in <- Result{Id: "a", Value: 1}
	in <- Result{Id: "b", Value: -2}
	close(in)

	var buf bytes.Buffer
	WriteData(in, &buf)

	want := "a:1\nb:-2\n"
	if buf.String() != want {
		t.Errorf("wrote %q, want %q", buf.String(), want)
	}
}

type errReader struct{}

// Always return an error
func (errReader) Read([]byte) (int, error) { return 0, errors.New("read error") }
func (errReader) Close() error             { return nil }

func TestNewController(t *testing.T) {
	testcases := []struct {
		name       string
		outBufSize int
		body       io.Reader
		wantStatus int
	}{
		{name: "success", outBufSize: 1, body: strings.NewReader("id1\n+\n1\n2"), wantStatus: http.StatusAccepted},
		{name: "queue full", outBufSize: 0, body: strings.NewReader("id1\n+\n1\n2"), wantStatus: http.StatusServiceUnavailable},
		{name: "bad input", outBufSize: 1, body: errReader{}, wantStatus: http.StatusBadRequest},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			handler := NewController(make(chan []byte, tc.outBufSize))

			req := httptest.NewRequest(http.MethodPost, "/", tc.body)
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)

			if rec.Code != tc.wantStatus {
				t.Errorf("got status %d, want %d", rec.Code, tc.wantStatus)
			}
		})
	}
}

func TestNewController_Concurrent(t *testing.T) {
	handler := NewController(make(chan []byte, 1000))

	var wg sync.WaitGroup
	for range 50 {
		wg.Go(func() {
			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("id\n+\n1\n2"))
			rec := httptest.NewRecorder()
			handler.ServeHTTP(rec, req)
		})
	}
	wg.Wait()
}
