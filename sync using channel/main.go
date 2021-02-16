package main

import "fmt"

func odd(n int, done chan bool) {
	msg := <-done
	if n%2 != 0 && !msg {
		fmt.Println("Odd", n)
	}
	done <- true
}

func even(n int, done chan bool) {
	msg := <-done
	if n%2 == 0 && msg {
		fmt.Println("Even", n)
	}
	done <- false
}

func main() {
	done := make(chan bool)
	go func() {
		done <- false
		for i := 1; i <= 10; i++ {
			odd(i, done)
		}
	}()

	go func() {
		for i := 1; i <= 10; i++ {
			even(i, done)
		}
	}()
	fmt.Scanln()
}
