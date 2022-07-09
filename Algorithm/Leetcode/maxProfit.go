package Leetcode

import "math"

func maxProfit(prices []int) int {
	//最小点买入，最大点卖出

	minPrice := math.MaxInt

	maxProfit := 0
	for _, price := range prices {
		minPrice = min(price, minPrice)

		maxProfit = max(maxProfit, price-minPrice)
	}
	return maxProfit
}
