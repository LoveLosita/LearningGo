package main

import "fmt"

func RomanToInt(s string) int {
	// 定义罗马数字对应的整数值映射
	romanDict := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	sum := 0
	for i := 0; i < len(s); i++ {
		currVal := romanDict[rune(s[i])]

		// 如果当前数字小于下一个数字，则减去当前值，否则加上当前值
		if i < len(s)-1 && currVal < romanDict[rune(s[i+1])] {
			sum -= currVal
		} else {
			sum += currVal
		}
	}
	return sum
}

func main() {
	fmt.Println(RomanToInt("III"))     // 3
	fmt.Println(RomanToInt("MCMXCIV")) // 1994
	fmt.Println(RomanToInt("LVIII"))   // 58
	fmt.Println(RomanToInt("IX"))      // 9
}
