package calculator

import (
	"strconv"
	"strings"
)

func Add(input string) int {
	inputChars := strings.Split(input, ",")
	var sum int
	for _, i := range inputChars {
		number, _ := strconv.Atoi(i)
		sum += number
	}
	return sum
}
