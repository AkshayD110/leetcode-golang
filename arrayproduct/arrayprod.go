package arrayproduct

import "fmt"

func ProductExceptSelfBrut(nums []int) []int {
	outputArray := make([]int, 0)
	for i, _ := range nums {
		prod := 1
		for j, n := range nums {
			if j != i {
				prod = prod * n
			}
		}
		outputArray = append(outputArray, prod)
	}
	return outputArray
}

func ProductExceptSelf(nums []int) []int {
	outputArray := make([]int, 0)
	arrayToRight := make([]int, len(nums))
	arrayToLeft := make([]int, 0)
	offset := 1
	arrayToLeft = append(arrayToLeft, 1)
	arrayToRight[len(nums)-1] = 1
	for i := 1; i < len(nums); i++ {
		value := offset * nums[i-1]
		arrayToLeft = append(arrayToLeft, value)
		offset = value
	}
	offset = 1
	for i := len(nums) - 2; i >= 0; i-- {
		value := offset * nums[i+1]
		arrayToRight[i] = value
		offset = value
	}
	fmt.Println(arrayToLeft)
	fmt.Println(arrayToRight)
	for i, _ := range nums {
		outputArray = append(outputArray, arrayToLeft[i]*arrayToRight[i])
	}
	return outputArray
}
