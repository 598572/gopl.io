// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 39.
//!+

// Package tempconv performs Celsius and Fahrenheit temperature computations.
package tempconv

import "fmt"

// 【Go vs Java】自定义类型（基础版本）
// Java:  需要创建完整的类，包含字段、构造器、getter/setter
// Go:    一行代码定义新类型，自动继承底层类型的所有操作
type Celsius float64
type Fahrenheit float64

// 【Go vs Java】常量组
// Java:  public static final Celsius ABSOLUTE_ZERO_C = new Celsius(-273.15);
//
//	public static final Celsius FREEZING_C = new Celsius(0);
//
// Go:    用括号组合多个常量，更简洁
const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

// CToF 摄氏度转华氏度
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

// FToC 华氏度转摄氏度
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//!-

// String 实现fmt.Stringer接口，自定义打印格式
// 【Go vs Java】接口实现
// Java:  public class Celsius implements Stringer { @Override public String toString() {...} }
// Go:    只要有String()方法就自动实现了fmt.Stringer接口（隐式实现）
// 注意：Go的接口是隐式实现的，不需要显式声明implements
func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }
