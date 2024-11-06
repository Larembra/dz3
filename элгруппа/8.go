package main

import (
	"fmt"
	"os"
)

func main() {
	var n1, k1, n2, k2, n3, k3 float64
	fmt.Scan(&n1, &k1)
	fmt.Scan(&n2, &k2)
	fmt.Scan(&n3, &k3)
	if k1 < n1 || k2 < n2 || k3 < n3 {
		fmt.Println("недопустимый отрезок")
		os.Exit(0)
	}
	if k1 >= n1 && k1 >= n2 && k1 >= n3 {
		if k2 >= n1 && k2 >= n2 && k2 >= n3 {
			if k3 >= n1 && k3 >= n2 && k3 >= n3 {
				fmt.Println("есть область пересечения 3-х отрезков")
				os.Exit(0)
			}
		}
	}
	fmt.Println("нет области пересечения 3-х отрезков")
}
