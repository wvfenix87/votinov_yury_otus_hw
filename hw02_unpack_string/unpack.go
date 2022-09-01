package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

type runeData struct {
	countOfRepeat    int
	value            rune
	isEscaped, isNum bool
}

func Unpack(stringForUnpack string) (string, error) {
	var someError error
	var data []runeData
	var result string
	data, someError = parse(stringForUnpack)
	if someError == nil {
		result = generateString(data)
	}
	return result, someError
}

func parse(stringForUnpack string) ([]runeData, error) {
	var result []runeData
	var parseError error
	var isEscaped bool
	stringRunes := []rune(stringForUnpack)
	for i := 0; i < len(stringRunes); i++ {
		currentRune := stringRunes[i]
		isStartRune := i == 0
		isEndRune := i == len(stringRunes)-1
		isNum := unicode.IsDigit(currentRune)
		if isNum && !isEscaped {
			if isStartRune {
				return nil, ErrInvalidString
			}
			if !isEndRune {
				if unicode.IsDigit(stringRunes[i+1]) && !isEscaped {
					return nil, ErrInvalidString
				}
			}
			var countOfRepeat int
			countOfRepeat, parseError = strconv.Atoi(string(currentRune))
			result[i-1].countOfRepeat = countOfRepeat
		}
		var data runeData
		data.isEscaped = isEscaped
		data.countOfRepeat = 1
		data.isNum = isNum
		data.value = currentRune
		result = append(result, data)
		if currentRune == '\\' && !isEscaped {
			if isEndRune {
				return nil, ErrInvalidString
			}
			isEscaped = true
		} else {
			isEscaped = false
		}
	}
	return result, parseError
}

func generateString(data []runeData) string {
	var unpackedStringBuilder strings.Builder
	for _, currentRuneData := range data {
		if (!currentRuneData.isNum || currentRuneData.isEscaped) &&
			(currentRuneData.value != '\\' || currentRuneData.isEscaped) && currentRuneData.countOfRepeat > 0 {
			unpackedStringBuilder.WriteString(strings.Repeat(string(currentRuneData.value),
				currentRuneData.countOfRepeat))
		}
	}
	return unpackedStringBuilder.String()
}
