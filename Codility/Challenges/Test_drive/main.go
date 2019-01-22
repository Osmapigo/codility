package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 3, 6, 4, 1, 2}
	//a := []int{-2, -5, -8}
	fmt.Println(solution(a))
}

func solution(input []int) int {
	sort.Ints(input)
	fmt.Println(input)
	answer := 1
	length := len(input)
	if input[length-1] < 0 {
		return answer
	}
	for i := 0; i < len(input); i++ {
		fmt.Println(input[i])
		if input[i] > 0 && input[i]+1 != input[i] {
			fmt.Println("Npe")
			answer = input[i] -1
			break
		} else if input[i] > 0 && i+1 == len(input) {
			answer = input[len(input)-1]
			break
		}
	}
	return answer
}
