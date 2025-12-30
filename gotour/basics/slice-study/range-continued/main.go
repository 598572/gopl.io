package main

import "fmt"

/*
range 遍历（续）
可以将下标或值赋予 _ 来忽略它。

for i, _ := range pow
for _, value := range pow
若你只需要索引，忽略第二个变量即可。

for i := range pow
*/
func main() {
	// ========== 第一步：创建切片 ==========
	// 【Go vs Java】创建切片/数组
	// Java:  int[] pow = new int[10];  // 数组，长度固定为10
	// Go:    pow := make([]int, 10)
	//
	// make([]int, 10) 的含义：
	// - 创建一个长度为10的 int 类型切片
	// - 所有元素初始化为零值（int 的零值是 0）
	// - 初始值：[0, 0, 0, 0, 0, 0, 0, 0, 0, 0]
	pow := make([]int, 10)
	// ========== 第二步：使用索引填充切片 ==========
	// 【Go vs Java】只要索引的遍历
	// Java:  for (int i = 0; i < pow.length; i++) {
	//            pow[i] = 1 << i;  // 或 (int)Math.pow(2, i);
	//        }
	// Go:    for i := range pow { pow[i] = 1 << uint(i) }
	//
	// 语法说明：
	// - for i := range pow 只要索引，不要值
	// - 等价于 for i, _ := range pow，但更简洁
	// - i 的值从 0 到 len(pow)-1（即 0 到 9）
	for i := range pow {
		// 【Go vs Java】位运算计算 2 的幂次方
		// Java:  pow[i] = 1 << i;  // 左移运算
		// Go:    pow[i] = 1 << uint(i)
		//
		// 位运算解释：
		// - 1 << n 表示把 1 左移 n 位，结果是 2^n
		// - 例如：
		//   1 << 0 = 1    (二进制：0001 左移0位 = 1)
		//   1 << 1 = 2    (二进制：0001 左移1位 = 0010 = 2)
		//   1 << 2 = 4    (二进制：0001 左移2位 = 0100 = 4)
		//   1 << 3 = 8    (二进制：0001 左移3位 = 1000 = 8)
		//
		// 为什么用 uint(i)？
		// - 左移运算符 << 要求右操作数是无符号整数
		// - i 是 int 类型，需要转换为 uint 类型
		// - uint(i) 是类型转换，将 int 转为 uint
		//
		// 等价写法：
		// - pow[i] = 1 << uint(i)  ← 位运算（推荐，高效）
		// - pow[i] = int(math.Pow(2, float64(i)))  ← 数学函数（慢）
		pow[i] = 1 << uint(i) // == 2**i（注释说明这等价于 2 的 i 次方）
	}
	// 执行后 pow 的值：[1, 2, 4, 8, 16, 32, 64, 128, 256, 512]

	// ========== 第三步：只遍历值，不要索引 ==========
	// 【Go vs Java】只要值的遍历
	// Java:  for (int value : pow) {  // 增强 for 循环
	//            System.out.println(value);
	//        }
	// Go:    for _, value := range pow { ... }
	//
	// 语法说明：
	// - for _, value := range pow 只要值，不要索引
	// - _ 是空白标识符（blank identifier），用于丢弃不需要的值
	// - _ 表示"我知道这里有值，但我不需要它"
	// - value 是当前元素的值
	for _, value := range pow {
		// 【Go vs Java】打印值
		// Java:  System.out.println(value);
		// Go:    fmt.Printf("%d\n", value)
		// 注意：%d 表示十进制整数，\n 表示换行
		fmt.Printf("%d\n", value)
	}
	// 输出结果：
	//   1
	//   2
	//   4
	//   8
	//   16
	//   32
	//   64
	//   128
	//   256
	//   512
}

// ========== range 的不同用法总结 ==========
// 1. 要索引和值：for i, v := range pow { ... }
// 2. 只要索引：   for i := range pow { ... }
// 3. 只要值：     for _, v := range pow { ... }
// 4. 都不要：     for range pow { ... }  （很少用）
//
// 关键点：
// - _ 是空白标识符，用于丢弃不需要的返回值
// - Go 要求所有变量都必须被使用，用 _ 可以"假装"使用了
// - 1 << n 是高效的位运算，计算 2 的 n 次方（比 math.Pow 快得多）
