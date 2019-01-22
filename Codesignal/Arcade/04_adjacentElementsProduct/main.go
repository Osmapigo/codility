package main

import (
	"fmt"
	"math"
)

func main() {
	// test1 := []int{4, 1, 2, 3, 1, 5}
	// fmt.Println(adjacentElementsProduct(test1))
	a := fmt.Sprint("as", "df")
	fmt.Println(a)
}

func adjacentElementsProduct(inputArray []int) int {
	max := int(math.Inf(-1))
	for i := 0; i < len(inputArray)-1; i++ {
		if max < inputArray[i]*inputArray[i+1] {
			max = inputArray[i] * inputArray[i+1]
		}
	}
	return max
}
