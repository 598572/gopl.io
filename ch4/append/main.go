// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 88.

// Append illustrates the behavior of the built-in append function.
package main

import "fmt"

func appendslice(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// There is room to expand the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}

// !+append
// appendInt 演示切片append的底层实现原理
// 【Go vs Java】main vs ArrayList
// Java:  ArrayList<Integer> list = new ArrayList<>(); list.add(y);
// Go:    slice = append(slice, y)
// 注意：Go的切片是对底层数组的引用，包含指针、长度、容量三个字段
//
//	len(x) 是当前元素个数，cap(x) 是底层数组容量
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1

	// 【Go vs Java】容量检查
	// Java:  ArrayList自动扩容
	// Go:    需要手动检查容量，或使用内置append函数
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		// 【Go vs Java】切片扩展
		// Java:  无对应操作（ArrayList自动管理）
		// Go:    z = x[:zlen] 重新切片，共享底层数组
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}

		// 【Go vs Java】创建切片
		// Java:  int[] z = new int[zcap];
		// Go:    z = make([]int, zlen, zcap)
		// 注意：make([]T, len, cap) 创建长度为len、容量为cap的切片
		z = make([]int, zlen, zcap)

		// 【Go vs Java】复制切片
		// Java:  System.arraycopy(x, 0, z, 0, x.length);
		// Go:    copy(z, x)
		// 注意：copy是内置函数，返回实际复制的元素个数
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}

//!-append

// !+growth
func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

//!-growth

/*
//!+output
0  cap=1   [0]
1  cap=2   [0 1]
2  cap=4   [0 1 2]
3  cap=4   [0 1 2 3]
4  cap=8   [0 1 2 3 4]
5  cap=8   [0 1 2 3 4 5]
6  cap=8   [0 1 2 3 4 5 6]
7  cap=8   [0 1 2 3 4 5 6 7]
8  cap=16  [0 1 2 3 4 5 6 7 8]
9  cap=16  [0 1 2 3 4 5 6 7 8 9]
//!-output
*/
