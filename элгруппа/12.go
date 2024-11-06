package main

import (
	"fmt"
)

func main() {
	var n, k float64
	fmt.Scan(&n, &k)
	f := 0
	for i := int(n / 1); i <= int(k); i++ {
		if float64(i) > n {
			f = 0
			for j := 2; j <= i/2; j++ {
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
