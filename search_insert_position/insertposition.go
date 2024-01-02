package search_insert_position

//ideally the name of package and folder should be same. since folder name sstarts with a number, keeping it different here.

//function name below should be capital, since you want to access it from outside.
func SearchInsert(nums []int, target int) int {
	for i, n := range nums {
		if n == target {
			return i
		} else if target-n < 0 {
			return i
		}
	}
	return len(nums)
}
