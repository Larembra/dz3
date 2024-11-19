package main

import (
	"fmt"
	"strconv"
)

var im []int

func scan1() {
	var s string
	im = []int{} //присвоили срезу пустой срез для его очистки
	fmt.Scanln(&s)
	im = append(im, 0)
	for _, i := range s {
		a, _ := strconv.Atoi(string(i))
		im = append(im, a)
	}
	im = append(im, 0)
}
func main() {
	var p [][]int //вводи начальное состояние без пробелов
	var t int64

	scan1()
	nul := make([]int, len(im))
	//nulmin := make([]int, 1)
	f := 0
	for len(im) != 2 {
		if f == 0 {
			p = append(p, nul)
			f++
		}
		p = append(p, im)

		scan1()
	}

	p = append(p, make([]int, len(p[len(p)-1])))
	fmt.Scanln(&t)
	for _, i := range p {
		fmt.Println(i)
	}
	ts := make([][]int, len(p))
	for i := range p {
		ts[i] = make([]int, len(p[i]))
		copy(ts[i], p[i]) // Копируем данные из p в ts, чтобы ts не менялся вместе с p
	}
	for k := int64(0); k < t; k++ {
		for i := int64(1); i < int64(len(p)-1); i++ {
			for j := int64(1); j < int64(len(p[i])-1); j++ {
				sm := p[i-1][j] + p[i-1][j+1] + p[i][j+1] + p[i+1][j+1] + p[i+1][j] + p[i+1][j-1] + p[i][j-1] + p[i-1][j-1]
				if sm == 3 && p[i][j] == 0 {
					ts[i][j] = 1
				} else if sm >= 4 && p[i][j] == 1 {
					ts[i][j] = 0
				} else if sm <= 1 && p[i][j] == 1 {
					ts[i][j] = 0
				}
			}
		}
		fmt.Println(k + 1)
		for ind, i := range ts {
			copy(p[ind], i)
			fmt.Println(i)
		}
	}

}
