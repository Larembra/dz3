package main

import (
	"fmt"
	"strconv"
)

func main() {
	var v0_str, a_str, t_str string
	var check = map[bool]func(){true: func() { panic("недопустимое значение") }, false: func() {}}
	fmt.Println("Введите нач скорость, ускорение и время в 1 строку")

	fmt.Scan(&v0_str, &a_str, &t_str)
	v0, err1 := strconv.ParseFloat(v0_str, 64)
	check[err1 != nil]()
	a, err2 := strconv.ParseFloat(a_str, 64)
	check[err2 != nil]()
	t, err3 := strconv.ParseFloat(t_str, 64)
	check[err3 != nil]()
	fmt.Println("Пройденное расстояние:", v0*t+(a*t*t)/2)
}
