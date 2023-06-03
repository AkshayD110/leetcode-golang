package main

import "fmt"

func main() {
	output := reverse(-120)
	fmt.Println(output)
}

//unoptimized first solution
func reverse(x int) int {
	var negative bool
	input := x
	if x < 0 {
		negative = true
		input = x * -1
	}
	output := 0
	for input >= 1 {
		quo, rem := input/10, input%10
		//below formula does the trick - "reverseNum=(reverseNum*10) + rem"
		output = (output * 10) + rem
		input = quo
	}
	if negative {
		output = output * -1
	}
	return output
}
