package main

import (
	"fmt"
)

func main() {
	var n int64
	fmt.Scan(&n)
	var ch1 int64 = 1
	var ch2 int64 = 2
	if n >= ch1 {
		fmt.Println(ch1)
	}
	if n >= ch2 {
		fmt.Println(ch2)
	}
	for ch1+ch2 <= n {
		ch1, ch2 = ch2, ch1
		ch2 = ch1 + ch2
		fmt.Println(ch2)
	}
}
