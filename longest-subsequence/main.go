package main

import "fmt"

// a binary tree like this one cannot find the longest increasing sequence :P
// it might take a balancing tree for that
type Tree struct {
	root *Node
}

func (t *Tree) lookUp(val int) bool {
	return lookUp(t.root, val)
}

func lookUp(n *Node, val int) bool {
	if n == nil {
		return false
	}

	if n.value == val {
		return true
	}

	return lookUp(n.left, val) || lookUp(n.right, val)
}

func (t *Tree) insert(val int) {
	if t.root == nil {
		println("creating new root", val)
		t.root = newNode(val)
		return
	}

	t.root.insert(val)
}

// func deepestBranch(n *Node, depth int) []int {

// }

func maxValue(n *Node, tmp int) int {
	if n == nil {
		return tmp
	}

	maxVal := tmp

	if n.value > maxVal {
		maxVal = n.value
	}

	return maxValue(n.right, maxVal)
}

func deepestBranch(n *Node, depth int) int {
	if n == nil {
		return 0
	}

	fmt.Printf("%d ", n.value)
	dr := len(traverse([]int{}, n.right))
	dl := len(traverse([]int{}, n.left))

	if dr > dl {
		return 1 + deepestBranch(n.right, depth)
	}

	return 1 + deepestBranch(n.left, depth)
}

func traverse(items []int, n *Node) []int {
	if n == nil {
		return items
	}

	values := append(items, n.value)
	values = append(values, traverse(items, n.left)...)
	return append(values, traverse(items, n.right)...)
}

func newNode(val int) *Node {
	return &Node{
		value: val,
	}
}

type Node struct {
	left  *Node
	right *Node
	value int
}

func (n *Node) insert(val int) {
	if n == nil {
		return
	}

	if val > n.value {
		if n.right == nil {
			n.right = newNode(val)
		} else {
			n.right.insert(val)
		}
	} else {
		if n.left == nil {
			n.left = newNode(val)
		} else {
			n.left.insert(val)
		}
	}
}

func main() {
	t := &Tree{}

	vals := []int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15}
	for _, v := range vals {
		t.insert(v)
	}

	fmt.Printf("traverse: %v\n", traverse([]int{}, t.root))
	fmt.Printf("max value from tree: %d\n", maxValue(t.root, 0))
	fmt.Printf("length of deepest tree branch: %d\n", deepestBranch(t.root, 0))
}
