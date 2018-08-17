package main

import (
	"jsq/stack"
	"unicode"
	"strconv"
	"fmt"
)

func main() {
	//print("请输入正确的数学表达式: ")
	//var stat string
	//reader := bufio.NewReader(os.Stdin)
	//stat, _ = reader.ReadString('\n')
	//stat = strings.TrimSpace(stat)
	//postfix := infix2ToPostfix(stat)
	//fmt.Printf("后缀表达式：%s\n", postfix)
	//fmt.Printf("计算结果: %d\n", calculate(postfix))
	
	str := `2018-1994`
	fmt.Println("表达式："+str)
	
	arr := ToArr(str)
	fmt.Println(arr)
	
	re := Change(arr)
	fmt.Println(re)
	
	result := Js(re)
	fmt.Println(result)
	
}

func ToArr(str string) []string{

	stackobject1 := new(stack.ItemStack)
	stackResult := stackobject1.New()
	s := ""
	for i, v := range str {
		if v > 47 && v < 59 || v==46{
			s += string(v)
			if i == len(str) - 1 {
				stackResult.Push(s)
			}
		}else{
			if s != ""{
				stackResult.Push(s)
				s = ""
			}
			stackResult.Push(string(v))
		}
	}
	
	return stackResult.Get()
}

func Change(str []string)[]string{
	//48-57 46
	stackobject1 := new(stack.ItemStack)
	stackResult := stackobject1.New()
	stackobject2 := new(stack.ItemStack)
	stackYs := stackobject2.New()
	
	for _, j := range str{
		 
		s := string(j) 
		switch s {
			case "(" :
				stackYs.Push(s)
			
			case ")":
				for !stackYs.IsEmpty() {
					preChar := stackYs.Top()
					if preChar == "(" {
						stackYs.Pop() // 弹出 "("
						break
					}
					stackResult.Push(preChar) 
					stackYs.Pop()
				}
			
			case "+", "-", "*", "/":
				for !stackYs.IsEmpty(){
	
					if  stackYs.Top() == "(" || isLower(stackYs.Top(), s){
						break
					}
					stackResult.Push(stackYs.Pop())
				}
				stackYs.Push(s)
	
			default:
				stackResult.Push(s)
			}
		 
	}
	
	
	
	
	for !stackYs.IsEmpty(){
		stackResult.Push(stackYs.Pop())
	}
	return stackResult.Get()	
}

func Js(arr []string)[]string{
	stackobject1 := new(stack.ItemStack)
	stackResult := stackobject1.New()
	for _, v := range arr{ 
		switch v {
		case "+", "-", "*", "/":
			if stackResult.IsEmpty(){
				fmt.Println("error")
			} 
			first, err := strconv.ParseFloat(stackResult.Pop(),32)
			CheckErr(err)
			second, err := strconv.ParseFloat(stackResult.Pop(),32)
			CheckErr(err)
			switch v {
			case "+":
				reStr := second + first				
				s1 := strconv.FormatFloat(reStr, 'f', -1, 32)
				stackResult.Push(s1)
			case "-":
				reStr := second - first 
				s1 := strconv.FormatFloat(reStr, 'f', -1, 32)
				stackResult.Push(s1)
			case "*":
				reStr := second * first
				s1 := strconv.FormatFloat(reStr, 'f', -1, 32)
				stackResult.Push(s1)
			case "/":
				reStr := second / first
				s1 := strconv.FormatFloat(reStr, 'f', -1, 32)
				stackResult.Push(s1)
			}			
		default:
			stackResult.Push(v)
		}		
	} 
	return stackResult.Get()	
}

func CheckErr(err error){
	if err != nil{
		fmt.Println(err)
	}
}

func calculate(postfix string) int {
	stack := stack.ItemStack{}
	fixLen := len(postfix)
	for i := 0; i < fixLen; i++ {
		nextChar := string(postfix[i])
		// 数字：直接压栈
		if unicode.IsDigit(rune(postfix[i])) {
			stack.Push(nextChar)
		} else {
			// 操作符：取出两个数字计算值，再将结果压栈
			num1, _ := strconv.Atoi(stack.Pop())
			num2, _ := strconv.Atoi(stack.Pop())
			switch nextChar {
			case "+":
				stack.Push(strconv.Itoa(num1 + num2))
			case "-":
				stack.Push(strconv.Itoa(num1 - num2))
			case "*":
				stack.Push(strconv.Itoa(num1 * num2))
			case "/":
				stack.Push(strconv.Itoa(num1 / num2))
			}
		}
	}
	result, _ := strconv.Atoi(stack.Top())
	return result
}

func infix2ToPostfix(exp string) string {
	stack := stack.ItemStack{}
	postfix := ""
	expLen := len(exp)

	// 遍历整个表达式
	for i := 0; i < expLen; i++ {

		char := string(exp[i])

		switch char {
		case " ":
			continue
		case "(":
			// 左括号直接入栈
			stack.Push("(")
		case ")":
			// 右括号则弹出元素直到遇到左括号
			for !stack.IsEmpty() {
				preChar := stack.Top()
				if preChar == "(" {
					stack.Pop() // 弹出 "("
					break
				}
				postfix += preChar
				stack.Pop()
			}

			// 数字则直接输出
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			j := i
			digit := ""
			for ; j < expLen && unicode.IsDigit(rune(exp[j])); j++ {
				digit += string(exp[j])
			}
			postfix += digit
			i = j - 1 // i 向前跨越一个整数，由于执行了一步多余的 j++，需要减 1

		default:
			// 操作符：遇到高优先级的运算符，不断弹出，直到遇见更低优先级运算符
			for !stack.IsEmpty() {
				top := stack.Top()
				if top == "(" || isLower(top, char) {
					break
				}
				postfix += top
				stack.Pop()
			}
			// 低优先级的运算符入栈
			stack.Push(char)
		}
	}

	// 栈不空则全部输出
	for !stack.IsEmpty() {
		postfix += stack.Pop()
	}

	return postfix
}

func isLower(top string, newTop string) bool {
	// 注意 a + b + c 的后缀表达式是 ab + c +，不是 abc + +
	switch top {
	case "+", "-":
		if newTop == "*" || newTop == "/" {
			return true
		}
	case "(":
		return true
	}
	return false
}
