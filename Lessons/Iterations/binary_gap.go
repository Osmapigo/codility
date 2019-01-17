package main

// you can also use imports, for example:
import "fmt"

// import "os"

func main() {
	fmt.Println(solution(561892))
}

func solution(N int) int {
	n := fmt.Sprintf("%b", N)
	fmt.Println(n)
	count := 0
	buffer := 0
	for _, v := range n {
		//fmt.Println(v == 48)
		//48 is unicode for zero
		if v == 48 {
			buffer++
		} else if count < buffer {
			count = buffer
			buffer = 0
		} else {
			buffer = 0
		}
	}
	return count
}
