package main

import "fmt"

func main() {
	input1 := "aabaa"
	input2 := "asdff"
	fmt.Println(checkPalindrome(input1))
	fmt.Println(checkPalindrome(input2))
}

func checkPalindrome(inputString string) bool {
	length := len(inputString)
	if length <= 1 {
		return true
	}

	steps := length / 2
	for i := 0; i < steps; i++ {
		if inputString[i] != inputString[length-1-i] {
			return false
		}
	}
	return true
}
