// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

// main 可以从标准输入或文件读取，统计重复行
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]

	// 【Go vs Java】条件判断
	// Java:  if (files.length == 0) { ... }
	// Go:    if len(files) == 0 { ... }
	// 注意：Go的if不需要括号，但{}必须有
	if len(files) == 0 {
		// 没有文件参数，从标准输入读取
		// 【Go vs Java】标准输入
		// Java:  System.in
		// Go:    os.Stdin (类型是 *os.File)
		countLines(os.Stdin, counts)
	} else {
		// 遍历每个文件
		for _, arg := range files {
			// 【Go vs Java】打开文件
			// Java:  File f = new File(arg); Scanner scanner = new Scanner(f);
			// Go:    f, err := os.Open(arg)
			// 注意：Go的函数常返回(结果, 错误)两个值，必须处理错误
			f, err := os.Open(arg)

			// 【Go vs Java】错误处理
			// Java:  try-catch 异常机制
			// Go:    显式检查 err != nil
			// 注意：Go没有异常，错误是普通的返回值
			if err != nil {
				// 【Go vs Java】格式化输出到标准错误
				// Java:  System.err.printf("dup2: %s\n", err);
				// Go:    fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				// 注意：%v 是通用格式化动词，可以打印任何类型
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				// 【Go vs Java】continue跳过本次循环
				// Java:  continue; (相同)
				// Go:    continue (相同，但不需要分号)
				continue
			}
			countLines(f, counts)
			// 【Go vs Java】关闭文件
			// Java:  try-with-resources 或 finally { f.close(); }
			// Go:    手动调用 f.Close()，通常用 defer f.Close()
			// 注意：这里应该用 defer，后续会学到
			f.Close()
		}
	}

	// 打印重复的行
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

// countLines 统计文件中每行出现的次数
// 【Go vs Java】函数参数传递
// Java:  void countLines(File f, Map<String, Integer> counts)
// Go:    func countLines(f *os.File, counts map[string]int)
// 注意：*os.File 是指针类型；map是引用类型，修改会影响原map
func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		// 【Go vs Java】map的修改
		// Java:  counts.put(line, counts.getOrDefault(line, 0) + 1);
		// Go:    counts[input.Text()]++ (直接修改，影响外部的map)
		// 注意：Go的map、slice、channel是引用类型，函数内修改会影响外部
		counts[input.Text()]++
	}
	// 注意：这里忽略了 input.Err() 的潜在错误
}

//!-
