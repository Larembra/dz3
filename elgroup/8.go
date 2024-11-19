package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var n, k float64
	fmt.Scan(&n, &k)
	var intn int64 = int64(n / 1)
	if float64(intn) < n {
		intn++
	}
	for i := intn; i <= int64(k); i++ {
		sm := 0
		for _, val := range strconv.Itoa(int(i)) {
			sm += int(math.Pow(float64(val-'0'), float64(len(strconv.Itoa(int(i))))))
		}
		if int64(sm) == i {
			fmt.Println(i)
		}
	}
}
