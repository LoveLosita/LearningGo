package main

import "fmt"

func Roman_to_Int(s string) int {
	sum := 0
	single_sum := 0
	pass_flag := false
	special_flag := false
	for index, value := range s {
		roman_dict := map[rune]int{
			'I': 1,
			'V': 5,
			'X': 10,
			'L': 50,
			'C': 100,
			'D': 500,
			'M': 1000,
		}
		if len(s) == 1 {
			return roman_dict[value]
		}
		if index == len(s)-1 { //最后一位了 直接结算 返回函数值
			if special_flag {
				sum += single_sum
				return sum
			}
			single_sum += roman_dict[value]
			sum += single_sum
			return sum
		}
		if pass_flag {
			pass_flag = false
			continue
		}
		if roman_dict[value] < roman_dict[rune(s[index+1])] { //右边的比自己大，可能要结算了，也可能是例外
			if (value == 'I' && s[index+1] == 'V') || (value == 'I' && s[index+1] == 'X') || //例外：大的在小的右边
				(value == 'X' && s[index+1] == 'L') || (value == 'X' && s[index+1] == 'C') ||
				(value == 'C' && s[index+1] == 'D') || (value == 'C' && s[index+1] == 'M') {
				single_sum += roman_dict[rune(s[index+1])] - roman_dict[value]
				pass_flag = true
				special_flag = true
			} else { //确实是结算
				single_sum += roman_dict[value]
				sum += single_sum
				single_sum = 0
			}
		} else if roman_dict[value] == roman_dict[rune(s[index+1])] && value == 'I' { //也就可能是I了
			single_sum += 1
		} else if roman_dict[value] > roman_dict[rune(s[index+1])] { //右边比自己大
			single_sum += roman_dict[value] //直接加上
		} else {
			fmt.Println("输入异常！")
			return -1
		}

	}
	return -1
}
func main() {
	fmt.Println(Roman_to_Int("III"))     //3
	fmt.Println(Roman_to_Int("MCMXCIV")) //1994
	fmt.Println(Roman_to_Int("LVIII"))   //58
	fmt.Println(Roman_to_Int("IX"))      //9
}
