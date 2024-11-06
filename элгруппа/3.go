package main

import (
	"fmt"
	"math"
)

func sorter(arr []float64) {
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr); j++ {
			if math.Abs(arr[j-1]) > math.Abs(arr[j]) {
				arr[j-1], arr[j] = arr[j], arr[j-1]
			}
		}
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	var arr []float64 = make([]float64, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	sorter(arr)
	fmt.Println(arr)
}
