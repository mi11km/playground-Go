package chapter1

import (
	"strings"
)

func CompressStr(str string) string {
	var (
		countConsecutive int
		charConsecutive  rune
	)
	compressed := make([]rune, 0, len(str))
	for _, r := range str {
		if countConsecutive == 0 {
			charConsecutive = r
			countConsecutive++
		} else {
			if charConsecutive != r {
				compressed = append(compressed, charConsecutive, rune(countConsecutive+48))
				countConsecutive = 1
				charConsecutive = r
			} else {
				countConsecutive++
			}
		}
	}
	compressed = append(compressed, charConsecutive, rune(countConsecutive+48))

	if len(compressed) >= len(str)-1 {
		return str
	}
	return string(compressed)
}

func CompressStrWithStringsBuilder(str string) string {
	var (
		countConsecutive int
		charConsecutive  rune
	)
	compressed := &strings.Builder{}
	compressed.Grow(len(str))
	for _, r := range str {
		if countConsecutive == 0 {
			charConsecutive = r
			countConsecutive++
		} else {
			if charConsecutive != r {
				compressed.WriteRune(charConsecutive)
				compressed.WriteRune(rune(countConsecutive + 48))
				countConsecutive = 1
				charConsecutive = r
			} else {
				countConsecutive++
			}
		}
	}
	compressed.WriteRune(charConsecutive)
	compressed.WriteRune(rune(countConsecutive + 48))

	if compressed.Len() >= len(str)-1 {
		return str
	}
	return compressed.String()
}
