package main

import "fmt"

// greatestProfit from one buy and sell
func greatestProfit(prices []int) int {
	lowest := prices[0]
	greatestProfit := 0

	for _, price := range prices {
		currentProfit := price - lowest

		if price < lowest {
			lowest = price
		} else if currentProfit > greatestProfit {
			greatestProfit = currentProfit
		}
	}

	return greatestProfit
}

func main() {
	prices := []int{10, 7, 5, 8, 11, 2, 6}
	fmt.Println(greatestProfit(prices)) // 6
}
