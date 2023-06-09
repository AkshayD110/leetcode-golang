package main

import (
	"fmt"
	"sort"
)

func main() {
	output := maxProfit([]int{2, 4, 1})
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
