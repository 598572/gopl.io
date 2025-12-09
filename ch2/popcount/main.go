// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
// !+
package popcount

// 【Go vs Java】数组声明
// Java:  private static final byte[] pc = new byte[256];
// Go:    var pc [256]byte
// 注意：[256]byte 是数组（固定长度），[]byte 是切片（动态长度）
//      数组的长度是类型的一部分，[256]byte 和 [128]byte 是不同类型

// pc[i] is the population count of i.
var pc [256]byte

// init 包初始化函数，在main之前自动执行
// 【Go vs Java】初始化块
// Java:  static { for (int i = 0; i < pc.length; i++) {...} }
// Go:    func init() { for i := range pc {...} }
// 注意：init函数无参数、无返回值，每个包可以有多个init函数
//
//	执行顺序：包级变量初始化 -> init函数 -> main函数
func init() {
	// 【Go vs Java】range遍历数组
	// Java:  for (int i = 0; i < pc.length; i++)
	// Go:    for i := range pc (只要索引，不要元素值)
	// 注意：range返回(索引, 值)，这里只要索引
	for i := range pc {
		// 【Go vs Java】位运算
		// Java:  pc[i] = (byte)(pc[i/2] + (i & 1));
		// Go:    pc[i] = pc[i/2] + byte(i&1)
		// 注意：i&1 判断最低位是否为1（奇偶性）
		//      这是动态规划：pc[i] = pc[i/2] + (i的最低位)
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
// 【Go vs Java】位操作和查表法
// Java:  public static int popCount(long x) { ... }
// Go:    func PopCount(x uint64) int { ... }
// 注意：将64位整数分成8个字节，分别查表，然后相加
func PopCount(x uint64) int {
	// 【Go vs Java】位移和类型转换
	// Java:  (byte)(x >> (0 * 8))
	// Go:    byte(x>>(0*8))
	// 注意：>> 是右移运算符，x>>(n*8) 取第n个字节
	//      byte(x) 取x的最低8位
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

//!-
