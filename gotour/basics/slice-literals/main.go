package main

import "fmt"

/*
*

切片字面量

切片字面量类似于没有长度的数组字面量。
这是一个数组字面量：
[3]bool{true, true, false}
下面这样则会创建一个和上面相同的数组，然后再构建一个引用了它的切片：

[]bool{true, true, false}
*/
func main() {
	//fmt.Println(q)

	// ========== 示例1：基础切片字面量 ==========
	// 【Go vs Java】切片字面量
	// Java:  List<Boolean> r = Arrays.asList(true, false, true, true, false, true);
	//        或者 boolean[] r = {true, false, true, true, false, true};
	// Go:    r := []bool{true, false, true, true, false, true}
	//
	// 解释：
	// - []bool 表示"bool类型的切片"（类似Java的List<Boolean>，但更底层）
	// - {true, false, ...} 是切片字面量，直接初始化元素
	// - Go会自动推断切片的长度（这里是6个元素）
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r) // 输出: [true false true true false true]

	// ========== 示例2：结构体切片字面量（重点！）==========
	// 【Go vs Java】匿名结构体切片
	// Java:  需要先定义一个类，然后创建List
	//        class Item { int i; boolean b; }
	//        List<Item> s = Arrays.asList(
	//            new Item(2, true), new Item(3, false), ...
	//        );
	//
	// Go:    可以直接用匿名结构体（不需要先定义类）
	//        []struct { i int; b bool }{ {2, true}, {3, false}, ... }
	//
	// 分步解释：
	// 1. []struct { i int; b bool } 表示：
	//    - [] 是切片
	//    - struct { i int; b bool } 是匿名结构体（临时定义，没有名字）
	//    - 结构体有两个字段：i (int类型) 和 b (bool类型)
	//
	// 2. { {2, true}, {3, false}, ... } 是结构体的初始化：
	//    - 每个 {2, true} 创建一个结构体实例
	//    - {2, true} 中：2 对应字段 i，true 对应字段 b
	//    - 按字段定义的顺序赋值（i在前，b在后）
	s := []struct {
		i int  // 结构体字段1：整数类型
		b bool // 结构体字段2：布尔类型
	}{
		{2, true},   // 创建第一个结构体：i=2, b=true
		{3, false},  // 创建第二个结构体：i=3, b=false
		{5, true},   // 创建第三个结构体：i=5, b=true
		{7, true},   // 创建第四个结构体：i=7, b=true
		{11, false}, // 创建第五个结构体：i=11, b=false
		{13, true},  // 创建第六个结构体：i=13, b=true
	}
	// 输出类似: [{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]
	fmt.Println(s)

	// ========== 补充说明 ==========
	// 如果你想更清晰地理解，可以这样拆分：
	//
	// // 1. 先定义结构体类型
	// type Item struct {
	//     i int
	//     b bool
	// }
	//
	// // 2. 然后用这个类型创建切片
	// s := []Item{
	//     {2, true},
	//     {3, false},
	//     ...
	// }
	//
	// 但是Go允许"匿名结构体"，不需要定义类型就可以直接使用（如上面的代码）
}
