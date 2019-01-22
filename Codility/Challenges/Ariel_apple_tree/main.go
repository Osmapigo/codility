package main

import "fmt"

func main() {
	test1 := []int{6, 1, 4, 6, 3, 2, 7, 4}
	fmt.Println(test1)
	apples := maxApples(test1, 3, 2)
	fmt.Println(apples)
}

func maxApples(trees []int, p1 int, p2 int) int {
	ans := 0
	length := len(trees) - p1
	// buffer := make([]int, len(trees))
	// copy(buffer, trees)
	// buffer[0] = 999999
	// fmt.Println(trees, buffer)
	// combinations := []int
	for i := 0; i <= length; i++ {
		sumP1 := sumSlice(trees[i : p1+i])
		sumP2 := getOtherSlices(trees, p1, p2, i)
		if sumP1+sumP2 > ans {
			ans = sumP1 + sumP2
		}
	}
	return ans
}

func sumSlice(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}
	return sum
}

func getOtherSlices(trees []int, p1 int, p2 int, i int) int {
	max := 0
	for j := 0; j < len(trees)-1; j++ {
		inRange := i <= j && j < i+p1
		sum := trees[j] + trees[j+1]
		if !inRange && sum > max {
			max = sum
		}
	}
	return max
}
