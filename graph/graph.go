package main

import (
	"errors"
	"fmt"
	"strings"
)

/*
	Implement a graph traversal with:
	* Depth-first search
	* Breadth-first search
*/

var (
	StackEmptyErr = errors.New("stack is empty")
)

type Node struct {
	Name  string
	Nodes []*Node
}

func NewNode(name string, nodes ...*Node) *Node {
	return &Node{
		Name:  name,
		Nodes: nodes,
	}
}

func (n *Node) IsConnected(node *Node) bool {
	for _, dot := range n.Nodes {
		if dot.Name == node.Name {
			return true
		}
	}

	return false
}

func (n *Node) AddNode(node *Node) {
	if n.IsConnected(node) {
		return
	}

	fmt.Printf("%s <--> %s\n", node.Name, n.Name)

	n.Nodes = append(n.Nodes, node)
	node.AddNode(n)
}

type Stack struct {
	Nodes []*Node
}

func (s *Stack) String() string {
	var stack strings.Builder
	for _, n := range s.Nodes {
		fmt.Fprintf(&stack, "%s ", n.Name)
	}

	return stack.String()
}

func NewStack(nodes ...*Node) *Stack {
	return &Stack{nodes}
}

func (s *Stack) Push(n *Node) {
	s.Nodes = append(s.Nodes, n)
}

func (s *Stack) Pop() (*Node, error) {
	i := len(s.Nodes) - 1
	if i < 0 {
		return nil, StackEmptyErr
	}

	n := s.Nodes[i]
	s.Nodes = append(s.Nodes[:i], s.Nodes[i+1:]...)

	return n, nil
}

func (s *Stack) Contains(node *Node) bool {
	for _, n := range s.Nodes {
		if n.Name == node.Name {
			return true
		}
	}

	return false
}

// because of the order things are added to nodes this is a broad/breadth-first traversal
// breadth first backtrack only a few nodes to continue.
// depth first goes backtracking as it encounters dead-ends,
// and ending only when the algorithm has backtracked past the original "root" vertex
// from the very first step
func BroadthFirstTraversal(node *Node, visited *Stack) {
	fmt.Println("visiting", node.Name)
	if visited.Contains(node) {
		// fmt.Println("already known")
		return
	}
	// fmt.Println("adding to knowns")
	visited.Push(node)

	for _, n := range node.Nodes {
		BroadthFirstTraversal(n, visited)
	}
}

func main() {
	root := NewNode("root")
	a := NewNode("a")
	b := NewNode("b")
	c := NewNode("c")
	d := NewNode("d")

	root.AddNode(a)
	a.AddNode(b)
	b.AddNode(c)
	c.AddNode(d)

	visited := NewStack()

	BroadthFirstTraversal(root, visited)
	fmt.Println(visited)
}
