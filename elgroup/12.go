package main

import (
	"fmt"
	"math/rand"
	"os"
)

func main() {
	ch := rand.Intn(101)
	var n int
	for i := 0; i < 10; i++ {
		fmt.Scan(&n)
		if n > ch {
			println("больше")
		} else if n < ch {
			println("меньше")
		} else {
			println("ты угадал")
			os.Exit(0)
		}
	}
	println("не угадал")
}
