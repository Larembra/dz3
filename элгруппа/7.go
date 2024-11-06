package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	var s string //длина 12, 2 символа - не буквы
	var sb strings.Builder
	fmt.Scan(&s)
	s = strings.ToLower(s)
	if unicode.IsLetter(rune(s[0])) {
		sb.WriteRune(rune(s[0]))
	}
	if unicode.IsLetter(rune(s[1])) {
		sb.WriteRune(rune(s[1]))
	}
	if unicode.IsLetter(rune(s[2])) {
		sb.WriteRune(rune(s[2]))
	}
	if unicode.IsLetter(rune(s[3])) {
		sb.WriteRune(rune(s[3]))
	}
	if unicode.IsLetter(rune(s[4])) {
		sb.WriteRune(rune(s[4]))
	}
	if unicode.IsLetter(rune(s[5])) {
		sb.WriteRune(rune(s[5]))
	}
	if unicode.IsLetter(rune(s[6])) {
		sb.WriteRune(rune(s[6]))
	}
	if unicode.IsLetter(rune(s[7])) {
		sb.WriteRune(rune(s[7]))
	}
	if unicode.IsLetter(rune(s[8])) {
		sb.WriteRune(rune(s[8]))
	}
	if unicode.IsLetter(rune(s[9])) {
		sb.WriteRune(rune(s[9]))
	}
	if unicode.IsLetter(rune(s[10])) {
		sb.WriteRune(rune(s[10]))
	}
	if unicode.IsLetter(rune(s[11])) {
		sb.WriteRune(rune(s[11]))
	}
	res := sb.String()
	if res[0] == res[9] {
		if res[1] == res[8] {
			if res[2] == res[7] {
				if res[3] == res[6] {
					if res[4] == res[5] {
						fmt.Println("палиндром")
						os.Exit(0)
					}
				}
			}
		}
	}
	fmt.Println("не палиндром")
}
