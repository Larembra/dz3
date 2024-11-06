package main

import (
	"fmt"
)

func main() {
	var s string //длина 10
	fmt.Scan(&s)
	var pds string //длина 3
	fmt.Scan(&pds)
	var bb = map[bool]int{false: 0, true: 1}
	i := 2
	ans1 := bb[s[i-2] == pds[0] && s[i-1] == pds[1] && s[i] == pds[2]]
	i += 1
	ans2 := bb[s[i-2] == pds[0] && s[i-1] == pds[1] && s[i] == pds[2]]
	i += 1
	ans3 := bb[s[i-2] == pds[0] && s[i-1] == pds[1] && s[i] == pds[2]]
	i += 1
	ans4 := bb[s[i-2] == pds[0] && s[i-1] == pds[1] && s[i] == pds[2]]
	i += 1
	ans5 := bb[s[i-2] == pds[0] && s[i-1] == pds[1] && s[i] == pds[2]]
	i += 1
	ans6 := bb[s[i-2] == pds[0] && s[i-1] == pds[1] && s[i] == pds[2]]
	i += 1
	ans7 := bb[s[i-2] == pds[0] && s[i-1] == pds[1] && s[i] == pds[2]]
	i += 1
	ans8 := bb[s[i-2] == pds[0] && s[i-1] == pds[1] && s[i] == pds[2]]
	i += 1
	ans9 := bb[s[i-2] == pds[0] && s[i-1] == pds[1] && s[i] == pds[2]]
	i += 1
	fmt.Println(-1 + int(bb[ans1 == 1])*(1) + bb[ans2 == 1]*(2) + bb[ans3 == 1]*(3) + bb[ans4 == 1]*(4) + bb[ans5 == 1]*(5) + bb[ans6 == 1]*(6) + bb[ans7 == 1]*(7) + bb[ans8 == 1]*(8) + bb[ans9 == 1]*(9))
}
