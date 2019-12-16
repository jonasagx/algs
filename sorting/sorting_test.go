package main

import (
	"fmt"
	"testing"
)

var (
	base = 100
)

// selection sort
func BenchmarkSelection(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		list := MakeList(base)
		Selection(list)
	}
}

// insertion sort
func BenchmarkInsertion(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		list := MakeList(base)
		Insertion(list)
	}
}

// shell sort
func BenchmarkShell(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		list := MakeList(base)
		Shell(list)
	}
}

// merge sort
func BenchmarkMerge(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		list := MakeList(base)
		MergeSort(list)
	}
}

type sortingFunc func(*testing.B)

func Test(t *testing.T) {
	var sf = map[string]sortingFunc{
		"merge":     BenchmarkMerge,
		"shell":     BenchmarkShell,
		"insertion": BenchmarkInsertion,
		"selection": BenchmarkSelection,
	}

	sizes := []int{
		10000,
		100000,
		1000000,
	}

	for _, s := range sizes {
		base = s
		fmt.Printf("testing with base %d\n", base)
		for n, f := range sf {
			r := testing.Benchmark(f)
			fmt.Printf("%s-sort took %s to sort %d numbers\n", n, r.T, base)
		}
		println()
	}
}
