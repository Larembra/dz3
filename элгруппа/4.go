package main

import (
	"fmt"
)

func main() {
	var arr1 [11]float64
	var arr2 [11]float64

	for i := 0; i < 10; i++ {
		fmt.Scan(&arr1[i])
	}
	arr1[10] = arr1[9] - 1

	for i := 0; i < 10; i++ {
		fmt.Scan(&arr2[i])
	}
	arr2[10] = arr2[9] - 1

	i1, i2, i3 := 0, 0, 0
	arr3 := make([]float64, 20)

	for i3 < 20 {
		var b bool
		b = ((!(i1 >= 10) && arr1[i1] <= arr2[i2]) || (!(i1 >= 10) && i2 >= 10))
		var bb = map[bool]float64{false: 0.0, true: 1.0}

		arr3[i3] = arr1[i1]*bb[b] + arr2[i2]*bb[!b]

		i1, i2 = i1+int(bb[b && i1 < 10]), i2+int(bb[!b && i2 < 10])

		i3++
	}

	fmt.Println(arr3)
}
