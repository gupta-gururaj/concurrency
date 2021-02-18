package main

import (
	"log"
	"time"
)

func makeRange(min, max int) []int {
	a := make([]int, max-min+1)
	for i := range a {
		a[i] = min + i
	}
	return a
}

func apiCall(i int) {
	log.Println("API call for", i, "started")
	time.Sleep(100 * time.Millisecond)
}

func main() {
	numArray := makeRange(0, 1000)
	start := time.Now()
	done := make(chan bool)
	count1 := 0
	go func() {
		for i := range numArray {
			go apiCall(i)
			count1++
			done <- true
		}
	}()
	for i := 1; i <= 1000; i++ {
		<-done
	}
	elapsed := time.Since(start)
	log.Println("Time taken", elapsed, "and Count", count1)
}
