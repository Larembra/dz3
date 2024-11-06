package main

import "fmt"

func main() {
	var n int64
	fmt.Scan(&n)
	var s string = ""
	for n != 0 {
		if n%2 == 1 {
			s += "1"
		} else {
			s += "0"
		}
		n /= 2
	}
	for i := int64(len(s) - 1); i >= 0; i-- {
		fmt.Print(string((s[i])))
	}

}
