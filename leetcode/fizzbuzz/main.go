package main

import (
	"fmt"
	"strconv"
)

func main() {
	numbers := make(chan int, 30)
	result1 := make(chan string, 100)

	//go worker(numbers, result1)
	go fizz()
	go buzz()
	go fizzbuzz()
	go num()

	for i := 1; i <= 30; i++ {
		numbers <- i
	}
	close(numbers)

	for msg1 := range result1 {
		fmt.Println(msg1)
	}
	fmt.Scanln()
}

func worker(jobs <-chan int, results chan<- string) {
	for n := range jobs {
		results <- fizzbuzz(n)
	}
	close(results)
}

func fizzbuzz(n int) string {
	if n%15 == 0 {
		return "fizzbuzz"
	} else if n%3 == 0 {
		return "fizz"
	} else if n%5 == 0 {
		return "buzz"
	}
	return strconv.Itoa(n)
}
