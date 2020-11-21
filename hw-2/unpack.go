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
	isEscapedChar := false
	repeat := 0

	for i, r := range runes {
		if !isEscapedChar && string(r) == "\\" {
			isEscapedChar = true
			continue
		}

		if unicode.IsDigit(r) {
			if i == 0 {
				return "", fmt.Errorf("string has to start from character, not a number")
			} else if (i + 1) < len(runes) && !isEscapedChar && unicode.IsDigit(runes[i+1]) {
				return "", fmt.Errorf("string can't contains numbers, you have to use only digits")
			}

			if isEscapedChar {
				char = string(r)
				isEscapedChar = false
			} else {
				digit, err := strconv.Atoi(string(r))

				if err != nil {
					return "", fmt.Errorf("cannot parse letter: %v", r)
				}

				repeat = digit
			}
		} else {
			isEscapedChar = false
			char = string(r)
		}

		// Check if char not repeated:
		// 1. end of string (ex: a3b2c)
		// 2. next char is string (ex: ab3)
		if repeat == 0 && ((i + 1) == len(runes) || !unicode.IsDigit(runes[i+1])) {
			repeat = 1
		}

		if char != "" && repeat > 0 {
			result += strings.Repeat(char, repeat)
			repeat = 0
		}
	}

	return result, nil
}
