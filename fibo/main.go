package main

func f(i int) int {
	if i < 2 {
		return 1
	}

	return f(i-1) + f(i-2)
}

func floop(n int) int64 {
	var f0 int64 = 0
	var f1 int64 = 1
	var fib int64 = 0

	for i := 0; i < n; i++ {
		fib = f0 + f1
		f0 = f1
		f1 = fib
	}
	return fib
}

func main() {
	var i int = 1
	for {
		v := floop(i)
		println(i, v)
		if v < 0 {
			break
		}
		i += 1
	}
}
