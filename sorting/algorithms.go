package main

import (
	"fmt"
	"math/rand"
	"time"
)

func MakeList(size int) []int {
	rand.Seed(time.Now().Unix())
	var list = make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = rand.Intn(size * size * size)
	}

	return list
}

func swap(i, j int, list []int) {
	ti := list[i]
	list[i] = list[j]
	list[j] = ti
}

func IsSorted(list []int) bool {
	for i := 0; i < len(list)-1; i++ {
		if list[i] > list[i+1] {
			return false
		}
	}

	return true
}

func Selection(list []int) []int {
	for i := 0; i < len(list)-1; i++ {
		var minIndex int = i
		for j := i + 1; j < len(list); j++ {
			if list[j] < list[minIndex] {
				minIndex = j
			}
		}

		if minIndex != i {
			swap(i, minIndex, list)
		}
	}
	return list
}

func Insertion(list []int) []int {
	var i, j int
	i = 1
	for i < len(list) {
		j = i
		for j > 0 && list[j-1] > list[j] {
			swap(j, j-1, list)
			j = j - 1
		}
		i = i + 1
	}
	return list
}

func Shell(list []int) []int {
	var gaps = []int{57, 23, 10, 4, 1}

	var i, j int
	for _, g := range gaps {
		for i = g; i < len(list); i++ {
			j = i
			for j > 0 && list[j-1] > list[j] {
				swap(j, j-1, list)
				j = j - 1
			}
		}
	}
	return list
}

func MergeSort(list []int) []int {
	if len(list) == 1 {
		return list
	}
	mid := len(list) / 2

	var left, right []int

	left = list[:mid]
	right = list[mid:]

	left = MergeSort(left)
	right = MergeSort(right)

	return merge(left, right)
}

func merge(a, b []int) (c []int) {
	for len(a) > 0 && len(b) > 0 {
		if a[0] > b[0] {
			c = append(c, b[0])
			b = b[1:]
		} else {
			c = append(c, a[0])
			a = a[1:]
		}
	}

	for len(a) > 0 {
		c = append(c, a[0])
		a = a[1:]
	}
	for len(b) > 0 {
		c = append(c, b[0])
		b = b[1:]
	}

	return
}

const max int = 100

func main() {
	list := MakeList(max)
	fmt.Printf("unsorted list with length: %d: %+v\n", len(list), list)
	sorted := MergeSort(list)
	fmt.Printf("sorted: %v %+v\n", IsSorted(sorted), sorted)
}
