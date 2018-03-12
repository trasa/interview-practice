package main

import (
	"fmt"
	"testing"
)

func TestGetProductsOfAllIntsExceptAtIndex(t *testing.T) {
	result := getProductsOfAllIntsExceptAtIndex([]int { 1, 7, 3, 4})
	fmt.Printf("%v\n", result)

	result = getProductsOfAllIntsExceptAtIndex([]int { 1, 0, 2 })
	fmt.Printf("%v\n", result)
}

func getProductsOfAllIntsExceptAtIndex(values []int) (result []int) {
	var beforeValues []int
	for i, val := range values {
		if i == 0 {
			beforeValues = append(beforeValues, 1)
		} else {
			beforeValues = append(beforeValues, beforeValues[i-1] * val)
		}
	}
	fmt.Printf("before values %v\n", beforeValues)
	for notPos := range values {
		amt := 1
		for pos, val := range values {
			if pos != notPos {
				amt = amt * val
			}
		}
		result = append(result, amt)
	}
	return
}


func TestHighestProduct(t *testing.T) {
	values := []int { 10, 2, 3, 4, 5}
	fmt.Printf("highest product %d\n", highestProductBruteForce(values))

	values = []int {-10, -10, 1, 3, 2}
	fmt.Printf("highest product %d\n", highestProductBruteForce(values))
}

func highestProductBruteForce(arrayOfInts []int) (best int) {
	for firstpos, first := range arrayOfInts {
		for secondpos, second := range arrayOfInts[firstpos+1:] {
			for _, third := range arrayOfInts[secondpos+2:] {
				fmt.Printf("check %d %d %d\n", first,second, third)
				candidate := first * second * third
				if candidate > best {
					best = candidate
				}
			}
		}
	}
	return
}
