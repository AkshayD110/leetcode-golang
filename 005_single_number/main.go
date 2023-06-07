package main

import (
	"fmt"
	"sort"
)

func main() {
	output := singleNumber([]int{1, 2, 1, 2, 3})
	fmt.Println(output)
}

func singleNumber(nums []int) int {

	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {

			return nums[i]
		}
		i++
	}
	return nums[len(nums)-1]
}
