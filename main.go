package main

import (
	"fmt"
	"leetcode_golang/maxsubarray"
)

func main() {
	fmt.Println("Solving prob 11")

	value := maxsubarray.MaxSubArray([]int{-1, 0, -2})
	fmt.Println(value)
}
