package main

import (
	"fmt"
	"math/rand"
	"time"
)

// func quick(v []int, lo, hi int) {
// 	if lo > hi {
// 		return
// 	}

// 	pivot := v[hi]
// 	i := lo
// 	for j := lo; j <= hi-1; j++ {
// 		if v[j] < pivot {
// 			swap(v, i, j)
// 			i++
// 		}
// 	}
// 	swap(v, i, hi)

// 	p := i
// 	quick(v, lo, p-1)
// 	quick(v, p+1, hi)
// }

func quick2(v []int, lo, hi int) {
	var stack []int

	stack = append(stack, lo)
	stack = append(stack, hi)

	for {
		h := stack[len(stack)-1]
		stack = append(stack[:len(stack)-1], stack[len(stack):]...)

		l := stack[len(stack)-1]
		stack = append(stack[:len(stack)-1], stack[len(stack):]...)

		p := partition(v, l, h)
		if p-1 > l {
			stack = append(stack, l)
			stack = append(stack, p-1)
		}

		if p+1 < h {
			stack = append(stack, p+1)
			stack = append(stack, h)
		}
	}
}

func partition(v []int, lo, hi int) int {
	pivot := v[hi]
	i := lo
	for j := lo; j <= hi-1; j++ {
		if v[j] < pivot {
			swap(v, i, j)
			i++
		}
	}
	swap(v, i, hi)
	return i
}

func swap(v []int, i, j int) {
	temp := v[i]
	v[i] = v[j]
	v[j] = temp
}

func isSorted(v []int) bool {
	for i := 0; i < len(v)-1; i++ {
		if v[i] > v[i+1] {
			return false
		}
	}

	return true
}

func main() {
	var gs int = 3
	rand.Seed(time.Now().Unix())

	size := gs
	l := make([]int, size)

	for i := 0; i < size; i++ {
		l[i] = rand.Intn(gs * 5)
	}

	quick2(l, 0, len(l)-1)

	fmt.Printf("\n%v sorted: %t\n", l, isSorted(l))
}
