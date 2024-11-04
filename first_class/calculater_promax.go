package main

import (
	"fmt"
	"strconv"
)

//整体思路：中缀表达式转后缀表达式，再利用栈将结果算出

func main() {
	/*
		//第一步，转换成中缀表达式
		output_list := []string{} //存放数字和符号，作为输出列表
		stack := make([]string, 1000)
		temp_num := ""
		stack_p := 0
		str := ""
		flag1 := false
		rank := map[string]int{"+": 1, "-": 1, "*": 2, "/": 2}
		fmt.Println("欢迎使用Go语言计算器ProMax版本！")
		fmt.Println("请输入一个合法的算术式")
		fmt.Scan(&str)
		for _, value := range str {
			if value >= '0' && value <= '9' { //遇到数字，存入临时数字tempnum
				temp_num += string(value)
			} else if value == '.' { //小数点也存入
				temp_num += string(value)
			} else if value == '+' || value == '-' || value == '/' || value == '*' { //遇到运算符，准备结算
				output_list = append(output_list, temp_num) //结算前面的数字，放入输出列表
				temp_num = ""
				//接下来是运算符是否入栈的一系列判断，以及处理
				if stack_p == 0 || stack[stack_p-1] == "(" || rank[stack[stack_p-1]] < rank[string(value)] { //优先级大，直接入栈
					stack[stack_p] = string(value)
					stack_p += 1
				} else { //优先级小，清空一些内容
					for stack[stack_p-1] != "(" && stack_p > 0 { //如果遇到优先级低或相等的运算符，弹出栈顶运算符并添加到输出，
						// 直到找到一个优先级低的运算符或左括号为止。
						if rank[stack[stack_p-1]] < rank[string(value)] {
							flag1 = true
							break
						}
						stack_p -= 1
						output_list = append(output_list, stack[stack_p])
					}
					if flag1 == false {
						stack_p -= 1
						flag1 = true
					}
				}
			} else if value == '(' {
				stack[stack_p] = string(value)
				stack_p += 1
			} else if value == ')' {
				for stack_p >= 0 && stack[stack_p] != "(" {
					stack_p -= 1
					output_list = append(output_list, stack[stack_p])
				}
				if stack[stack_p] == "(" {
					stack_p -= 1
				}
			}
		}
		output_list = append(output_list, temp_num)
		for stack_p > 0 {
			stack_p -= 1
			output_list = append(output_list, stack[stack_p])
		}*/
	//第二步，开始使用后缀表达式计算值
	output_list := []string{"5.1", "3.0", "*", "2.0", "-", "1.0", "+"}
	stack2 := make([]float64, 1000)
	stack2_p := 0
	for _, value := range output_list { //遍历后缀表达式
		if value == "+" || value == "-" || value == "/" || value == "*" { //如果遇到运算符，准备计算栈中元素
			temp_sum := 0.0
			num1 := 0.0
			num2 := 0.0
			num1 = stack2[stack2_p-2]
			num2 = stack2[stack2_p-1]
			if value == "+" {
				temp_sum += num1 + num2
			} else if value == "-" {
				temp_sum += num1 - num2
			} else if value == "/" {
				temp_sum += num1 / num2
			} else if value == "*" {
				temp_sum += num1 * num2
			}
			stack2_p -= 2
			//截止此处都是计算过程，接着需要把算出来的元素放回栈中
			stack2[stack2_p] = temp_sum //放回栈中
		} else { //如果是数字，入栈
			stack2[stack2_p], _ = strconv.ParseFloat(value, 64)
			stack2_p += 1
		}
	}
	//最后算完了，输出
	fmt.Printf("计算结果是:%.2f", stack2[stack2_p])
}
