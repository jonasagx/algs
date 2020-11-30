package main

import "fmt"

func fibo(n int) int {
	if n <= 1 {
		return 1
	}

	return fibo(n-1) + fibo(n-2)
}

func linearFibo(n int) int {
	if n == 0 {
		return 0
	}

	var f, s, c int = 0, 1, 2
	for c <= n {
		temp := s
		s = f + s
		f = temp
		c++
	}

	return f + s
}

func main() {
	for _, i := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20} {
		fmt.Printf("%d = %d\n", fibo(i), linearFibo(i))
	}
	// fmt.Printf("%d = %d\n", fibo(5), linearFibo(5))
}
