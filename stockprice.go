package main

import (
	"fmt"
	"math"
)

func runStockPriceTests() {
	prices := []int {
		 10, 7, 5, 8, 11, 9,
	}
	verifyProfit(getMaxProfitSquared(prices), 6)
	verifyProfit(getMaxProfitOnePass(prices), 6)

	negativeProfit := []int {
		10, 7,
	}
	verifyProfit(getMaxProfitSquared(negativeProfit), -3)
	verifyProfit(getMaxProfitOnePass(negativeProfit), -3)
}

func verifyProfit(profit int, expected int) {
	fmt.Printf("max: %d\n", profit)
	if profit != expected {
		panic("wrong answer!")
	}
}

func getMaxProfitSquared(prices []int) int {
	bestProfit := math.MinInt32

	for buyCandidatePos, buyCandidate := range prices {
		fmt.Printf("buy candidate %d: %d\n", buyCandidatePos, buyCandidate)
		// if we buy at this point...
		for _, sellCandidate := range prices[buyCandidatePos+1:] {
			fmt.Printf("sell candidate: %d, current best %d, profit %d\n", sellCandidate, bestProfit, sellCandidate - buyCandidate)
			if bestProfit < sellCandidate - buyCandidate {
				bestProfit = sellCandidate - buyCandidate
				fmt.Printf("new best! %d\n", bestProfit)
			}
		}
	}
	return bestProfit
}

func getMaxProfitOnePass(prices []int) int {
	if len(prices) < 2 {
		panic("not enough prices")
	}
	bestProfit := prices[1] - prices[0]
	minPrice := prices[0]

	for _, curPrice := range prices[1:]{
		profit := curPrice - minPrice
		if bestProfit < profit {
			bestProfit = profit
		}
		if curPrice < minPrice {
			minPrice = curPrice
		}
	}
	return bestProfit
}
