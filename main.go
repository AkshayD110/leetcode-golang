package main

import (
	"fmt"
	"leetcode_golang/remove_element"
)

func main() {
	fmt.Println("Solving prob 11")

	value := remove_element.RemoveElement([]int{3, 2, 2, 3}, 3)
	fmt.Println(value)
}
