package main

import (
	"Calculator/stack"
	"unicode"
	"strconv"
	"fmt"
)

func main() {
	
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
