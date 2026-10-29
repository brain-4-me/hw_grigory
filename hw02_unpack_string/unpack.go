package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	stringRune := []rune(input)
	newString := strings.Builder{}
	skipNextChar := false

	for index, char := range stringRune {
		if skipNextChar {
			skipNextChar = false
			continue
		}
		switch {
		case unicode.IsDigit(char):
			if index == 0 {
				return "Invalid string", ErrInvalidString
			}

			prev := stringRune[index-1]
			prevWasCount := unicode.IsDigit(prev) && (index < 2 || stringRune[index-2] != '\\')
			if prevWasCount {
				return "", ErrInvalidString
			}

			repeats, err := strconv.Atoi(string(char))
			if err == nil {
				if repeats == 0 {
					cutString, _ := strings.CutSuffix(newString.String(), string(stringRune[index-1]))
					newString.Reset()
					newString.WriteString(cutString)
				} else {
					newString.WriteString(strings.Repeat(string(stringRune[index-1]), repeats-1))
				}
			}
		case char == '\\':
			if index+1 >= len(stringRune) {
				return "", ErrInvalidString
			}
			nextChar := stringRune[index+1]
			if nextChar != '\\' && !unicode.IsDigit(nextChar) {
				return "", ErrInvalidString
			}
			newString.WriteString(string(nextChar))
			skipNextChar = true
			continue
		default:
			newString.WriteString(string(char))
		}
	}

	return newString.String(), nil
}
