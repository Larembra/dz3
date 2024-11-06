package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	var n int64

	fmt.Scan(&n)
	n = int64(math.Abs(float64(n)))
	var i int64
	for i = 2; i <= int64(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			fmt.Println("не простое")
			os.Exit(0)
		}

	}
	fmt.Println("простое")
}
