package main

import (
	"fmt"
	"math"
)

var numbers = [...]int{2, 4, 6, 7, 10, 12, 14, 16} //массив

var words = []string{"one", "two", "three"}
var number1 = 2.0
var number2 float32 = 3.0

func main() {

	a := 4
	fmt.Println(a)
	fmt.Println(numbers[2]) //6
	fmt.Println(math.Pow(a, 4))
	//fmt.Println(number1*number2)
}
