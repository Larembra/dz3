package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	viragenie := strings.Split(input, "")

	stack := list.New()
	polska := ""

	var oper1 = map[string]func() float64{
		")": func() float64 { return 0 },
		"(": func() float64 { return 1 },
		"+": func() float64 { return 2 },
		"-": func() float64 { return 2 },
		"*": func() float64 { return 3 },
		"/": func() float64 { return 3 },
	}
	for i := 0; i < len(viragenie); i++ {
		char := string(viragenie[i]) // Преобразуем текущий символ в строку
		if viragenie[i] >= "0" && viragenie[i] <= "9" {
			polska += char
		} else {
			prior := oper1[char]() // Приоритет текущего оператора
			if prior == 1 {
				stack.PushBack(char)
			} else if char == ")" {
				for stack.Back().Value.(string) != "(" {
					polska += stack.Back().Value.(string) // Добавляем оператор в польскую нотацию
					stack.Remove(stack.Back())
				}
				stack.Remove(stack.Back())
			} else if stack.Len() == 0 || oper1[stack.Back().Value.(string)]() < prior {
				stack.PushBack(char)
			} else {
				for stack.Len() > 0 && oper1[stack.Back().Value.(string)]() >= prior {
					polska += stack.Back().Value.(string) // Добавляем оператор в польскую нотацию
					stack.Remove(stack.Back())            // Удаляем из стека
				}
				stack.PushBack(char) // Добавляем текущий оператор
			}
		}
	}
	fmt.Println(polska)

	//valfloat1, _ := strconv.ParseFloat(viragenie[0], 64)
	//valfloat2, _ := strconv.ParseFloat(viragenie[1], 64)
	//viragenie[len(viragenie)-1] = viragenie[len(viragenie)-1][:1]
	//for i := 2; i < len(viragenie); {
	//	//valfloat2, _ := strconv.ParseFloat(viragenie[i], 64)
	//	var oper = map[string]func() float64{
	//		"+": func() float64 { return valfloat1 + valfloat2 },
	//		"-": func() float64 { return valfloat1 - valfloat2 },
	//		"*": func() float64 { return valfloat1 * valfloat2 },
	//		"/": func() float64 { return valfloat1 / valfloat2 },
	//	}
	//	valfloat1 = oper[viragenie[i]]()
	//	var check = map[bool]func(){true: func() { i++ }, false: func() { fmt.Println(valfloat1); os.Exit(0) }}
	//	check[i < len(viragenie)-1]()
	//	valfloat2, _ = strconv.ParseFloat(viragenie[i], 64)
	//	check[i < len(viragenie)-1]()
	//}
}
