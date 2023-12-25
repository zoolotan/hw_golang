package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(s string) (string, error) {
	stringRunes := []rune(s)
	var result strings.Builder

	for i, item := range stringRunes {
		if unicode.IsDigit(item) {
			if i == 0 {
				return "", ErrInvalidString
			}
			if unicode.IsDigit(stringRunes[i+1]) {
				return "", ErrInvalidString
			}
			n, err := strconv.Atoi(string(item))
			if err != nil {
				return "", ErrInvalidString
			}
			if n == 0 {
				trimStr := result.String()[:result.Len()-1]
				result.Reset()
				result.WriteString(trimStr)
				continue
			}
			result.WriteString(strings.Repeat(string(stringRunes[i-1]), n-1))
			continue
		}
		result.WriteString(string(item))
	}
	return result.String(), nil
}
