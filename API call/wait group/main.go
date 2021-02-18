package main

import (
	"log"
	"sync"
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
	count1 := 0
	s1 := numArray
	s2 := s1[500:1000]
	s1 = s1[:500]
	start := time.Now()
	var wg sync.WaitGroup
	wg.Add(1)
	for i := range s1 {
		go apiCall(i)
		count1++
	}
	wg.Done()
	wg.Add(1)
	for i := range s2 {
		go apiCall(i)
		count1++
	}
	wg.Done()
	wg.Wait()

	elapsed := time.Since(start)
	log.Println("Time taken", elapsed, "and Count", count1)
}
