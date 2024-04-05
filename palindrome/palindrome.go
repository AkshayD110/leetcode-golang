package palindrome

import "fmt"

func IsPalindrome(x int) bool {
	isBool := false
	orgVal := x
	if x < 0 {
		return isBool
	} else {
		output := 0
		for x > 0 {
			quo, rem := x/10, x%10
			output = (output * 10) + rem
			x = quo
		}
		fmt.Println(output)
		if orgVal == output {

			isBool = true
		}
	}
	return isBool
}
