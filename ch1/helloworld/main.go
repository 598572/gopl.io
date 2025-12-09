// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 1.

// Helloworld is our first Go program.
// !+

// 【Go vs Java】包声明
// Java:  public class HelloWorld { public static void main(String[] args) {...} }
// Go:    package main + func main() 即可，无需类包装
// 注意：Go的可执行程序必须在 package main 中，且必须有 main() 函数
package main

// 【Go vs Java】导入包
// Java:  import java.lang.System; (可省略java.lang)
// Go:    import "fmt" (格式化输入输出包)
// 注意：Go的import路径用双引号，未使用的import会编译错误
import "fmt"

// main 是程序的入口函数
// 【Go vs Java】main函数签名
// Java:  public static void main(String[] args)
// Go:    func main() - 无参数、无返回值、无修饰符
// 注意：命令行参数通过 os.Args 获取，不是函数参数
func main() {
	// 【Go vs Java】打印输出
	// Java:  System.out.println("Hello, world");
	// Go:    fmt.Println("Hello, world")
	// 注意：Go的函数名首字母大写表示公开（类似Java的public）
	fmt.Println("Hello, world，我是蝎子莱莱，我来了")
}

//!-
