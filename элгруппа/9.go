package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	//длина 14, слов 3
	reader := bufio.NewReader(os.Stdin)
	s, _ := reader.ReadString('\n') //не fmt.Scan() для чтения всего предложения, а не одного слова
	var s1, s2, s3 strings.Builder
	f := 0
	for i := 0; i < 14; i++ {
		//fmt.Println(i)
		switch f {
		case 0:
			if !unicode.IsLetter(rune(s[i])) && len(s1.String()) == 0 { //не string(s1)!

			} else if unicode.IsLetter(rune(s[i])) {
				s1.WriteRune(rune(s[i]))
			} else if !unicode.IsLetter(rune(s[i])) && len(s1.String()) != 0 {
				f++
			}

		case 1:
			if !unicode.IsLetter(rune(s[i])) && len(s2.String()) == 0 {

			} else if unicode.IsLetter(rune(s[i])) {
				s2.WriteRune(rune(s[i]))
			} else if !unicode.IsLetter(rune(s[i])) && len(s2.String()) != 0 {
				f++
			}

		case 2:
			if !unicode.IsLetter(rune(s[i])) && len(s3.String()) == 0 {

			} else if unicode.IsLetter(rune(s[i])) {
				s3.WriteRune(rune(s[i]))
			} else if !unicode.IsLetter(rune(s[i])) && len(s3.String()) != 0 {
				f++
			}

		default: //когда f == 3
			if len(s1.String()) >= len(s2.String()) && len(s1.String()) >= len(s3.String()) {
				fmt.Println(s1.String())
				f++
			} else if len(s2.String()) >= len(s1.String()) && len(s2.String()) >= len(s3.String()) {
				fmt.Println(s2.String())
				f++
			} else {
				fmt.Println(s3.String())
				f++
			}

		}

	}
	if f != 4 {
		if len(s1.String()) >= len(s2.String()) && len(s1.String()) >= len(s3.String()) {
			fmt.Println(s1.String())
			f++
		} else if len(s2.String()) >= len(s1.String()) && len(s2.String()) >= len(s3.String()) {
			fmt.Println(s2.String())
			f++
		} else {
			fmt.Println(s3.String())
			f++
		}
	}

}
