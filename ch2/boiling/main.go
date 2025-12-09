// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 29.
//!+

// Boiling prints the boiling point of water.
package main

import "fmt"

// 【Go vs Java】包级常量
// Java:  public static final double BOILING_F = 212.0;
// Go:    const boilingF = 212.0
// 注意：Go的常量首字母小写表示包内可见，大写表示导出（public）
const boilingF = 212.0

// main 计算并打印水的沸点
func main() {
	// 【Go vs Java】变量声明并初始化
	// Java:  double f = boilingF;
	// Go:    var f = boilingF (类型自动推断)
	// 注意：也可以写成 var f float64 = boilingF 或 f := boilingF
	var f = boilingF
	var c = (f - 32) * 5 / 9

	// 【Go vs Java】格式化输出
	// Java:  System.out.printf("boiling point = %f°F or %f°C\n", f, c);
	// Go:    fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	// 注意：%g 自动选择 %e 或 %f 格式，去掉无意义的尾零
	fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	// Output:
	// boiling point = 212°F or 100°C
}

//!-
