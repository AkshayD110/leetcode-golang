package containingduplicates

func ContainsDuplicate(nums []int) bool {
	output_map := make(map[int]int)
	for _, n := range nums {
		_, ok := output_map[n]
		if ok {
			return true
		} else {
			output_map[n] = 1
		}
	}
	return false
}
