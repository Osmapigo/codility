package solution

// you can also use imports, for example:
// import "fmt"
import "math"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(X int, Y int, D int) int {
	var minSteps float64
	minSteps = float64((Y - X))
	return int(math.Ceil(minSteps / float64(D)))
	// write your code in Go 1.4
}
