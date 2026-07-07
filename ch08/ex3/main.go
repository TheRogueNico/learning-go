package main

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

type List[T comparable] struct {
	Head *Node[T]
	Tail *Node[T]
}

func (l *List[T]) Add(v T) {
	newNode := &Node[T]{
		Value: v,
	}

	// If list is empty
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}

func (l *List[T]) Insert(v T, index int) {
	newNode := &Node[T]{
		Value: v,
	}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}

	if index <= 0 {
		newNode.Next = l.Head
		l.Head = newNode
		return
	}

	curNode := l.Head
	for i := 1; i < index; i++ {
		if curNode.Next == nil {
			curNode.Next = newNode
			l.Tail = curNode.Next
			return
		}
		curNode = curNode.Next
	}

	newNode.Next = curNode.Next
	curNode.Next = newNode
	if l.Tail == curNode {
		l.Tail = newNode
	}
}

func (l *List[T]) Index(v T) int {
	i := 0
	for curNode := l.Head; curNode != nil; curNode = curNode.Next {
		if curNode.Value == v {
			return i
		}
		i++
	}

	return -1
}

func main() {
}
