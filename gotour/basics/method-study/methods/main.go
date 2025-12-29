package main

import (
	"fmt"
	"math"
)

/*
*

方法
Go 没有类。不过你可以为类型定义方法。

方法就是一类带特殊的 接收者 参数的函数。

方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。

在此例中，Abs 方法拥有一个名字为 v，类型为 Vertex 的接收者。
*/
// ========== 定义结构体类型 ==========
// 【Go vs Java】结构体定义
// Java:  class Vertex {
//            private double x;
//            private double y;
//            public Vertex(double x, double y) { this.x = x; this.y = y; }
//        }
// Go:    type Vertex struct { X, Y float64 }
//
// 解释：
// - Vertex 是一个结构体类型，表示二维坐标点
// - X, Y 是字段名（首字母大写表示导出/public）
// - float64 是字段类型（浮点数）
type Vertex struct {
	X, Y float64 // X 坐标和 Y 坐标
}

// ========== 为类型定义方法 ==========
// 【Go vs Java】方法定义
//
//	Java:  class Vertex {
//	           public double abs() {
//	               return Math.sqrt(x * x + y * y);
//	           }
//	       }
//
// Go:    func (v Vertex) Abs() float64 { ... }
//
// 方法语法说明：
// - func (v Vertex) Abs() float64
//   - func 关键字
//   - (v Vertex) 是接收者（receiver）
//   - v 是接收者变量名（类似 Java 的 this，但可以自定义）
//   - Vertex 是接收者类型（方法属于这个类型）
//   - Abs 是方法名
//   - float64 是返回类型
//
// 接收者说明：
// - 接收者写在 func 关键字和方法名之间
// - (v Vertex) 表示这个方法是 Vertex 类型的方法
// - v 是方法内部的变量，代表调用该方法的 Vertex 实例
// - 在方法内部可以通过 v.X 和 v.Y 访问字段
//
// 功能：计算点到原点的距离（向量的模长）
// 公式：√(x² + y²)
func (v Vertex) Abs() float64 {
	// v.X 和 v.Y 访问接收者的字段
	// v.X*v.X + v.Y*v.Y 计算 x² + y²
	// math.Sqrt() 计算平方根
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// ========== 创建结构体实例 ==========
	// 【Go vs Java】创建实例
	// Java:  Vertex v = new Vertex(3, 4);
	// Go:    v := Vertex{3, 4}
	//
	// 创建坐标点 (3, 4)
	v := Vertex{3, 4}

	// ========== 调用方法 ==========
	// 【Go vs Java】方法调用
	// Java:  double result = v.abs();
	// Go:    v.Abs()
	//
	// 语法说明：
	// - v.Abs() 调用 Vertex 类型的方法 Abs
	// - v 是接收者（类似 Java 中的 this）
	// - 计算：Abs(3, 4) = √(3² + 4²) = √(9 + 16) = √25 = 5
	fmt.Println(v.Abs())       // 输出: 5
	fmt.Println(Vertex.Abs(v)) // 输出: 5 和上边的等价

}

// ========== 核心概念总结 ==========
// 1. Go 没有类（class），但有方法（method）
//    - 方法可以绑定到任何类型（不仅仅是结构体）
//    - 方法通过接收者（receiver）实现
//
// 2. 接收者（Receiver）
//    - 语法：(变量名 类型) 写在 func 和方法名之间
//    - 类似 Java 的 this，但可以自定义名字
//    - 可以是值接收者 (v Vertex) 或指针接收者 (v *Vertex)
//
// 3. 方法 vs 函数
//    - 方法：属于某个类型，通过 实例.方法名() 调用
//    - 函数：独立存在，通过 函数名() 调用
//
// 4. 与 Java 的对比
//    - Java: 方法必须在类中定义
//    - Go:   方法可以绑定到任何类型，更灵活
//
// 5. 方法调用的本质
//    - v.Abs() 等价于 Vertex.Abs(v)（函数调用）
//    - Go 编译器会自动转换

// ========== 方法的不同写法 ==========
// 值接收者（当前代码）：
//   func (v Vertex) Abs() float64 { ... }
//   - 方法内部操作的是 v 的副本
//   - 不会修改原始值
//
// 指针接收者（可以修改原值）：
//   func (v *Vertex) Scale(f float64) {
//       v.X = v.X * f
//       v.Y = v.Y * f
//   }
//   - 方法内部操作的是指针指向的值
//   - 可以修改原始值
