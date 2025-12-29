// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//!+array
	// 【Go vs Java】数组 vs main
	// Java:  int[] a = {0, 1, 2, 3, 4, 5}; (数组)
	// Go:    a := [...]int{0, 1, 2, 3, 4, 5} (数组，...自动推断长度)
	// 注意：[...]int 是数组，[]int 是切片，数组长度固定且是类型的一部分
	a := [...]int{0, 1, 2, 3, 4, 5}

	// 【Go vs Java】数组转切片
	// Java:  无需转换（数组和List是不同类型）
	// Go:    a[:] 将数组转为切片
	// 注意：a[:] 创建引用整个数组的切片，修改切片会影响原数组
	reverse(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array

	//!+slice
	// 【Go vs Java】切片字面量
	// Java:  List<Integer> s = Arrays.asList(0, 1, 2, 3, 4, 5);
	// Go:    s := []int{0, 1, 2, 3, 4, 5}
	s := []int{0, 1, 2, 3, 4, 5}

	// Rotate s left by two positions.
	// 【Go vs Java】切片操作实现旋转
	// Java:  Collections.rotate(s, -2);
	// Go:    通过三次反转实现：reverse(s[:2]) + reverse(s[2:]) + reverse(s)
	// 注意：s[:2] 是前2个元素，s[2:] 是从索引2到末尾
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"
	//!-slice

	// Interactive test of reverse.
	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		var ints []int
		for _, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				fmt.Fprintln(os.Stderr, err)
				continue outer
			}
			ints = append(ints, int(x))
		}
		reverse(ints)
		fmt.Printf("%v\n", ints)
	}
	// NOTE: ignoring potential errors from input.Err()
}

// !+rev
// reverse reverses a slice of ints in place.
// 【Go vs Java】原地反转
// Java:  void reverse(int[] s) { for (int i=0, j=s.length-1; i<j; i++, j--) {...} }
// Go:    func reverse(s []int) { for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {...} }
func reverse(s []int) {
	// 【Go vs Java】多变量初始化和更新
	// Java:  for (int i = 0, j = s.length - 1; i < j; i++, j--)
	// Go:    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1
	// 注意：Go的for支持多变量同时初始化和更新
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		// 【Go vs Java】交换两个变量
		// Java:  int temp = s[i]; s[i] = s[j]; s[j] = temp;
		// Go:    s[i], s[j] = s[j], s[i]
		// 注意：Go支持多重赋值，可以优雅地交换变量，无需临时变量
		s[i], s[j] = s[j], s[i]
	}
}

//!-rev
