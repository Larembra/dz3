package main

import (
	"fmt"
	"math"
)

func nod(a int64, b int64) int64 {
	for a != 0 && b != 0 {
		var r int64
		if b > a {
			a, b = b, a
		}
		r = a % b
		if r == 0 {
			return b
		}
		a = b
		b = r

	}
	return 1
}
func main() {
	var a, b int64
	fmt.Scan(&a, &b)
	fmt.Println(nod(int64(math.Abs(float64(a))), int64(math.Abs(float64(b)))))
}
