package main

import "github.com/google/btree"

func main() {
	vals := []btree.Int{0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15}
	t := btree.New(2)

	for _, i := range vals {
		t.ReplaceOrInsert(i)
	}

	println(t.Max().(btree.Int))
	println(t.Len())
}
