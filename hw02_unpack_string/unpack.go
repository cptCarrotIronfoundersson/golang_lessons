package hw02unpackstring

import (
	"errors"
	"strconv"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var unpackedRunes []rune
	for index, symbol := range str {
		SymbolSize := len(string(symbol))
		switch {
		case (index == 0 || index == len(str)-SymbolSize) && unicode.IsDigit(symbol):
			return "", ErrInvalidString
		case unicode.IsDigit(symbol):
			if unicode.IsDigit(rune(str[index+SymbolSize])) {
				return "", ErrInvalidString
			}
			continue
		}
		if index == len(str)-SymbolSize {
			unpackedRunes = append(unpackedRunes, symbol)
			break
		}
		nextIndex := rune(str[index+SymbolSize])
		if unicode.IsDigit(nextIndex) {
			if unicode.IsDigit(rune(index)) {
				return "", nil
			}
			nextIndex, _ := strconv.Atoi(string(nextIndex))
			for i := 0; i < nextIndex; i++ {
				unpackedRunes = append(unpackedRunes, symbol)
			}
		} else {
			unpackedRunes = append(unpackedRunes, symbol)
		}
	}
	return string(unpackedRunes), nil
}
