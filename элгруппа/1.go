package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s1, s2 int
	var n string
	fmt.Scan(&n, &s1, &s2)
	n10, err := strconv.ParseInt(n, s1, 0)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println(strconv.FormatInt(n10, s2))
}
