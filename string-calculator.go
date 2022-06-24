package calculator

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

func Add(input string) (int, error) {
	delimiters := regexp.MustCompile("[0-9]")
	matchDelimeters := delimiters.FindAllString(input, -1)

	lastTwoCharacter := strings.LastIndex(input, "\n")
	if lastTwoCharacter == len(input)-1 {
		err := errors.New("the input is not valid")
		return 0, err

	} else {
		var sum int
		for _, i := range matchDelimeters {
			number, _ := strconv.Atoi(i)
			sum += number
		}
		return sum, nil
	}

}
