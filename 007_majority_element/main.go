package main

import "fmt"

func main() {
	output := majorityElement([]int{1})
	fmt.Println(output)
}

func majorityElement(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	halflength := len(nums) / 2
	reverseMap := make(map[int]int)
	for _, num := range nums {

		val, ok := reverseMap[num]
		reverseMap[num] = val + 1
		if ok {
			if val+1 > halflength {
				return num
			}
		}
	}
	return 0
}
