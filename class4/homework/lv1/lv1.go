package lv1

import (
	"strings"
)

func ReverseWords(s string) string {
	word := strings.Fields(s)
	word2 := ""
	for i := len(word) - 1; i >= 0; i-- {
		if i == 0 {
			word2 += word[i]
			break
		}
		word2 += word[i]
		word2 += " "
	}
	return word2
}
