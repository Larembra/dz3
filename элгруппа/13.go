package main

import (
	"fmt"
	"math"
)

func main() {
	var n int64
	fmt.Scan(&n)
	r := make([][]string, n)
	for i := range r {
		r[i] = make([]string, n)
	}
	for i := int64(0); i < n; i++ {
		for j := int64(0); j < n; j++ {
			if j >= int64(math.Abs(float64((n/2)-i))) && j <= n-1-int64(math.Abs(float64((n/2)-i))) {
				r[i][j] = "*"
			} else {
				r[i][j] = "."
			}
		}
	}
	for _, i := range r {
		fmt.Println(i)
	}
}
