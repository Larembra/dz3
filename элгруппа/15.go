package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b int64
	fmt.Scan(&a, &b)
	if b == 0 {
		fmt.Println("на 0 делить нельзя!")
	}
	if b > a {
		a, b = b, a
	}
	if a < 0 {
		a = -1 * a
	}
	if b < 0 {
		b = -1 * b
	}
	r := a % b
	for r != 0 {
		r = a % b
		if r == 0 {
			fmt.Println(b)
			os.Exit(0)
		}
		a = b
		b = r
	}
}
