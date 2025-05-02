package main

import (
	"fmt"
	"log"
	"sort"
)

func ballsAreSortable(containers [][]int) bool {
	numOfBallsWithColor := make([]int, len(containers))
	numOfBallsInContainer := make([]int, len(containers))
	for i := 0; i < len(containers); i++ {
		numOfBallsInContainer[i] = 0
		numOfBallsWithColor[i] = 0
		for j := 0; j < len(containers); j++ {
			numOfBallsInContainer[i] += containers[i][j]
			numOfBallsWithColor[i] += containers[j][i]
		}
	}

	sort.Ints(numOfBallsWithColor)
	sort.Ints(numOfBallsInContainer)

	for i := 0; i < len(containers); i++ {
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

	fmt.Println()
}
