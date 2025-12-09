// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 41.

//!+

package tempconv

// CToF converts a Celsius temperature to Fahrenheit.
// 【Go vs Java】函数命名和导出
// Java:  public static Fahrenheit cToF(Celsius c) { ... }
// Go:    func CToF(c Celsius) Fahrenheit { ... }
// 注意：首字母大写的函数是导出的（public），小写是包内私有的
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC converts a Fahrenheit temperature to Celsius.
// 【Go vs Java】类型转换
// Java:  return new Celsius((f.getValue() - 32) * 5 / 9);
// Go:    return Celsius((f - 32) * 5 / 9)
// 注意：Celsius(x) 是类型转换，将float64转为Celsius类型
//
//	自定义类型可以直接参与运算，结果是底层类型（float64）
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//!-
