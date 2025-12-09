// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 72.

// Basename2 reads file names from stdin and prints the base name of each one.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// main 从标准输入读取文件路径，提取并打印基本文件名（使用strings包）
func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}
	// NOTE: ignoring potential errors from input.Err()
}

// basename removes directory components and a trailing .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
// !+
func basename(s string) string {
	// 【Go vs Java】查找子串位置
	// Java:  int slash = s.lastIndexOf("/");
	// Go:    slash := strings.LastIndex(s, "/")
	// 注意：LastIndex返回最后一次出现的位置，未找到返回-1
	slash := strings.LastIndex(s, "/") // -1 if "/" not found

	// 【Go vs Java】切片操作的妙用
	// Java:  s = slash >= 0 ? s.substring(slash + 1) : s;
	// Go:    s = s[slash+1:]
	// 注意：即使slash=-1，s[-1+1:] = s[0:] 也是正确的（从开头到末尾）
	s = s[slash+1:]

	// 【Go vs Java】if语句中的短变量声明
	// Java:  int dot = s.lastIndexOf(".");
	//        if (dot >= 0) { s = s.substring(0, dot); }
	// Go:    if dot := strings.LastIndex(s, "."); dot >= 0 { s = s[:dot] }
	// 注意：dot的作用域仅在if块内，这是Go的惯用法
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

//!-
