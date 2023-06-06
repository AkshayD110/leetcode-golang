package main

import (
	"fmt"
	"sort"
)

func main() {
	output := canMakeArithmeticProgression([]int{3, 5, 1})
	fmt.Println(output)
}

func canMakeArithmeticProgression(arr []int) bool {

	// the sort function is available from go 1.8 version. So through built in sort functions in golang. There is also sortslice.
	sort.Ints(arr)
	var diff int
	if len(arr) > 0 {
		diff = arr[1] - arr[0]
	}
	for i := 2; i < len(arr); i++ {
		currDiff := arr[i] - arr[i-1]
		if currDiff != diff {
			return false
		}
	}
	return true
}
