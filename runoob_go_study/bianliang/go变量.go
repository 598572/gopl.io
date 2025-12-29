package main

import "fmt"

func main() {
	var a string = "Runoob" // 声明变量并初始化
	fmt.Println(a)          // 输出变量值

	var b, c int = 1, 2 // 声明多个变量并初始化
	fmt.Println(b, c)   // 输出变量值

	var d = true   // 声明变量并不指定类型，系统会根据初始值自动推断类型
	fmt.Println(d) // 输出变量值

	//	Go 语言：= 运算符可以简化变量声明和初始化
	e := "Hello, World!" // 声明变量并初始化
	fmt.Println(e)       // 输出变量值
}
