package main

import (
	"io"
	"os"
	"sort"
)

type Ranker interface {
	Ranking() []string
}

func RankPrinter(r Ranker, w io.Writer) {
	result := r.Ranking()
	for _, v := range result {
		io.WriteString(w, v)
		w.Write([]byte("\n"))
	}
}

type Team struct {
	Name    string
	Players []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

func (l *League) MatchResult(team1 string, score1 int, team2 string, score2 int) {
	if score1 > score2 {
		l.Wins[team1]++
	} else if score1 < score2 {
		l.Wins[team2]++
	} else {
		// It's a draw, no wins
		return
	}
}

func (l League) Ranking() []string {
	teams := make([]string, len(l.Teams))
	for i, team := range l.Teams {
		teams[i] = team.Name
	}

	// sort bassed on most wins
	sort.Slice(teams, func(i, j int) bool {
		return l.Wins[teams[i]] > l.Wins[teams[j]]
	})

	return teams
}

func main() {
	myTeams := []Team{
		{Name: "Red", Players: []string{"Rose", "Remi", "Ranch"}},
		{Name: "Green", Players: []string{"Gabe", "Garry"}},
		{Name: "Blue", Players: []string{"Buck", "Bill", "Bush"}},
	}

	league := League{
		Teams: myTeams,
		Wins:  map[string]int{},
	}

	league.MatchResult("Red", 100, "Blue", 10)
	league.MatchResult("Red", 100, "Blue", 900)
	league.MatchResult("Green", 200, "Blue", 50)
	league.MatchResult("Green", 1000, "Red", 80)

	RankPrinter(league, os.Stdout)
}
