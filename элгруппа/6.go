package main

import "fmt"

var mem = []int64{1, 1} //чтобы менять её как глобальную

func fib(n int64) int64 {
	if n == 0 || n == 1 {
		return 1
	} else {
		if n-2 >= int64(len(mem)) {
			_ = fib(n - 2) //просто выполнится и я смогу использовать схраненные значения для m
		}
		if n-1 >= int64(len(mem)) {
			_ = fib(n - 1)
		}
		var m = mem[n-2] + mem[n-1]
		if mem[len(mem)-1] < m {
			mem = append(mem, m)
		}
		return m
	}
}
func main() {
	var n int64
	fmt.Scan(&n)
	fmt.Println("рекурсия:")
	_ = fib(n - 1)
	fmt.Println(mem)
	fmt.Println("динамика")
	var mem1 = []int64{1, 1}
	var i int64 = 2
	for i < n {
		mem1 = append(mem1, mem1[len(mem1)-1]+mem1[len(mem1)-2])
		i++
	}
	fmt.Println(mem1)
}
