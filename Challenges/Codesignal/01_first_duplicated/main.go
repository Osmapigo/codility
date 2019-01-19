package main

import "fmt"

func main() {
	a := []int{2, 2}
	fmt.Println(firstDuplicate(a))
}

func firstDuplicate(a []int) int {
	buffer := make(map[int]bool)
	for i := 0; i < len(a); i++ {
		if buffer[a[i]] == false {
			buffer[a[i]] = true
			continue
		}
		return a[i]
	}
	return -1
}
