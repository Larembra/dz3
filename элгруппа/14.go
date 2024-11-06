package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	s, _ := r.ReadString('\n')
	var srev strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		srev.WriteRune(rune(s[i]))
	}
	fmt.Println(srev.String())
}
