// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 74.

// Printints demonstrates the use of bytes.Buffer to format a string.
package main

import (
	"bytes"
	"fmt"
)

// !+
// intsToString is like fmt.Sprint(values) but adds commas.
// 【Go vs Java】高效字符串拼接
// Java:  StringBuilder buf = new StringBuilder();
// Go:    var buf bytes.Buffer
// 注意：频繁拼接字符串应该用bytes.Buffer或strings.Builder，不要用+
func intsToString(values []int) string {
	// 【Go vs Java】Buffer声明
	// Java:  StringBuilder buf = new StringBuilder();
	// Go:    var buf bytes.Buffer
	// 注意：Go的Buffer是值类型，不需要new，零值可用
	var buf bytes.Buffer

	// 【Go vs Java】写入单个字节
	// Java:  buf.append('[');
	// Go:    buf.WriteByte('[')
	// 注意：WriteByte写入一个字节，参数是byte类型
	buf.WriteByte('[')

	// 【Go vs Java】遍历切片
	// Java:  for (int i = 0; i < values.length; i++) { int v = values[i]; ... }
	// Go:    for i, v := range values { ... }
	// 注意：range同时返回索引和值
	for i, v := range values {
		if i > 0 {
			// 【Go vs Java】写入字符串
			// Java:  buf.append(", ");
			// Go:    buf.WriteString(", ")
			buf.WriteString(", ")
		}

		// 【Go vs Java】格式化写入
		// Java:  buf.append(String.format("%d", v));
		// Go:    fmt.Fprintf(&buf, "%d", v)
		// 注意：&buf 取Buffer的地址，因为Fprintf需要io.Writer接口
		//      bytes.Buffer的指针类型实现了io.Writer接口
		fmt.Fprintf(&buf, "%d", v)
	}

	buf.WriteByte(']')

	// 【Go vs Java】转换为字符串
	// Java:  return buf.toString();
	// Go:    return buf.String()
	return buf.String()
}

func main() {
	// 【Go vs Java】切片字面量
	// Java:  int[] arr = {1, 2, 3}; 或 new int[]{1, 2, 3}
	// Go:    []int{1, 2, 3}
	// 注意：[]int{...} 是切片字面量，不需要指定长度
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
}

//!-
