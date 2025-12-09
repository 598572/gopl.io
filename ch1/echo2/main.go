// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 6.
//!+

// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

// main 打印所有命令行参数（方式2：range遍历）
func main() {
	// 【Go vs Java】短变量声明
	// Java:  String s = "", sep = "";
	// Go:    s, sep := "", ""
	// 注意：:= 只能在函数内使用，自动推断类型，相当于 var s = ""
	s, sep := "", ""

	// 【Go vs Java】切片操作
	// Java:  Arrays.copyOfRange(args, 1, args.length)
	// Go:    os.Args[1:] (切片语法，从索引1到末尾)
	// 注意：[start:end] 左闭右开，省略end表示到末尾

	// 【Go vs Java】range遍历
	// Java:  for (String arg : args) { ... }
	// Go:    for index, value := range slice { ... }
	// 注意：range返回两个值：索引和元素值
	for _, arg := range os.Args[1:] {
		// 【Go vs Java】空白标识符 _
		// Java:  无对应概念，不需要的变量可以不声明
		// Go:    _ 表示丢弃该值（这里丢弃索引，只要元素值）
		// 注意：Go不允许声明未使用的变量，用_可以忽略不需要的返回值
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

//!-
