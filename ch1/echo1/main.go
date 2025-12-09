// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

// 【Go vs Java】多个import的写法
// Java:  import java.util.*; 或多行 import xxx;
// Go:    使用括号组合多个import，更清晰
import (
	"fmt" // 格式化I/O
	"os"  // 操作系统功能，包括命令行参数
)

// main 打印所有命令行参数（方式1：传统for循环）
func main() {
	// 【Go vs Java】变量声明
	// Java:  String s = "", sep = "";
	// Go:    var s, sep string (声明多个同类型变量)
	// 注意：Go的string零值是""（空字符串），所以这里自动初始化为空
	var s, sep string

	// 【Go vs Java】获取命令行参数
	// Java:  args[0] 是第一个参数
	// Go:    os.Args[0] 是程序名，os.Args[1] 才是第一个参数
	// 注意：os.Args 是 []string 类型（字符串切片）

	// 【Go vs Java】传统for循环
	// Java:  for (int i = 1; i < args.length; i++)
	// Go:    for i := 1; i < len(os.Args); i++
	// 注意：Go的for不需要括号，但{}必须有；len()是内置函数
	for i := 1; i < len(os.Args); i++ {
		// 【Go vs Java】字符串拼接
		// Java:  s += sep + args[i]; (或用StringBuilder)
		// Go:    s += sep + os.Args[i] (简单场景直接用+)
		// 注意：频繁拼接应用 strings.Builder，这里仅示例
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

//!-
