package main

import "fmt"

func main() {
	var n, k int64
	fmt.Scan(&n, &k)
	for i := n; i <= k; i++ {
		fmt.Println(i)
		for j := 1; j <= 10; j++ {
			fmt.Println(i, "*", j, "=", i*int64(j))
		}
	}
}
