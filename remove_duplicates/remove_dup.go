package remove_duplicates

import "fmt"

func RemoveDuplicates(nums []int) int {
	current_val := nums[0]
	change_ind := 1
	for _, n := range nums {
		if current_val != n {
			nums[change_ind] = n
			change_ind++
			current_val = n
		}
	}
	fmt.Println(nums)
	return change_ind
}
