package main

import (
	"fmt"
	"math"
)

/*
函数值
函数也是值。它们可以像其他值一样传递。

函数值可以用作函数的参数或返回值。
*/
// ========== 函数作为参数 ==========
// compute 函数接收一个函数作为参数
// 【Go vs Java】函数作为参数
// Java:  @FunctionalInterface
//        interface BiFunction<T, U, R> { R apply(T t, U u); }
//
//        double compute(BiFunction<Double, Double, Double> fn) {
//            return fn.apply(3.0, 4.0);
//        }
//
// Go:    func compute(fn func(float64, float64) float64) float64 { ... }
//
// 函数类型说明：
// - fn func(float64, float64) float64
//   - fn 是参数名
//   - func(float64, float64) float64 是函数类型
//     - 接收两个 float64 参数
//     - 返回一个 float64 值
//
// 功能：
// - compute 函数接收一个函数 fn，然后用参数 (3, 4) 调用它
// - 这展示了"函数作为值"的概念
func compute(fn func(float64, float64) float64) float64 {
	// 调用传入的函数 fn，参数是 (3, 4)
	// 例如：如果 fn 是 math.Pow，则执行 math.Pow(3, 4) = 81
	return fn(3, 4)
}

func main() {
	// ========== 匿名函数（函数字面量）==========
	// 【Go vs Java】匿名函数/Lambda
	// Java:  BiFunction<Double, Double, Double> hypot = (x, y) -> {
	//            return Math.sqrt(x * x + y * y);
	//        };
	//
	// Go:    hypot := func(x, y float64) float64 {
	//            return math.Sqrt(x*x + y*y)
	//        }
	//
	// 解释：
	// - func(x, y float64) float64 { ... } 是匿名函数（函数字面量）
	// - hypot 是函数变量，类型是 func(float64, float64) float64
	// - 这个函数计算直角三角形的斜边长度（勾股定理）
	//   - 公式：√(x² + y²)
	//   - 例如：hypot(3, 4) = √(9 + 16) = √25 = 5
	hypot := func(x, y float64) float64 {
		// math.Sqrt 计算平方根
		// x*x + y*y 计算 x² + y²
		return math.Sqrt(x*x + y*y)
	}

	// ========== 调用函数变量 ==========
	// 【Go vs Java】调用函数变量
	// Java:  double result = hypot.apply(5, 12);
	// Go:    result := hypot(5, 12)
	//
	// 计算：hypot(5, 12) = √(5² + 12²) = √(25 + 144) = √169 = 13
	fmt.Println(hypot(5, 12)) // 输出: 13

	// ========== 将函数作为参数传递 ==========
	// 【Go vs Java】函数作为参数传递
	// Java:  compute(hypot);
	// Go:    compute(hypot)  // 直接传递函数变量
	//
	// 执行过程：
	// 1. compute 接收 hypot 函数作为参数
	// 2. compute 内部执行 hypot(3, 4)
	// 3. hypot(3, 4) = √(3² + 4²) = √(9 + 16) = √25 = 5
	fmt.Println(compute(hypot)) // 输出: 5
	// 等价于：fmt.Println(hypot(3, 4))

	// ========== 传递标准库函数 ==========
	// 【Go vs Java】传递标准库函数
	// Java:  compute((x, y) -> Math.pow(x, y));
	// Go:    compute(math.Pow)  // 直接传递函数名
	//
	// math.Pow 是 Go 标准库函数，计算 x 的 y 次方
	// 执行过程：
	// 1. compute 接收 math.Pow 函数作为参数
	// 2. compute 内部执行 math.Pow(3, 4)
	// 3. math.Pow(3, 4) = 3⁴ = 81
	fmt.Println(compute(math.Pow)) // 输出: 81
	// 等价于：fmt.Println(math.Pow(3, 4))
}

// ========== 核心概念总结 ==========
// 1. 函数是"一等公民"（First-class citizen）
//    - 函数可以作为变量
//    - 函数可以作为参数传递
//    - 函数可以作为返回值
//
// 2. 匿名函数（函数字面量）
//    - func(x, y float64) float64 { ... } 创建匿名函数
//    - 可以赋值给变量
//
// 3. 函数类型
//    - func(float64, float64) float64 是函数类型
//    - 描述了函数的签名（参数类型和返回类型）
//
// 4. 与 Java 的对比
//    - Java 8+ 有 Lambda 表达式和函数式接口
//    - Go 的函数作为值更加自然和简洁
//    - Go 不需要接口，函数类型直接作为参数类型
