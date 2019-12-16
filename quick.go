package test1

import (
	"fmt"
	"math/rand"
)

func quick(v []int) {
	if n <= 1 {
		return
	}

	// swap(v, 0, rand.Intn(n)%n)
	last := 0

	for i := 1; i < n; i++ {
		if v[i] < v[0] {
			last++
			swap(v, last, i)
		}
	}

	swap(v, 0, last)
	quick(v, last)
	quick(v, n-last-1)
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
	// gs := rand.Intn(30)
	gs := 10

	size := gs
	l := make([]int, size)

	for i := 0; i < size; i++ {
		l[i] = rand.Intn(gs)
	}

	fmt.Printf("s: %v\nis sorted?: %t\n", l, isSorted(l))

	quick(l, len(l))

	fmt.Printf("s: %v\nis sorted?: %t\n", l, isSorted(l))
}
