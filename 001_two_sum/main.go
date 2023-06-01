package main

import "fmt"

func main() {
	nums := []int{2, 3, 5, 6}
	output := twoSum(nums, 9)
	fmt.Println(output)
}

//brute force method
func twoSum(nums []int, target int) []int {
	//note that since we need the index, we are not using the range kind of for loop. -- "for i, num := range nums{}""
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				//a return is simple and best. I was trying a "break" here earlier.
				return []int{i, j}
			}
		}
	}
	return []int{}
}
