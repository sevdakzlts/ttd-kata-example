package calculator

import (
	"strconv"
	"strings"
)

func Add(numbers string) int {
	splitString := strings.Split(numbers, ",")
	var sum int
	for _, i := range splitString {
		num, _ := strconv.Atoi(i)
		sum += num
	}
	return sum
}
