package calculator

import (
	"strconv"
	"strings"
)

func Add(input string) (int, error) {
	inputWithoutNewLine := strings.Replace(input, "\n", ",", -1)
	inputChars := strings.Split(inputWithoutNewLine, ",")
	var sum int
	for _, i := range inputChars {
		number, _ := strconv.Atoi(i)
		sum += number
	}
	return sum, nil
}
