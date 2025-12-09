// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 33.
//!+

// Echo4 prints its command-line arguments.
package main

import (
	"flag" // 命令行标志解析包
	"fmt"
	"strings"
)

// 【Go vs Java】命令行参数解析
// Java:  使用第三方库如 Apache Commons CLI 或 JCommander
// Go:    内置 flag 包，非常简洁
// 注意：flag.Bool/String等返回指针，需要用*解引用

// 【Go vs Java】包级变量初始化
// Java:  private static boolean n = parseFlag(...);
// Go:    var n = flag.Bool("n", false, "omit trailing newline")
// 注意：flag.Bool返回 *bool 指针，默认值false，描述用于-h帮助
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

// main 打印命令行参数，支持自定义分隔符和换行控制
// 使用示例：go run main.go -s "," -n arg1 arg2 arg3
func main() {
	// 【Go vs Java】解析命令行标志
	// Java:  CommandLine cmd = parser.parse(options, args);
	// Go:    flag.Parse()
	// 注意：Parse()之后，flag.Args()返回非标志参数
	flag.Parse()

	// 【Go vs Java】指针解引用
	// Java:  无指针概念
	// Go:    *sep 解引用指针获取值
	// 注意：sep是*string类型，*sep是string类型
	fmt.Print(strings.Join(flag.Args(), *sep))

	// 【Go vs Java】布尔取反
	// Java:  if (!n) { ... }
	// Go:    if !*n { ... }
	// 注意：!*n 表示"对指针解引用后取反"
	if !*n {
		fmt.Println()
	}
}

//!-
