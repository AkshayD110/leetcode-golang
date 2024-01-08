package main

import (
	"fmt"
	"leetcode_golang/arrayproduct"
)

func main() {
	fmt.Println("Solving prob 11")
	value := arrayproduct.ProductExceptSelf([]int{1, 0, 18})
	fmt.Println(value)
}
