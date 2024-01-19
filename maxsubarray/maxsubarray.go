package maxsubarray

//brute force
func MaxSubArray(nums []int) int {
	currentMax := nums[0]
	for i := 0; i < len(nums); i++ {
		adder := nums[i]
		if adder > currentMax {
			currentMax = adder
		}
		for j := i + 1; j < len(nums); j++ {
			currentSum := adder + nums[j]
			if currentSum > currentMax {
				currentMax = currentSum
			}
			adder = currentSum
		}
	}
	return currentMax

}
