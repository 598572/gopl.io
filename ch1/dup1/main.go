// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.
//!+

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
package main

import (
	"bufio"
	"fmt"
	"os"
)

// main 从标准输入读取内容，打印出现次数大于1的行及其次数
// 【Go vs Java】Go的注释风格统一使用 // ，不推荐使用 /** */ 的JavaDoc风格
func main() {
	// 【Go vs Java】创建map的方式
	// Java:  Map<String, Integer> counts = new HashMap<>();
	// Go:    使用 make() 内置函数创建，语法为 make(map[KeyType]ValueType)
	// 注意：Go的类型声明是「变量名在前，类型在后」，与Java相反
	counts := make(map[string]int)

	// 【Go vs Java】创建Scanner读取标准输入
	// Java:  Scanner input = new Scanner(System.in);
	// Go:    使用 bufio.NewScanner 包装 os.Stdin
	// 注意：:= 是「短变量声明」，自动推断类型，相当于 var input = ...
	input := bufio.NewScanner(os.Stdin)

	// 【Go vs Java】循环读取输入
	// Java:  while (input.hasNextLine()) { String line = input.nextLine(); ... }
	// Go:    for 是唯一的循环关键字（没有while），input.Scan() 返回 bool
	// 注意：Go的 for 条件不需要括号，但 {} 是必须的
	for input.Scan() {
		// 【Go vs Java】获取当前行并更新计数
		// Java:  counts.put(line, counts.getOrDefault(line, 0) + 1);
		// Go:    直接用 [] 访问和赋值，未初始化的int默认值为0，可以直接 ++
		counts[input.Text()]++
	}
	// 注意：这里忽略了 input.Err() 的潜在错误，生产代码应检查

	// 【Go vs Java】遍历map
	// Java:  for (Map.Entry<String, Integer> entry : counts.entrySet()) { ... }
	// Go:    使用 range 关键字，可同时获取 key 和 value
	// 注意：map遍历顺序是随机的，与Java的LinkedHashMap不同
	for line, n := range counts {
		if n > 1 {
			// 【Go vs Java】格式化输出
			// Java:  System.out.printf("%d\t%s%n", n, line);
			// Go:    fmt.Printf()，换行用 \n（没有 %n）
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//!-
