package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	viragenie := strings.Split(input, " ")

	valfloat1, _ := strconv.ParseFloat(viragenie[0], 64)
	valfloat2, _ := strconv.ParseFloat(viragenie[1], 64)
	viragenie[len(viragenie)-1] = viragenie[len(viragenie)-1][:1]
	for i := 2; i < len(viragenie); {
		//valfloat2, _ := strconv.ParseFloat(viragenie[i], 64)
		var oper = map[string]func() float64{
			"+": func() float64 { return valfloat1 + valfloat2 },
			"-": func() float64 { return valfloat1 - valfloat2 },
			"*": func() float64 { return valfloat1 * valfloat2 },
			"/": func() float64 { return valfloat1 / valfloat2 },
		}
		valfloat1 = oper[viragenie[i]]()
		var check = map[bool]func(){true: func() { i++ }, false: func() { fmt.Println(valfloat1); os.Exit(0) }}
		check[i < len(viragenie)-1]()
		valfloat2, _ = strconv.ParseFloat(viragenie[i], 64)
		check[i < len(viragenie)-1]()
	}
}
