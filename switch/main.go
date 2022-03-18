package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday, time.Friday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}
	fmt.Println(time.Now())
	t := time.Now()
	time.Sleep(time.Second)
	elapsed := time.Now().Sub(t)
	fmt.Println(elapsed)
	fmt.Println(time.Second)
	location, _ := time.LoadLocation("Asia/Shanghai")
	fmt.Println(time.Now().In(location))
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Println("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI(1.111)
}
