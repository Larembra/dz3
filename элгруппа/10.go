package main

import (
	"fmt"
)

func main() {
	var god int
	fmt.Scan(&god)
	if (god%4 == 0 && god%100 != 0) || (god%100 == 0 && god%400 == 0) {
		fmt.Println("год високосный")
	} else {
		fmt.Println("год не високосный")
	}
}
