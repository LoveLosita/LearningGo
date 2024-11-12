package lv1

import (
	"fmt"
	"testing"
)

func TestReverseWords(t *testing.T) {
	cases := []struct {
		input    string
		expected string
	}{
		{"the sky is blue", "blue is sky the"},
		{"  hello world  ", "world hello"},
		{"a good   example", "example good a"},
	}
	for _, c := range cases {
		got := ReverseWords(c.input)
		fmt.Println(got)
		if got != c.expected {
			// 如果结果不符预期，记录错误信息并标记测试失败
			t.Errorf("ReverseWords(%s) = %s; want %s", c.input, got, c.expected)
		}
	}
}
