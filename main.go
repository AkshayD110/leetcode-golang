package main

import (
	"fmt"
	"leetcode_golang/remove_duplicates"
)

func main() {
	fmt.Println("Solving prob 11")
	value := remove_duplicates.RemoveDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4})
	fmt.Println(value)
}
