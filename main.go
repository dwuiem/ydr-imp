package main

import (
	"fmt"
	"log"
	"sort"
)

// ballsAreSortable returns true if there is a solution to sort balls
// using swapping operation in different containers
func ballsAreSortable(containers [][]int) bool {
	numOfBallsWithColor := make([]int, len(containers))
	numOfBallsInContainer := make([]int, len(containers))

	// Calculate sums of rows and columns
	for i := 0; i < len(containers); i++ {
		numOfBallsInContainer[i] = 0
		numOfBallsWithColor[i] = 0
		for j := 0; j < len(containers); j++ {
			numOfBallsInContainer[i] += containers[i][j]
			numOfBallsWithColor[i] += containers[j][i]
		}
	}

	// Sorting slices to find the correct container order
	sort.Ints(numOfBallsWithColor)
	sort.Ints(numOfBallsInContainer)

	// Comparing slices (sums of balls by color and containers)
	for i := 0; i < len(containers); i++ {
		// Check if sums are equal
		if numOfBallsInContainer[i] != numOfBallsWithColor[i] {
			return false
		}
	}
	return true
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil || n < 1 || n > 100 {
		log.Fatal()
	}
	containers := make([][]int, n)
	for i := 0; i < n; i++ {
		containers[i] = make([]int, n)
		for j := 0; j < n; j++ {
			_, err := fmt.Scan(&containers[i][j])
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	if ballsAreSortable(containers) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
