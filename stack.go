package main

import (
	"errors"
	"fmt"
)

type Node struct {
	Value int
}

type Stack struct {
	Items []Node
}

// Push puhses element to top of stack
func (s *Stack) Push(n Node) {
	s.Items = append(s.Items, n)
}

func (s *Stack) Pop() (Node, error) {
	if len(s.Items) == 0 {
		return Node{}, errors.New("empty stack")
	}

	l := len(s.Items)
	item := s.Items[l-1]
	s.Items = s.Items[:l-1]
	return item, nil
}

func main() {
	s := Stack{}
	s.Push(Node{1})
	s.Push(Node{2})
	s.Push(Node{3})

	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
