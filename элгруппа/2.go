package main

import "fmt"
import "math/cmplx"

func main() {
	var reala, realb, realc float64
	fmt.Scan(&reala, &realb, &realc)
	a := complex(reala, 0) //переводит в complex(128)
	b := complex(realb, 0) //0 значит мнимая часть = 0
	c := complex(realc, 0)
	d := (b * b) - 4*a*c
	if d == 0 {
		fmt.Println((-1 * b) / (2 * a))
	} else {
		fmt.Println(((-1 * b) - (cmplx.Sqrt(d))) / (2 * a))
		fmt.Println(((-1 * b) + (cmplx.Sqrt(d))) / (2 * a))
	}

}
