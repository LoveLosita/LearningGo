package main

import "fmt"

// 确定运算符优先级
func precedence(op byte) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}
	return 0
}

// 进行运算
func applyOperation(a, b int, op byte) int {
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	}
	return 0
}

// 计算的主函数
func calculate(expression string) int {
	// 两个栈数字栈和操作符栈
	numStack := []int{}
	opStack := []byte{}

	for i := 0; i < len(expression); i++ {
		char := expression[i]
		// 如果是数字
		if char >= '0' && char <= '9' {
			// 将多位数拼接起来
			num := int(char - '0') //转换为整数
			// 检查后续字符是否为数字，用于处理多位数
			for i+1 < len(expression) && expression[i+1] >= '0' && expression[i+1] <= '9' {
				i++
				num = num*10 + int(expression[i]-'0')
			}
			// 将数字压入数字栈
			numStack = append(numStack, num)
		} else if char == '+' || char == '-' || char == '*' || char == '/' {
			// 如果是运算符
			// 将栈内优先级高于或等于当前运算符的运算符进行计算
			for len(opStack) > 0 && precedence(opStack[len(opStack)-1]) >= precedence(char) {
				// 从数字栈中取出两个数字
				b := numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]
				a := numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]
				// 从运算符栈中取出一个运算符
				op := opStack[len(opStack)-1]
				opStack = opStack[:len(opStack)-1]

				// 进行运算并将结果压入数字栈
				result := applyOperation(a, b, op)
				numStack = append(numStack, result)
			}
			//直接压入运算符栈
			opStack = append(opStack, char)
		} else if char == '(' {
			// 遇到左括号，直接压入操作符栈
			opStack = append(opStack, char)
		} else if char == ')' {
			// 遇到右括号，依次弹出操作符栈直到遇到左括号
			for len(opStack) > 0 && opStack[len(opStack)-1] != '(' {
				// 弹出栈顶的两个操作数
				b := numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]
				a := numStack[len(numStack)-1]
				numStack = numStack[:len(numStack)-1]
				// 弹出栈顶的操作符
				op := opStack[len(opStack)-1]
				opStack = opStack[:len(opStack)-1]

				// 进行运算并将结果压入数字栈
				result := applyOperation(a, b, op)
				numStack = append(numStack, result)
			}
			// 弹出左括号
			opStack = opStack[:len(opStack)-1]
		}
	}

	// 处理操作符栈中剩余的操作符
	for len(opStack) > 0 {
		// 弹出栈顶的两个操作数
		b := numStack[len(numStack)-1]
		numStack = numStack[:len(numStack)-1]
		a := numStack[len(numStack)-1]
		numStack = numStack[:len(numStack)-1]
		// 弹出栈顶的操作符
		op := opStack[len(opStack)-1]
		opStack = opStack[:len(opStack)-1]

		result := applyOperation(a, b, op)
		numStack = append(numStack, result)
	}

	return numStack[0]
}

func main() {
	var expression string
	fmt.Scanln(&expression)
	fmt.Println(calculate(expression))
}
