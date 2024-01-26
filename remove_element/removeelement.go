package remove_element

import "sort"

func RemoveElement(nums []int, val int) int {
	k := len(nums)
	for i, n := range nums {
		if n == val {
			k = k - 1
			nums[i] = 101
		}
	}
	sort.Ints(nums)
	return k
}
