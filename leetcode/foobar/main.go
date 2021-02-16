package main

import (
	"fmt"
)

func main() {
	c1 := make(chan string, 10)
	c2 := make(chan string, 10)
	go func() {
		for {
			go foo(c1)
		}
	}()
	go func() {
		for {
			go bar(c2)
		}
	}()
	for {
		select {
		case msg1 := <-c1:
			fmt.Print(msg1)
		case msg2 := <-c2:
			fmt.Print(msg2)
		}
		fmt.Print()
	}
	fmt.Scanln()
}

func foo(c1 chan<- string) {
	c1 <- "foo"
}

func bar(c2 chan<- string) {
	c2 <- "bar"
}
