package main

import (
	"fmt"
)

const HighestPossibleScore = 100

func topScores() {
	unsorted := []int { 37, 89, 41, 65, 91, 53 , 53, 100, 0 }
	result := topScoreSorter(unsorted)
	for _, val := range result {
		fmt.Printf("%d ", val)
	}
}

func topScoreSorter(scores []int) (result []int) {
	var counter [HighestPossibleScore + 1]int
	for _, s := range scores {
		counter[s] = counter[s] + 1
	}

	for i := HighestPossibleScore; i >= 0; i-- {
		count := counter[i]
		for occurrences := 0; occurrences < count; occurrences++ {
			result = append(result, i)
		}
	}
	return
}
