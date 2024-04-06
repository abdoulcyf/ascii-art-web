package util

import (
	"errors"
	"strings"
)

func StrMaker(cliStr string, patternMap map[byte][]string) string {
	var finalStr strings.Builder

	maxWidth := 0
	for _, char := range patternMap {
		if len(char[0]) > maxWidth {
			maxWidth = len(char[0])
		}
	}

	cliStrArr := strings.Split(strings.TrimSpace(cliStr), "\n\n")

	for row := range patternMap[' '] {
		for _, strItem := range cliStrArr {
			for _, v := range strItem {
				char, ok := patternMap[byte(v)]
				if !ok || row >= len(char) {
					finalStr.WriteString(strings.Repeat(" ", maxWidth))
				} else {
					finalStr.WriteString(char[row])
					if len(char[row]) < maxWidth {
						finalStr.WriteString(strings.Repeat(" ", maxWidth-len(char[row])))
					}
				}
			}
			if strItem != cliStrArr[len(cliStrArr)-1] {
				finalStr.WriteString(" ")
			}
		}
		if row != len(patternMap[' '])-1 {
			finalStr.WriteString("\n")
		}
	}

	return finalStr.String()
}

func ContentToMap(content string, chLength int) (map[byte][]string, error) {
	contentArr := PatternsToArr(content)
	contentMap, err := ToMap(contentArr, chLength)
	return contentMap, err
}

func ToMap(strArr []string, chLength int) (map[byte][]string, error) {
	if len(strArr) < 94*chLength {
		return nil, errors.New("not enough lines")
	}
	resultMap := make(map[byte][]string)
	for i := ' '; i <= '~'; i++ {
		startIndex := int((i - ' ')) * (chLength)
		endIndex := startIndex + chLength
		resultMap[byte(i)] = strArr[startIndex:endIndex]
	}
	return resultMap, nil
}

func PatternsToArr(str string) []string {
	str = strings.Trim(str, "\n")
	str = strings.ReplaceAll(str, "\n\n", "\n")
	chArr := strings.Split(str, "\n")
	return chArr
}
