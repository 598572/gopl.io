package main

import "fmt"

/*
*

range 遍历
for 循环的 range 形式可遍历切片或映射。

当使用 for 循环遍历切片时，每次迭代都会返回两个值。 第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。
*/
// 【Go vs Java】包级变量
// Java:  private static final int[] pow = {1, 2, 4, 8, 16, 32, 64, 128};
// Go:    var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
// 注意：这是2的幂次方序列：2^0=1, 2^1=2, 2^2=4, ..., 2^7=128
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// ========== range 遍历切片 ==========
	// 【Go vs Java】遍历数组/列表
	// Java:  for (int i = 0; i < pow.length; i++) {
	//            int v = pow[i];
	//            System.out.printf("2**%d = %d\n", i, v);
	//        }
	//        或者：
	//        for (int i = 0; i < pow.size(); i++) {
	//            int v = pow.get(i);
	//            ...
	//        }
	// Go:    for i, v := range pow { ... }
	//
	// range 关键字说明：
	// - range pow 遍历切片 pow
	// - 每次迭代返回两个值：
	//   - i：当前元素的索引（index）
	//   - v：当前元素的值（value）的副本
	// - i 和 v 是循环内部声明的变量，作用域仅在 for 循环内
	//
	// 执行过程：
	// 第1次迭代：i=0, v=1  (pow[0])
	// 第2次迭代：i=1, v=2  (pow[1])
	// 第3次迭代：i=2, v=4  (pow[2])
	// ...
	// 第8次迭代：i=7, v=128 (pow[7])
	for i, v := range pow {
		// 【Go vs Java】格式化输出
		// Java:  System.out.printf("2**%d = %d\n", i, v);
		// Go:    fmt.Printf("2**%d = %d\n", i, v)
		// 注意：%d 是整数格式化，** 表示幂运算符号（这里是数学表示）
		//       输出类似：2**0 = 1, 2**1 = 2, 2**2 = 4, ...
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// ========== range 的其他用法 ==========
	// 1. 只要索引，不要值：
	//    for i := range pow { ... }
	//
	// 2. 只要值，不要索引（用 _ 丢弃索引）：
	//    for _, v := range pow { ... }
	//
	// 3. 只要值（简写，不推荐，因为不清楚）：
	//    for v := range pow { ... }  // 注意：这样 v 得到的是索引，不是值！
	//
	// 4. 空循环（只要索引）：
	//    for range pow { ... }
}

// ========== 输出示例 ==========
// 2**0 = 1
// 2**1 = 2
// 2**2 = 4
// 2**3 = 8
// 2**4 = 16
// 2**5 = 32
// 2**6 = 64
// 2**7 = 128
