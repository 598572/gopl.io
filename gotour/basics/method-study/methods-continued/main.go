package main

import (
	"fmt"
	"math"
)

/*
*
方法（续）
你也可以为非结构体类型声明方法。

在此例中，我们看到了一个带 Abs 方法的数值类型 MyFloat。

你只能为在同一个包中定义的接收者类型声明方法，而不能为其它别的包中定义的类型 （包括 int 之类的内置类型）声明方法。

（译注：就是接收者的类型定义和方法声明必须在同一包内。）
*/
// ========== 为非结构体类型定义方法 ==========
// 【核心问题】如何为基本类型添加方法？
//
// ========== Java 的做法 ==========
// Java 中，float 是基本类型，不能直接添加方法。
// 如果你想要一个"有 abs() 方法的 float"，必须创建一个包装类：
//
//   class MyFloat {
//       private float value;  // 包装一个 float 值
//
//       public MyFloat(float value) {
//           this.value = value;  // 构造函数
//       }
//
//       public float abs() {  // 添加方法
//           return Math.abs(value);
//       }
//
//       public float getValue() {  // 还需要 getter
//           return value;
//       }
//   }
//
//   // 使用：
//   MyFloat f = new MyFloat(-1.414f);
//   float result = f.abs();  // 调用方法
//
// 缺点：
// - 需要写很多样板代码（类、字段、构造函数、getter）
// - 不是真正的 float，是一个对象
// - 使用起来比较繁琐
//
// ========== Go 的做法 ==========
// Go 中，可以直接通过类型别名来扩展基本类型：
//
//   type MyFloat float64  // 一行代码定义新类型
//
//   func (f MyFloat) Abs() float64 {  // 直接为类型添加方法
//       if f < 0 {
//           return float64(-f)
//       }
//       return float64(f)
//   }
//
//   // 使用：
//   f := MyFloat(-1.414)
//   result := f.Abs()  // 调用方法，就像调用 float 的方法一样！
//
// 优点：
// - 代码简洁（不需要类、字段、构造函数）
// - 本质上还是 float64，只是有别名
// - 使用起来很自然，就像 float 原生有这个方法
//
// ========== 类型说明 ==========
// type MyFloat float64 的含义：
// - 创建一个新类型 MyFloat
// - 它"基于" float64（底层存储方式相同）
// - 但它们是"不同的类型"（不能直接混用）
//
// 类比理解：
// - 就像 Java 中的：class MyFloat extends Number { ... }
// - 但 Go 的更轻量，不需要类的开销
//
// 重要限制：
// - 只能为"你自己定义的类型"添加方法
// - 不能直接为别人的类型（如标准库的 float64）添加方法
// - 但通过自定义类型，可以"扩展"基本类型的功能
type MyFloat float64

// ========== 为自定义类型定义方法 ==========
// 【Java 对比】
// Java 中，方法必须在类里面定义：
//
//	class MyFloat {
//	    private float value;
//	    public float abs() {  // 方法在类里面
//	        if (value < 0) {
//	            return -value;
//	        }
//	        return value;
//	    }
//	}
//
// 【Go 做法】
// Go 中，方法可以"挂"在类型上：
//
//	type MyFloat float64
//	func (f MyFloat) Abs() float64 { ... }  // 方法"属于"这个类型
//
// 方法说明：
// - func (f MyFloat) Abs() float64
//   - (f MyFloat) 是接收者（类似 Java 的 this）
//   - f 代表调用这个方法的 MyFloat 实例
//   - Abs() 方法计算绝对值
//
// 功能：计算绝对值
// - 如果 f < 0，返回 -f（取反）
// - 否则返回 f（本身）
func (f MyFloat) Abs() float64 {
	// 【Java 对比】
	// Java:  if (value < 0) { return -value; }
	// Go:    if f < 0 { return float64(-f) }
	//
	// 关键区别：
	// - Java: 访问类字段 this.value（或直接 value）
	// - Go:   直接使用接收者 f（就像使用变量一样）
	//
	// 类型转换说明：
	// - f 是 MyFloat 类型
	// - 返回值要求是 float64 类型
	// - 所以需要转换：float64(-f)
	if f < 0 {
		return float64(-f) // 负数：-(-1.414) = 1.414，然后转换为 float64
	}
	return float64(f) // 正数：直接转换为 float64
}

