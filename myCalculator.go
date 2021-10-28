package myCalculator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type calc struct {
}

func (calc) operate(entry []string) int {
	operator1 := parse(entry[0])
	operator2 := parse(entry[1])
	operator := entry[2]
	switch operator {
	case "+":
		return operator1 + operator2
	case "-":

		return operator1 - operator2
	case "*":

		return operator1 * operator2
	case "/":

		return operator1 / operator2
	default:
		fmt.Println(operator, "Not supported")
		return 0
	}

}
func parse(entry string) int {
	operator, _ := strconv.Atoi(entry)
	return operator

}
func readEntry() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()

}
