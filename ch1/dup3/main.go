// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 12.

//!+

// Dup3 prints the count and text of lines that
// appear more than once in the named input files.
package main

import (
	"fmt"
	"io/ioutil" // I/O工具包（注意：Go 1.16+推荐用os.ReadFile）
	"os"
	"strings"
)

// main 一次性读取整个文件内容，统计重复行
func main() {
	counts := make(map[string]int)

	// 遍历所有文件参数
	for _, filename := range os.Args[1:] {
		// 【Go vs Java】一次性读取整个文件
		// Java:  String data = Files.readString(Path.of(filename));
		// Go:    data, err := ioutil.ReadFile(filename)
		// 注意：ReadFile返回 []byte（字节切片），不是string
		data, err := ioutil.ReadFile(filename)

		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}

		// 【Go vs Java】字符串分割
		// Java:  String[] lines = data.split("\n");
		// Go:    lines := strings.Split(string(data), "\n")
		// 注意：需要将 []byte 转换为 string
		for _, line := range strings.Split(string(data), "\n") {
			// 【Go vs Java】类型转换
			// Java:  new String(bytes) 或 new String(bytes, charset)
			// Go:    string([]byte) 直接转换
			// 注意：Go的类型转换语法是 Type(value)，不是 new Type(value)
			counts[line]++
		}
	}

	// 打印重复的行
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// 【dup1 vs dup2 vs dup3 对比】
// dup1: 只从标准输入读取，流式处理
// dup2: 可从标准输入或文件读取，流式处理（逐行）
// dup3: 只从文件读取，一次性读取整个文件（适合小文件）

//!-
