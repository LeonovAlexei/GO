package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func valueSplit(text, actions string) []string {
	return strings.Split(text, actions)
}
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Input:\n")
		var value []string
		var actions string
		text, _ := reader.ReadString('\n')
		var textTS string = strings.TrimSpace(text)
		fmt.Println(text)
		fmt.Println(textTS)
		//fmt.Printf("%q\n", strings.Split(text, "+"))
		switch {
		case strings.Contains(textTS, "+"):
			actions = "+"
			fmt.Println("ЭТО +.")
		case strings.Contains(textTS, "-"):
			fmt.Println("ЭТО -.")
		case strings.Contains(textTS, "*"):
			fmt.Println("ЭТО *")
		case strings.Contains(textTS, "/"):
			fmt.Println("ЭТО /")
		default:
			fmt.Println("Знак не указан")
		}
		value = valueSplit(textTS, actions)

		//fmt.Println(toNumber + 8)
		//var a string = value[0]
		//b := value[1]
		aToNumber, _ := strconv.Atoi(value[0])
		bToNumber, _ := strconv.Atoi(value[1])
		//toNumber, _ := strconv.Atoi(a)
		//fmt.Println(a, b)
		fmt.Println(aToNumber)
		fmt.Println(bToNumber)
		fmt.Println(textTS)
	}
}
