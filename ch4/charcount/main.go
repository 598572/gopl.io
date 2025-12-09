// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

// main 统计Unicode字符及其UTF-8编码长度
func main() {
	// 【Go vs Java】rune类型
	// Java:  Map<Integer, Integer> counts = new HashMap<>(); (用int表示字符)
	// Go:    counts := make(map[rune]int)
	// 注意：rune是int32的别名，表示Unicode码点，Go原生支持Unicode
	counts := make(map[rune]int) // counts of Unicode characters

	// 【Go vs Java】固定长度数组
	// Java:  int[] utflen = new int[utf8.UTFMax + 1];
	// Go:    var utflen [utf8.UTFMax + 1]int
	// 注意：数组长度必须是常量表达式
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	// 【Go vs Java】创建Reader
	// Java:  BufferedReader in = new BufferedReader(new InputStreamReader(System.in));
	// Go:    in := bufio.NewReader(os.Stdin)
	in := bufio.NewReader(os.Stdin)

	// 【Go vs Java】无限循环
	// Java:  while (true) { ... }
	// Go:    for { ... }
	// 注意：Go的for不带条件就是无限循环
	for {
		// 【Go vs Java】读取Unicode字符
		// Java:  int r = in.read(); (返回int，-1表示EOF)
		// Go:    r, n, err := in.ReadRune() (返回rune、字节数、错误)
		// 注意：ReadRune返回三个值，r是字符，n是UTF-8编码字节数
		r, n, err := in.ReadRune() // returns rune, nbytes, error

		// 【Go vs Java】EOF检查
		// Java:  if (r == -1) break;
		// Go:    if err == io.EOF { break }
		// 注意：Go用错误值表示EOF，不用特殊返回值
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		// 【Go vs Java】Unicode替换字符检查
		// Java:  if (r == 0xFFFD && n == 1) { ... }
		// Go:    if r == unicode.ReplacementChar && n == 1 { ... }
		// 注意：ReplacementChar (U+FFFD) 用于表示无效的UTF-8序列
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		// 【Go vs Java】map自增
		// Java:  counts.put(r, counts.getOrDefault(r, 0) + 1);
		// Go:    counts[r]++
		// 注意：Go的map访问不存在的key返回零值，可以直接++
		counts[r]++
		utflen[n]++
	}

	// 打印统计结果
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		// 【Go vs Java】字符格式化
		// Java:  System.out.printf("'%c'\t%d\n", c, n);
		// Go:    fmt.Printf("%q\t%d\n", c, n)
		// 注意：%q 输出带引号的字符，会转义特殊字符
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

//!-
