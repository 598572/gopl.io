// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 77.

// Netflag demonstrates an integer type used as a bit field.
package main

import (
	"fmt"
	// 【Go vs Java】点导入
	// Java:  import static java.net.Flags.*;
	// Go:    import . "net"
	// 注意：. 导入后可以直接使用包内的导出标识符，不需要包名前缀
	//      不推荐使用，容易造成命名冲突，这里仅作演示
	. "net"
)

//!+
// 【Go vs Java】位运算操作
// Java:  使用 & | ^ ~ 等位运算符
// Go:    同样使用位运算符，但有特殊的 &^ (AND NOT) 运算符

// IsUp 检查是否设置了FlagUp标志
// 【Go vs Java】位与操作检查标志
// Java:  boolean isUp(int v) { return (v & FLAG_UP) == FLAG_UP; }
// Go:    func IsUp(v Flags) bool { return v&FlagUp == FlagUp }
func IsUp(v Flags) bool { return v&FlagUp == FlagUp }

// TurnDown 清除FlagUp标志
// 【Go vs Java】清除位标志
// Java:  void turnDown(Flags v) { v &= ~FLAG_UP; }
// Go:    func TurnDown(v *Flags) { *v &^= FlagUp }
// 注意：&^ 是Go特有的"AND NOT"运算符，x &^ y 等价于 x & (^y)
//
//	*v 解引用指针，修改原值
func TurnDown(v *Flags) { *v &^= FlagUp }

// SetBroadcast 设置FlagBroadcast标志
// 【Go vs Java】设置位标志
// Java:  void setBroadcast(Flags v) { v |= FLAG_BROADCAST; }
// Go:    func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
// 注意：|= 是位或赋值运算符
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }

// IsCast 检查是否设置了广播或多播标志
func IsCast(v Flags) bool { return v&(FlagBroadcast|FlagMulticast) != 0 }

func main() {
	// 【Go vs Java】位或组合多个标志
	// Java:  Flags v = FLAG_MULTICAST | FLAG_UP;
	// Go:    var v Flags = FlagMulticast | FlagUp
	var v Flags = FlagMulticast | FlagUp

	// 【Go vs Java】二进制格式化输出
	// Java:  System.out.printf("%s %b\n", Integer.toBinaryString(v), isUp(v));
	// Go:    fmt.Printf("%b %t\n", v, IsUp(v))
	// 注意：%b 输出二进制格式，%t 输出布尔值
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"

	// 【Go vs Java】传递指针修改值
	// Java:  turnDown(v); (Java对象引用自动传递)
	// Go:    TurnDown(&v) (必须显式传递指针)
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"

	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
}

//!-
