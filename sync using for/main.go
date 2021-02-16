package main

import (
	"fmt"
)

func odd(n int, result chan<- int) {
	for i := 1; i <= n; i++ {
		if i%2 != 0 {
			result <- i
		}
	}
}

func even(n int, result chan<- int) {
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			result <- i
		}
	}
}

func main() {
	c1 := make(chan int, 10)
	c2 := make(chan int, 10)
	go func() {
		for i := 0; true; i++ {
			odd(i, c1)
		}
	}()

	go func() {
		for i := 0; true; i++ {
			even(i, c2)
		}
	}()

	for {
		fmt.Println("Odd", <-c1)
		fmt.Println("Even", <-c2)
	}
}
