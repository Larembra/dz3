package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	d := b*b - 4*a*c
	if a == 0 {
		if c != 0 && b == 0 {
			fmt.Println("нет решений")
		} else if b != 0 && c == 0 {
			fmt.Println(0)
		} else if b != 0 && c != 0 {
			fmt.Println((-1 * b) / c)
		} else {
			fmt.Println("любое число")
		}
	} else {
		if d < 0 {
			fmt.Println("нет решений")
		} else if d == 0 {
			fmt.Println((-1 * b) / (2 * a))
		} else {
			fmt.Println((-1*b + math.Sqrt(d)) / (2 * a))
		}
	}

}
