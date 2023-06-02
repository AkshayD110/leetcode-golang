package main

import "fmt"

func main() {
	nums := []int{2, 3, 5, 6}
	output := twoSumOpti(nums, 9)
	fmt.Println(output)
}

//brute force method - O(n2)
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

//optimized - O(n)
func twoSumOpti(nums []int, target int) []int {
	reversemap := make(map[int]int)
	for i, num := range nums {
		diffVal := target - num
		val, ok := reversemap[num]
		if ok {
			return []int{i, val} //note how to return a slice directly without creating a variable.
		}
		// The above two lines can be written in a single line statement link below
		// if val, ok := myMap["foo"]; ok {
		// do something here
		// }

		reversemap[diffVal] = i
	}

	return []int{}
}
