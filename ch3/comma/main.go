// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
//
//	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
//	1
//	12
//	123
//	1,234
//	1,234,567,890
package main

import (
	"fmt"
	"os"
)

// main 为每个命令行参数添加千位分隔符
func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// !+
// comma inserts commas in a non-negative decimal integer string.
// 【Go vs Java】递归函数
//
//	Java:  public static String comma(String s) {
//	         if (s.length() <= 3) return s;
//	         return comma(s.substring(0, s.length()-3)) + "," + s.substring(s.length()-3);
//	       }
//
// Go:    func comma(s string) string { ... }
// 注意：这是一个优雅的递归实现，每次处理最后3位
func comma(s string) string {
	n := len(s)

	// 【Go vs Java】递归终止条件
	// Java:  if (s.length() <= 3) return s;
	// Go:    if n <= 3 { return s }
	if n <= 3 {
		return s
	}

	// 【Go vs Java】字符串拼接和切片
	// Java:  return comma(s.substring(0, n-3)) + "," + s.substring(n-3);
	// Go:    return comma(s[:n-3]) + "," + s[n-3:]
	// 注意：s[:n-3] 前n-3个字符，s[n-3:] 最后3个字符
	//      例如："1234567" -> comma("1234") + "," + "567"
	return comma(s[:n-3]) + "," + s[n-3:]
}

//!-
