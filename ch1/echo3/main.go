// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
	"strings" // 字符串处理包
)

// !+
// main 打印所有命令行参数（方式3：使用strings.Join）
func main() {
	// 【Go vs Java】字符串连接
	// Java:  String.join(" ", Arrays.copyOfRange(args, 1, args.length))
	// Go:    strings.Join(os.Args[1:], " ")
	// 注意：这是最高效和简洁的方式，推荐使用
	fmt.Println(strings.Join(os.Args[1:], " "))
}

//!-
