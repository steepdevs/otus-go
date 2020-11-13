package hw_2

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Unpack(s string) (string, error) {
	if len(s) == 0 {
		return "", nil
	}

	result, char := "", ""
	runes := []rune(s)

	for i, r := range runes {
		if unicode.IsDigit(r) {
			if i == 0 {
				return "", fmt.Errorf("string has to start from character, not a number")
			} else if unicode.IsDigit(runes[i-1]) {
				return "", fmt.Errorf("string can't contains numbers, you have to use only digits")
			}

			repeat, err := strconv.Atoi(string(r))

			if err != nil {
				return "", fmt.Errorf("cannot parse letter: %v", r)
			}

			result += strings.Repeat(char, repeat)
		 } else {
		 	char = string(r)

		 	if (i + 1) == len(runes) || !unicode.IsDigit(runes[i+1]) {
		 		result += char
			}
		}
	}

	return result, nil
}
