package main

import (
	"testing"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestMergeSortedArray(t *testing.T) {
	first := []int{3, 4, 6, 10, 11, 15}
	second := []int{1, 5, 8, 12, 14, 19}
	result := mergeSortedArray(first, second)
	verifyMergeSortedArray(t, first, second, result)

	first = []int{3, 4, 6}
	second = []int{1}
	result = mergeSortedArray(first, second)
	verifyMergeSortedArray(t, first, second, result)

	result = mergeSortedArray(second, first)
	verifyMergeSortedArray(t, first, second, result)
}

func verifyMergeSortedArray(t *testing.T, first []int, second []int, result []int){
	fmt.Printf("Sorted Array Result: %v\n", result)
	assert.Equal(t, len(first) + len(second), len(result))
	prev :=  -1000
	for _, i := range result {
		assert.True(t, prev <= i)
		prev = i
	}
}

func mergeSortedArray(first []int, second []int) (result []int) {
	firstPtr := 0
	secondPtr := 0
	for ; firstPtr < len(first) && secondPtr < len(second) ; {
		if first[firstPtr] <= second[secondPtr] {
			result = append(result, first[firstPtr])
			firstPtr++
		} else {
			result = append(result, second[secondPtr])
			secondPtr++
		}
	}
	// are there any left overs? (should only be one list)
	if firstPtr < len(first) {
		for i := firstPtr; i < len(first); i++ {
			result = append(result, first[i])
		}
	} else {
		for i := secondPtr; i < len(second); i++ {
			result = append(result, second[i])
		}
	}
	return
}
