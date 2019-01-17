package solution

// you can also use imports, for example:
import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int) int {
    n := fmt.Sprintf("%b", N)
    count := 0
    buffer := 0
    for _, v := range (n) {
        //fmt.Println(v == 48)
        //48 is unicode for zero
        if v == 48 {
            buffer ++
        } else if count < buffer {
            count = buffer
            buffer = 0
        }
    }
    return count
}