package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []int{1, 2, 2, 3, 5, 3, 1, 1, 1}
	fmt.Println(solution(input))
}

func solution(A []int) int {
	sort.Ints(A)

	for i := 0; i < len(A); i++ {
		if i < len(A)-1 && A[i] == A[i+1] {
			i++
		} else {
			return A[i]
		}
	}
	return 1
}
