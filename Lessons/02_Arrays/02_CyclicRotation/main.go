package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println(solution(a, 2))
}

func solution(A []int, K int) []int {
	for i := 0; i < K; i++ {
		A = append(append([]int(nil), A[len(A)-1]), A[0:len(A)-1]...)
	}
	return A
}
