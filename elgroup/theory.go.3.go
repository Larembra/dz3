package main

import "fmt"

func main() {
	fmt.Println(summ(2, 5))
	fmt.Println(summm(2, 5, 8))
}
func summ(num1, num2 int) int {
	return num1 + num2
}

func summm(num1, num2, num3 int) int {
	return num1 + num2 + num3
}
