package main

import (
	"fmt"
)

func main() {
	moveZeroes2([]int{0, 1, 0, 3, 12})

}

//brute force way
func moveZeroes(nums []int) {
	ZeroCount := 0
	var newArray []int
	for _, val := range nums {
		if val != 0 {
			newArray = append(newArray, val)
		} else {
			ZeroCount++
		}

	}
	for i := 0; i < ZeroCount; i++ {
		newArray = append(newArray, 0)
	}
	nums = newArray
	fmt.Println(nums)
}

//o(n) way of solving. Keeping two pointers to move and swap elements
func moveZeroes2(nums []int) {
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
		}
		if nums[j] != 0 {
			j++
		}
	}
	fmt.Println(nums)
}
