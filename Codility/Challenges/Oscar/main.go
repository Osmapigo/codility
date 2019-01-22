package main

// you can also use imports, for example:
import "fmt"

func main() {
	fmt.Println(Solution(3, 3))
}

func Solution(A int, B int) string {
	if A >= B {
		return mixString(A, B, "a", "b")
	} else {
		return mixString(B, A, "b", "a")
	}
}

func mixString(A int, B int, first string, last string) string {
	total := A + B
	str := ""
	for k := 1; k <= total; k++ {
		for i := 0; i < 2 && A > 0; i++ {
			str = fmt.Sprint(str, first)
			A--
		}
		for j := 0; j < 2 && B > 0; j++ {
			str = fmt.Sprint(str, last)
			B--
		}

	}

	return str
}
