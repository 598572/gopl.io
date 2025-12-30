package main

import (
	"fmt"
	"math"
)

/*
*
方法即函数
记住：方法只是个带接收者参数的函数。

现在这个 Abs 的写法就是个正常的函数，功能并没有什么变化。
*/
type Vertex struct {
	X, Y float64 // 二维坐标点
}

// ========== 函数（Function）写法 ==========
// 【Go vs Java】函数定义
// Java:  没有独立的函数，只有方法
// Go:    func Abs(v Vertex) float64 { ... }
//
// 这是普通的函数（不是方法）
// 语法说明：
// - func Abs(v Vertex) float64
//   - func 关键字
//   - Abs 是函数名
//   - v Vertex 是参数（类型 Vertex）
//   - float64 是返回类型
//
// 函数 vs 方法对比：
// - 方法：func (v Vertex) Abs() float64 { ... }  ← 有接收者
// - 函数：func Abs(v Vertex) float64 { ... }     ← 没有接收者，参数在括号内
//
// 功能：计算点到原点的距离（与方法的实现完全相同）
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}

	// ========== 函数调用 ==========
	// 【Go vs Java】函数调用
	// Java:  没有独立函数，只有方法调用 v.abs()
	// Go:    Abs(v)  ← 函数调用，需要显式传入参数
	//
	// 函数调用语法：
	// - Abs(v) 调用函数 Abs，传入参数 v
	// - 与方法调用 v.Abs() 功能相同，但语法不同
	//
	// 方法调用 vs 函数调用对比：
	// - 方法：v.Abs()      ← 实例在前，方法名在后（面向对象风格）
	// - 函数：Abs(v)       ← 函数名在前，参数在后（函数式风格）
	//
	// 计算：Abs(3, 4) = √(3² + 4²) = √25 = 5
	fmt.Println(Abs(v)) // 输出: 5
}

// ========== 方法即函数：核心理解 ==========
// 在 Go 中，方法本质上就是带接收者参数的函数！
//
// 等价关系：
// 方法写法：func (v Vertex) Abs() float64 { ... }
//          调用：v.Abs()
//
// 函数写法：func Abs(v Vertex) float64 { ... }
//          调用：Abs(v)
//
// 编译器视角：这两种写法在底层是等价的
// - v.Abs() 会被编译器转换为 Vertex.Abs(v)
// - 方法只是语法糖，让代码更符合面向对象的习惯
//
// 选择建议：
// 1. 如果是类型相关的操作 → 使用方法（更符合 OOP 习惯）
//    func (v Vertex) Abs() float64 { ... }
//
// 2. 如果是通用工具函数 → 使用函数（更简洁）
//    func Abs(v Vertex) float64 { ... }
//
// 3. 需要实现接口 → 必须使用方法
//    （接口要求方法签名，不能用函数）

// ========== 实际对比示例 ==========
// 方法写法（上一段代码）：
//   type Vertex struct { X, Y float64 }
//   func (v Vertex) Abs() float64 { return math.Sqrt(v.X*v.X + v.Y*v.Y) }
//   v := Vertex{3, 4}
//   fmt.Println(v.Abs())  // 调用方法
//
// 函数写法（当前代码）：
//   type Vertex struct { X, Y float64 }
//   func Abs(v Vertex) float64 { return math.Sqrt(v.X*v.X + v.Y*v.Y) }
//   v := Vertex{3, 4}
//   fmt.Println(Abs(v))   // 调用函数
//
// 功能完全相同，只是调用方式不同！
