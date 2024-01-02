package main

import (
	"fmt"
	"leetcode_golang/search_insert_position"
)

func main() {
	fmt.Println("Solving prob 10")
	value := search_insert_position.SearchInsert([]int{1, 3, 5, 6}, 4)
	fmt.Println(value)
}
