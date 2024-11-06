package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	var as, op, bs string
	var a, b float64
	var err error

	fmt.Scan(&as, &op, &bs)
	if as == "pi" {
		a = float64(math.Pi)
	} else if as == "e" {
		a = float64(math.E)
	} else {
		a, err = strconv.ParseFloat(as, 64)
		if err != nil {
			fmt.Println("недопустимое число")
			os.Exit(-1)
		}
	}

	if bs == "pi" {
		b = float64(math.Pi)
	} else if bs == "e" {
		b = float64(math.E)
	} else {
		b, err = strconv.ParseFloat(bs, 64)
		if err != nil {
			fmt.Println("недопустимое число")
			os.Exit(-1)
		}
	}

	if op == "+" {
		fmt.Println(a + b)
	} else if op == "-" {
		fmt.Println(a - b)
	} else if op == "*" {
		fmt.Println(a * b)
	} else if op == "/" {
		if b == 0 {
			fmt.Println("на 0 делить нельзя!")
		} else {
			fmt.Println(a / b)
		}
	} else if op == "^" {
		fmt.Println(math.Pow(a, b))
	} else if op == "%" {
		if b == 0 {
			fmt.Println("на 0 делить нельзя!")
		} else {
			if math.Mod(a, b) < 0 {
				if b < 0 {
					fmt.Println(math.Mod(a, b) - b)
				} else {
					fmt.Println(math.Mod(a, b) + b)
				}
			} else {
				fmt.Println(math.Mod(a, b))
			}
		}
	} else {
		fmt.Println("недопустимая операция")
	}
}
