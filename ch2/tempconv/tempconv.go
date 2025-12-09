// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

//!+

// Package tempconv performs Celsius and Fahrenheit conversions.
// 【Go vs Java】包注释
// Java:  使用 /** */ JavaDoc 注释在类上
// Go:    在 package 语句前用 // 注释整个包
// 注意：包注释应该以 "Package 包名" 开头
package tempconv

import "fmt"

// 【Go vs Java】自定义类型
// Java:  public class Celsius { private double value; ... }
// Go:    type Celsius float64
// 注意：Go的type定义了新类型，基于已有类型（这里是float64）
//
//	Celsius和float64是不同类型，不能直接混用，需要显式转换
type Celsius float64
type Fahrenheit float64

// 【Go vs Java】类型化常量
// Java:  public static final Celsius ABSOLUTE_ZERO_C = new Celsius(-273.15);
// Go:    const AbsoluteZeroC Celsius = -273.15
// 注意：Go的常量可以指定类型，首字母大写表示导出（public）
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// String 方法让Celsius类型实现fmt.Stringer接口
// 【Go vs Java】方法定义
// Java:  public String toString() { return String.format("%g°C", value); }
// Go:    func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
// 注意：(c Celsius) 是接收者，表示这是Celsius类型的方法
//
//	c是接收者的名字（类似Java的this，但可以自定义名字）
func (c Celsius) String() string    { return fmt.Sprintf("%g°C", c) }
func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

//!-
