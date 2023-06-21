package main

import "fmt"

func main() {
	output := search([]int{-1, 0, 3, 5, 9, 12}, 5)
	fmt.Println(output)
}

//brute force
func search(nums []int, target int) int {
	lowIndex := 0
	highIndex := len(nums) - 1
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}
	for lowIndex < highIndex {
		currentIndex := (highIndex + lowIndex) / 2
		if nums[currentIndex] == target {
			return currentIndex
		} else if nums[currentIndex] > target {
			highIndex = currentIndex
		} else if nums[currentIndex] < target {
			lowIndex = currentIndex
		}
	}
	return -1
}

//optimized
func search1(nums []int, target int) int {
	lowIndex := 0
	highIndex := len(nums) - 1

	//note the condition below : lowindex and highindex comparisoin below v/s one in above function
	for lowIndex <= highIndex {
		currentIndex := (highIndex + lowIndex) / 2
		if nums[currentIndex] == target {
			return currentIndex
		} else if nums[currentIndex] > target {
			//note that index is moved by -1 from current vs in above function
			highIndex = currentIndex - 1
		} else if nums[currentIndex] < target {
			//same, movement of index
			lowIndex = currentIndex + 1
		}
	}
	return -1
}
