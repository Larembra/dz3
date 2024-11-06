package main

import (
	"fmt"
	"math"
)

func main() {
	var n, k float64
	fmt.Scan(&n, &k)
	f := 0
	for i := int64(n / 1); i <= int64(k); i++ {
		if float64(i) > n {
			f = 0
			for j := int64(2); j <= int64(math.Sqrt(float64(i))); j++ {
				if i%j == 0 {
					f = 1
					break
				}
			}
			if f == 0 {
				fmt.Println(i)
			}
		}
	}
}
