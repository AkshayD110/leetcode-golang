package main

import "fmt"

func main() {
	output := generate(5)
	fmt.Println(output)
}

//https://leetcode.com/problems/pascals-triangle/
func generate(numRows int) [][]int {
	output := make([][]int, 0)
	// output = append(output, []int{2, 3})

	output = append(output, []int{1})
	for i := 1; i < numRows; i++ {
		currentRow := make([]int, 0)
		currentRow = append(currentRow, 1)
		previousRow := output[i-1]
		for j := 1; j < len(previousRow); j++ {
			currentRow = append(currentRow, previousRow[j]+previousRow[j-1])
			// output = append(output, []int{1})
		}
		currentRow = append(currentRow, 1)
		output = append(output, currentRow)
	}
	return output

}
