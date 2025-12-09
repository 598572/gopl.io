// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 72.

// Basename1 reads file names from stdin and prints the base name of each one.
package main

import (
	"bufio"
	"fmt"
	"os"
)

// main 从标准输入读取文件路径，提取并打印基本文件名
func main() {
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		fmt.Println(basename(input.Text()))
	}
	// NOTE: ignoring potential errors from input.Err()
}

// !+
// basename removes directory components and a .suffix.
// e.g., a => a, a.go => a, a/b/c.go => c, a/b.c.go => b.c
func basename(s string) string {
	// 【Go vs Java】字符串操作（手动实现版本）
	// Java:  String basename = new File(path).getName();
	//        basename = basename.substring(0, basename.lastIndexOf('.'));
	// Go:    手动遍历字符串，使用切片操作

	// Discard last '/' and everything before.
	// 【Go vs Java】倒序遍历字符串
	// Java:  for (int i = s.length() - 1; i >= 0; i--)
	// Go:    for i := len(s) - 1; i >= 0; i--
	// 注意：len(s) 返回字符串的字节数（不是字符数，对于中文要注意）
	for i := len(s) - 1; i >= 0; i-- {
		// 【Go vs Java】字符串索引
		// Java:  if (s.charAt(i) == '/') { ... }
		// Go:    if s[i] == '/' { ... }
		// 注意：s[i] 返回byte类型（uint8），不是rune（字符）
		if s[i] == '/' {
			// 【Go vs Java】字符串切片
			// Java:  s = s.substring(i + 1);
			// Go:    s = s[i+1:]
			// 注意：s[i+1:] 表示从索引i+1到末尾的子串
			s = s[i+1:]
			break
		}
	}

	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			// 【Go vs Java】字符串切片（到指定位置）
			// Java:  s = s.substring(0, i);
			// Go:    s = s[:i]
			// 注意：s[:i] 表示从开头到索引i（不包含i）的子串
			s = s[:i]
			break
		}
	}
	return s
}

//!-
