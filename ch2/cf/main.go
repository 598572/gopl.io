// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 43.
//!+

// Cf converts its numeric argument to Celsius and Fahrenheit.
package main

import (
	"fmt"
	"os"
	"strconv" // 字符串转换包

	// 【Go vs Java】导入自定义包
	// Java:  import com.example.tempconv.Celsius;
	// Go:    import "gopl.io/ch2/tempconv"
	// 注意：Go的import路径是从模块根目录开始的完整路径
	"gopl.io/ch2/tempconv"
)

// main 将命令行参数转换为摄氏度和华氏度
func main() {
	for _, arg := range os.Args[1:] {
		// 【Go vs Java】字符串转浮点数
		// Java:  double t = Double.parseDouble(arg);
		// Go:    t, err := strconv.ParseFloat(arg, 64)
		// 注意：64表示float64（双精度），返回(结果, 错误)两个值
		t, err := strconv.ParseFloat(arg, 64)

		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}

		// 【Go vs Java】类型转换（自定义类型）
		// Java:  Fahrenheit f = new Fahrenheit(t);
		// Go:    f := tempconv.Fahrenheit(t)
		// 注意：Fahrenheit是基于float64的自定义类型，可以直接转换
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)

		// 【Go vs Java】调用包中的函数
		// Java:  TempConv.fToC(f)
		// Go:    tempconv.FToC(f)
		// 注意：Go的包名通常是路径的最后一部分（这里是tempconv）
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}

//!-
