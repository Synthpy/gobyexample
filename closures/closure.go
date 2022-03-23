package main

import "fmt"

func a() func() {
	i := 1
	return func() {
		fmt.Println(i)
		i++
	}
}

func main() {
	af := a()
	af()
	af()
	af()
	af()
}
