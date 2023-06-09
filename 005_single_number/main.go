package main

import (
	"fmt"
	"sort"
)

func main() {
	output := singleNumber_mod([]int{1, 2, 1, 2, 3})
	fmt.Println(output)
}

func singleNumber(nums []int) int {
	//sort the array first. Below is just a way to sort on an Int array
	//https://pkg.go.dev/sort#example-package
	//https://pkg.go.dev/sort#Slice
	sort.Ints(nums)

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {

			return nums[i]
		}
		i++
	}
	return nums[len(nums)-1]
}

//xor method of solving
func singleNumber_mod(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
	}
	return nums[0]
}
