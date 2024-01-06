package main

import (
	"fmt"
	"leetcode_golang/containingduplicates"
)

func main() {
	fmt.Println("Solving prob 11")
	value := containingduplicates.ContainsDuplicate([]int{1, 2, 3})
	fmt.Println(value)
}
