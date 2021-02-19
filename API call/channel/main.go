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
	s1 := numArray
	s1 = s1[:500]
	start := time.Now()
	var wg sync.WaitGroup
	func() {
		for i := range s1 {
			wg.Add(1)
			go apiCall(i)
			wg.Done()
			j := 500 + i
			wg.Add(1)
			go apiCall(j)
			wg.Done()
		}
	}()
	wg.Wait()

	elapsed := time.Since(start)
	log.Println("Time taken", elapsed)
}
