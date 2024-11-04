package main

import (
	"fmt"
	"strconv"
)

// 将中缀表达式转换为后缀表达式的思路可以简化为以下几个步骤：
// 使用栈：使用一个栈来存放运算符和括号，以确保运算符的优先级和顺序。
// 遍历表达式：从左到右遍历中缀表达式的每个符号（数字、运算符、括号）。
// 处理操作数：如果遇到操作数（如数字），直接将其添加到输出列表中。done
// 处理运算符：如果遇到运算符，比较它的优先级：(1)如果栈为空或栈顶是左括号，将运算符压入栈。done(2)如果遇到优先级高的运算符，压入栈。done
// (3)如果遇到优先级低或相等的运算符，弹出栈顶运算符并添加到输出，直到找到一个优先级低的运算符或左括号为止。done
// 处理括号：(1)如果遇到左括号，压入栈。done (2)如果遇到右括号，弹出栈顶运算符到输出，直到找到左括号。done
// 结束处理：当表达式遍历完成后，弹出栈中所有的运算符并添加到输出。
func Middle_to_Latter(str string) []string {
	stack_p := 0                                           //栈的指针
	output_list := []string{}                              //存放数字和符号，作为输出列表
	temp_num := ""                                         //用于获取整个数字
	stack := make([]string, 1000)                          //存放运算符和括号
	rank := map[string]int{"+": 1, "-": 1, "*": 2, "/": 2} //运算符优先级
	for index, value := range str {
		if value <= '9' && value >= '0' || value == '.' { //入数字
			temp_num += string(value)
			if index == len(str)-1 {
				output_list = append(output_list, temp_num)
				temp_num = ""
				for stack_p > 0 {
					stack_p -= 1
					output_list = append(output_list, stack[stack_p])
				}
			}
		} else if value == '+' || value == '-' || value == '/' || value == '*' || value == ')' { //遇到右括号或者运算符，准备结算
			if temp_num != "" { //不为空才结算
				output_list = append(output_list, temp_num) //直接结算掉数字，将其加入输出列表
				temp_num = ""
			}
			if value == ')' { //是右括号
				for stack_p != 0 {
					stack_p -= 1
					if stack[stack_p] == "(" { //弹出栈顶运算符到输出，直到找到左括号
						break
					}
					output_list = append(output_list, stack[stack_p])
				}
			} else { //是运算符
				if stack_p == 0 || stack[stack_p-1] == "(" { //如果栈为空或栈顶是左括号
					stack[stack_p] = string(value) //入栈
					stack_p += 1
				} else if stack_p != 0 && rank[string(value)] > rank[stack[stack_p-1]] { //如果遇到优先级高的运算符，压入栈。
					stack[stack_p] = string(value) //入栈
					stack_p += 1
				} else if stack_p != 0 && rank[string(value)] <= rank[stack[stack_p-1]] { //如果遇到优先级低或相等的运算符
					for stack_p != 0 {
						stack_p -= 1
						if stack_p >= 0 && rank[string(value)] > rank[stack[stack_p]] { //直到找到一个优先级低的运算符或左括号为止
							stack_p += 1
							break //此时指针停在了该优先级低的运算符之上的可用位置上
						}
						output_list = append(output_list, stack[stack_p]) //弹出栈顶运算符并添加到输出
					}
					stack[stack_p] = string(value) //处理完了，入栈
					stack_p += 1
				}
			}
		} else if value == '(' { //如果遇到左括号，压入栈。
			stack[stack_p] += string(value)
			stack_p += 1
		}
	}
	return output_list
}

func Latter_to_answer(output_list []string) float64 {
	stack := make([]float64, 1000)
	stack_p := 0
	temp := 0.0
	for _, value := range output_list {
		if stack_p >= 2 && (value == "+" || value == "-" || value == "*" || value == "/") { //如果是运算符
			num1 := 0.0
			num2 := 0.0
			num1 = stack[stack_p-2]
			num2 = stack[stack_p-1]
			stack_p -= 2
			if value == "+" {
				temp = num1 + num2
			} else if value == "-" {
				temp = num1 - num2
			} else if value == "*" {
				temp = num1 * num2
			} else if value == "/" {
				if num2 == 0.0 {
					fmt.Println("除数不能为0！后缀表达式错误！")
					return -1.0
				}
				temp = num1 / num2
			} else {
				fmt.Println("输入错误！")
				return -1.0
			}
			stack[stack_p] = temp //算出来后放回栈中
			temp = 0.0
			stack_p += 1
		} else { //如果是数字
			stack[stack_p], _ = strconv.ParseFloat(value, 64)
			stack_p += 1
		}
	}
	if stack_p <= 0 {
		fmt.Println("错误！指针不在它应该在的位置！")
		return -1.0
	}
	result := stack[stack_p-1]
	return result
}
func main() {
	fmt.Println("欢迎使用高级计算机！请输入正确算式：")
	str := ""
	fmt.Scanf("%s", &str)
	out := Middle_to_Latter(str)
	out2 := Latter_to_answer(out)
	fmt.Println(out2)
}
