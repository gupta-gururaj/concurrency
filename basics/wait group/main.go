package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	go func() {
		wg.Add(1)
		count("sheep", &wg)
	}()

	go func() {
		wg.Add(1)
		count("cow", &wg)
	}()

	wg.Wait()
}

func count(thing string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
	}
}
