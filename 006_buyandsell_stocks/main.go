package main

import (
	"fmt"
	"sort"
)

func main() {
	output := maxProfit_opti([]int{7, 1, 5, 3, 6, 4})
	fmt.Println(output)
}

func maxProfit(prices []int) int {
	unsorted := make([]int, len(prices))
	copy(unsorted, prices)
	mapIndexPrice := make(map[int]int)
	for i, num := range prices {
		mapIndexPrice[num] = i
	}
	sort.Ints(prices)
	leastPrice := prices[0]
	indexOfLeast := mapIndexPrice[leastPrice]
	maxPrice := 0
	for i := indexOfLeast; i < len(unsorted); i++ {
		if unsorted[i] > leastPrice && unsorted[i] > maxPrice {
			maxPrice = unsorted[i]

		}
	}
	// fmt.Println(maxPrice)
	if maxPrice == 0 {
		return 0
	}
	return maxPrice - leastPrice
}

func maxProfit_opti(prices []int) int {
	maxProfitValue := 0
	leastBuyValue := prices[0]
	for _, val := range prices {
		if val < leastBuyValue {
			leastBuyValue = val
		} else {
			currentProValue := val - leastBuyValue
			if currentProValue > maxProfitValue {
				maxProfitValue = currentProValue
			}
		}
	}
	return maxProfitValue
}
