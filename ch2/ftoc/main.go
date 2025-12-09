// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 29.
//!+

// Ftoc prints two Fahrenheit-to-Celsius conversions.
package main

import "fmt"

// main 演示华氏度到摄氏度的转换
func main() {
	// 【Go vs Java】多变量常量声明
	// Java:  final double freezingF = 32.0, boilingF = 212.0;
	// Go:    const freezingF, boilingF = 32.0, 212.0
	// 注意：Go支持多重赋值，可以同时声明多个常量
	const freezingF, boilingF = 32.0, 212.0

	// 【Go vs Java】函数调用
	// Java:  fToC(freezingF) (相同)
	// Go:    fToC(freezingF) (相同)
	fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF)) // "32°F = 0°C"
	fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))   // "212°F = 100°C"
}

// fToC 将华氏度转换为摄氏度
// 【Go vs Java】函数定义
// Java:  double fToC(double f) { return (f - 32) * 5 / 9; }
// Go:    func fToC(f float64) float64 { return (f - 32) * 5 / 9 }
// 注意：Go的返回类型在参数列表之后
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

//!-
