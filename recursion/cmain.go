package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func cfib() func(int) int {
	counter := 0
	var fib func(int) int
	fib = func(a int) int {
		fmt.Println("counter: n", counter, a)
		counter++
		if a < 2 {
			return a
		}
		return fib(a-2) + fib(a-1)
	}
	return fib
}

func main() {
	fib := cfib()
	fmt.Println("result:", fib(30))
}
