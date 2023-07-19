package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
)

var romanNumerals map[string]int

func valueSplit(text, actions string) []string {
	text = strings.ReplaceAll(text, " ", "")
	return strings.Split(text, actions)
}
func believe(a, b int, actions string) int {
	switch actions {
	case "+":
		return a + b

	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:

	}
	return 0
}
func comvertAnswerR(num int) string {
	var roman string = ""
	var numbers = []int{1, 4, 5, 9, 10, 40, 50, 90, 100}
	var romans = []string{"I", "IV", "V", "IX", "X", "XL", "L", "XC", "C"}
	var index = len(romans) - 1

	for num > 0 {
		for numbers[index] <= num {
			roman += romans[index]
			num -= numbers[index]
		}
		index -= 1
	}

	return roman
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	romanNumerals = make(map[string]int)
	romanNumerals["I"] = 1
	romanNumerals["II"] = 2
	romanNumerals["III"] = 3
	romanNumerals["IV"] = 4
	romanNumerals["V"] = 5
	romanNumerals["VI"] = 6
	romanNumerals["VII"] = 7
	romanNumerals["VIII"] = 8
	romanNumerals["IX"] = 9
	romanNumerals["X"] = 10

	for {
		fmt.Println("Input:")
		var values []string
		var actions string
		var aar int
		var bar int

		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)

		switch {
		case strings.Contains(text, "+"):
			actions = "+"
		case strings.Contains(text, "-"):
			actions = "-"
		case strings.Contains(text, "*"):
			actions = "*"
		case strings.Contains(text, "/"):
			actions = "/"
		default:
			fmt.Println("Output:\n Вывод ошибки, так как строка не является математической операцией.")
			runtime.Goexit()
		}
		values = valueSplit(text, actions)

		if len(values) != 2 {
			if len(values) < 2 {
				fmt.Println("Output:\n Вывод ошибки, так как строка не является математической операцией.")
				runtime.Goexit()
			} else {
				fmt.Println("Output:\n Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
				runtime.Goexit()
			}
		}
		ar, arok := romanNumerals[values[0]]
		br, brok := romanNumerals[values[1]]
		if arok != brok {
			fmt.Println("Output:\n Вывод ошибки, так как используются одновременно разные системы счисления.")
			runtime.Goexit()
		}
		if !arok && !brok {
			aar, _ = strconv.Atoi(values[0])
			bar, _ = strconv.Atoi(values[1])

			if ((aar < 1) || (aar > 10)) || ((bar < 1) || (bar > 10)) {
				fmt.Println("Output:\n Вывод ошибки, так как формат математической операции не удовлетворяет заданию на вход римские или апабские числа от 1 до 10 включительно, не более.")
				runtime.Goexit()
			}
			answer := believe(aar, bar, actions)
			fmt.Println(answer)
		}
		if arok {
			answer := believe(ar, br, actions)
			if answer < 1 {
				fmt.Println("Output:\n Вывод ошибки, так как в римской системе нет отрицательных чисел.")
			} else {
				fmt.Println(comvertAnswerR(answer))
			}

		}
	}
}
