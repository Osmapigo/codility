package main

import (
	"fmt"
	"sort"
)

func main() {
	A := []int{1, 2, 3}
	fmt.Println(solution(A))
}
func solution(A []int) int {
	sort.Ints(A)
	length := len(A)
	if length == 0 || length == 1 {
		return length + 1
	}
	for i := 0; i < length-1; i++ {
		if A[i+1] != A[i]+1 {
			return A[i] + 1
		}
	}
	return A[length-1]
}
