package main

import (
	"fmt"
	"sync"
)

func isValidSlice(slice []int, results chan<- bool, wg *sync.WaitGroup) {
	defer wg.Done()
	exists := make(map[int]bool)
	for _, v := range slice {
		if exists[v] {
			results <- false
		} else {
			exists[v] = true
		}
	}
	results <- true
}

func isValidSudoku(sudoku [][]int, wg *sync.WaitGroup) bool {
	c := make(chan bool, 100)

	//row check
	for x := 0; x < 9; x++ {
		wg.Add(1)
		go isValidSlice(append([]int{}, sudoku[x]...), c, wg)
	}

	//column check
	for y := 0; y < 9; y++ {
		wg.Add(1)
		go isValidSlice(append([]int{}, sudoku[0:9][y]...), c, wg)
	}

	// Check every 3x3 block
	for x := 0; x <= 6; x += 3 {
		for y := 0; y <= 6; y += 3 {
			smallSudoku := append([]int{}, sudoku[x][y:y+3]...)      //row-1
			smallSudoku = append(smallSudoku, sudoku[x+1][y:y+3]...) //row-2
			smallSudoku = append(smallSudoku, sudoku[x+2][y:y+3]...) //row-3
			wg.Add(1)
			go isValidSlice(smallSudoku, c, wg) //smallSudoku = [9 elements]
		}
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for result := range c {
		if !result {
			return false
		}
	}
	return true
}

func main() {
	var wg sync.WaitGroup
	sudoku := [][]int{
		{5, 3, 4, 6, 7, 8, 9, 1, 2},
		{6, 7, 2, 1, 9, 5, 3, 4, 8},
		{1, 9, 8, 3, 4, 2, 5, 6, 7},
		{8, 5, 9, 7, 6, 1, 4, 2, 3},
		{4, 2, 6, 8, 5, 3, 7, 9, 1},
		{7, 1, 3, 9, 2, 4, 8, 5, 6},
		{9, 6, 1, 5, 3, 7, 2, 8, 4},
		{2, 8, 7, 4, 1, 9, 6, 3, 5},
		{3, 4, 5, 2, 8, 6, 1, 7, 9},
	}
	if isValidSudoku(sudoku, &wg) {
		fmt.Println("Valid")
	} else {
		fmt.Println("Invalid")
	}
}