func main() {
	// ========== 创建和使用对比 ==========
	//
	// 【Java 代码】
	//   MyFloat f = new MyFloat(-Math.sqrt(2));  // 1. 创建对象
	//   float result = f.abs();                   // 2. 调用方法
	//   System.out.println(result);               // 3. 打印结果
	//
	// 【Go 代码】
	//   f := MyFloat(-math.Sqrt2)  // 1. 类型转换（不是创建对象！）
	//   result := f.Abs()           // 2. 调用方法
	//   fmt.Println(result)         // 3. 打印结果
	//
	// 关键区别：
	// - Java: new MyFloat(...) 创建对象，有对象开销
	// - Go:   MyFloat(...) 只是类型转换，没有对象开销，还是原始值
	//
	// 说明：
	// - math.Sqrt2 是常量，值为 √2 ≈ 1.4142135623730951
	// - -math.Sqrt2 是 -√2 ≈ -1.414...
	// - MyFloat(-math.Sqrt2) 将 float64 类型转换为 MyFloat 类型
	f := MyFloat(-math.Sqrt2) // f = -√2 ≈ -1.414

	// ========== 方法调用 ==========
	// 【Java】
	//   float result = f.abs();  // 调用方法
	//
	// 【Go】
	//   result := f.Abs()  // 调用方法（语法几乎相同！）
	//
	// 执行过程：
	// 1. f 的值是 -√2（负数）
	// 2. 调用 f.Abs()，进入 Abs() 方法
	// 3. 判断 f < 0 为真，执行 return float64(-f)
	// 4. -f = -(-√2) = √2
	// 5. float64(√2) 转换为 float64 类型
	// 6. 返回 √2 ≈ 1.4142135623730951
	fmt.Println(f.Abs()) // 输出: 1.4142135623730951
}

// ========== 完整对比示例 ==========
//
// 【Java 完整代码】
//   class MyFloat {
//       private float value;
//       public MyFloat(float value) { this.value = value; }
//       public float abs() {
//           return value < 0 ? -value : value;
//       }
//   }
//
//   public class Main {
//       public static void main(String[] args) {
//           MyFloat f = new MyFloat(-1.414f);
//           System.out.println(f.abs());  // 输出: 1.414
//       }
//   }
//
// 【Go 完整代码】
//   type MyFloat float64
//
//   func (f MyFloat) Abs() float64 {
//       if f < 0 {
//           return float64(-f)
//       }
//       return float64(f)
//   }
//
//   func main() {
//       f := MyFloat(-1.414)
//       fmt.Println(f.Abs())  // 输出: 1.414
//   }
//
// 对比总结：
// - Java：需要类、字段、构造函数 → 代码多，有对象开销
// - Go：  只需要类型定义和方法 → 代码少，无对象开销，更简洁

// ========== 核心概念总结 ==========
// 1. 可以为非结构体类型定义方法
//    - 可以为自定义类型（如 MyFloat）定义方法
//    - 不能直接为内置类型（如 float64、int）定义方法
//    - 但可以通过自定义类型间接实现
//
// 2. 类型限制
//    - 只能为"同一个包内定义的类型"定义方法
//    - 接收者类型和方法必须在同一包内
//
// 3. 类型转换
//    - MyFloat 和 float64 是不同的类型
//    - 需要显式转换：MyFloat(value) 或 float64(value)
//
// 4. 与 Java 的对比
//    - Java：基本类型不能有方法，需要包装类
//    - Go：通过自定义类型可以为任何类型添加方法
//
// 5. 实际应用
//    - 可以为基本类型添加便捷方法
//    - 例如：为 int 类型添加 IsEven()、IsOdd() 等方法
//    - 例如：为 string 类型添加自定义的格式化方法

// ========== 扩展示例 ==========
// 为 int 类型添加方法：
//   type MyInt int
//   func (i MyInt) IsEven() bool {
//       return i%2 == 0
//   }
//
//   i := MyInt(4)
//   fmt.Println(i.IsEven())  // true
//
// 为 string 类型添加方法：
//   type MyString string
//   func (s MyString) Reverse() string {
//       // 反转字符串的逻辑
//   }
