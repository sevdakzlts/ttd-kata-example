package calculator

import (
	"errors"
	"strconv"
	"strings"
)

func Add(input string) (int, error) {
	lastTwoCharacter := strings.LastIndex(input, "\n")
	if lastTwoCharacter == len(input)-1 {
		err := errors.New("the input is not valid")
		return 0, err

	} else {

		inputWithoutNewLine := strings.Replace(input, "\n", ",", -1)
		inputChars := strings.Split(inputWithoutNewLine, ",")
		var sum int
		for _, i := range inputChars {
			number, _ := strconv.Atoi(i)
			sum += number
		}
		return sum, nil
	}

}
