package lab2

import (
	"fmt"
	"strconv"
	"strings"
)

func PostfixToInfix(input string) (string, error) {
	operators := []string{"+", "-", "*", "/", "^"}
	symbols := strings.Fields(input)

	if input == "" {
		return "", fmt.Errorf("empty input")
	}

	for key, value := range symbols {
		_, err := strconv.Atoi(value)

		if key < 2 && err != nil {
			return "", fmt.Errorf("invalid input")
		}

		if err != nil && !includes(operators, value) {
			return "", fmt.Errorf("invalid input")
		}
	}

	for i := 0; i < len(symbols); i++ {
		if !includes(operators, symbols[i]) {
			continue
		}

		tmp := symbols[i-2] + symbols[i] + symbols[i-1]
		if includes(operators[0:2], symbols[i]) &&
			len(symbols) > 3 {

			tmp = "(" + tmp + ")"
		}

		tmpSli := append([]string{tmp}, symbols[i+1:]...)
		symbols = append(symbols[0:i-2], tmpSli...)
		i = 0
	}

	return symbols[0], nil
}

func includes(sli []string, item string) bool {
	for _, x := range sli {
		if x == item {
			return true
		}
	}

	return false
}
