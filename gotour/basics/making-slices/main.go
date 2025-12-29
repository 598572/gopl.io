package main

import "fmt"

/*
*

用 make 创建切片

切片可以用内置函数 make 来创建，这也是你创建动态数组的方式。
make 函数会分配一个元素为零值的数组并返回一个引用了它的切片：
a := make([]int, 5)  // len(a)=5
要指定它的容量，需向 make 传入第三个参数：
b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
*/
func main() {
	// ========== 示例1：make([]T, length) ==========
	// 【Go vs Java】创建切片/数组
	// Java:  int[] a = new int[5];  // 数组，长度固定为5
	//        List<Integer> a = new ArrayList<>(5);  // 列表，初始容量5
	// Go:    a := make([]int, 5)
	//
	// make([]int, 5) 的含义：
	// - 创建一个 int 类型的切片
	// - 长度（len）= 5，容量（cap）= 5（默认等于长度）
	// - 所有元素自动初始化为零值（int 的零值是 0）
	//
	// 底层数组：分配了一个长度为5的数组 [0, 0, 0, 0, 0]
	// 切片窗口：能看到全部5个元素
	a := make([]int, 5)
	printSlice("a", a) // 输出: a len=5 cap=5 [0 0 0 0 0]

	// ========== 示例2：make([]T, length, capacity) ==========
	// 【Go vs Java】指定容量
	// Java:  ArrayList<Integer> b = new ArrayList<>(5);  // 初始容量5，但size=0
	// Go:    b := make([]int, 0, 5)
	//
	// make([]int, 0, 5) 的含义：
	// - 创建一个 int 类型的切片
	// - 长度（len）= 0（当前没有元素）
	// - 容量（cap）= 5（底层数组可以容纳5个元素）
	//
	// 底层数组：分配了一个长度为5的数组 [0, 0, 0, 0, 0]
	// 切片窗口：当前长度为0（看不到任何元素），但容量是5
	//
	// 用途：当你知道大概需要多少个元素，但暂时不添加时，可以先预留容量
	// 这样可以避免频繁扩容（类似 Java 的 ArrayList 预分配容量）
	b := make([]int, 0, 5)
	printSlice("b", b) // 输出: b len=0 cap=5 []

	// ========== 示例3：从空切片扩展 ==========
	// 【Go vs Java】扩展切片
	// Java:  b.add(0); b.add(0);  // 添加元素
	// Go:    c := b[:2]  // 扩展窗口，从底层数组中"看到"前2个元素
	//
	// 执行 c = b[:2] 后：
	// - 底层数组：[0, 0, 0, 0, 0]（和 b 共享同一个底层数组）
	// - c 的窗口：[0, 0]（能看到前2个元素）
	// - len(c) = 2, cap(c) = 5（容量还是5，因为从索引0开始）
	//
	// 关键点：虽然 b 的长度是0，但底层数组已分配，所以可以扩展窗口看到数据
	c := b[:2]
	printSlice("c", c) // 输出: c len=2 cap=5 [0 0]

	// ========== 示例4：从中间位置切片 ==========
	// 【Go vs Java】从指定位置切片
	// Java:  int[] d = Arrays.copyOfRange(c, 2, 5);
	// Go:    d := c[2:5]
	//     ^^
	// 执行 d = c[2:5] 后：
	// - 底层数组：[0, 0, 0, 0, 0]（和 c、b 共享同一个底层数组）
	// - d 的窗口：[0, 0, 0]（从索引2到索引5，能看到索引2、3、4的元素）
	// - len(d) = 3, cap(d) = 3（容量变成3，因为从索引2开始，后面只剩3个位置）
	//
	// 注意：虽然 c 的长度是2（只能看到索引0、1），但可以切到索引2-5
	// 这是因为切片是"窗口"的概念，只要底层数组容量足够就可以
	d := c[2:5]
	printSlice("d", d) // 输出: d len=3 cap=3 [0 0 0]
}

// printSlice 打印切片的名称、长度、容量和内容
// 【Go vs Java】格式化输出
// Java:  System.out.printf("%s len=%d cap=%d %s\n", name, arr.length, arr.capacity, Arrays.toString(arr));
// Go:    fmt.Printf("%s len=%d cap=%d %v\n", name, len(x), cap(x), x)
func printSlice(s string, x []int) {
	// %s - 字符串（切片名称）
	// len(x) - 切片长度（当前元素个数）
	// cap(x) - 切片容量（从起始位置到底层数组末尾的元素个数）
	// %v - 通用格式化，打印切片内容
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

// ========== make() 函数总结 ==========
// make([]T, length)        - 长度和容量都等于 length
// make([]T, length, cap)   - 长度=length，容量=cap（cap >= length）
//
// make() 的优势：
// 1. 预分配底层数组，避免频繁扩容
// 2. 可以创建长度和容量不同的切片（长度=0，容量>0）
// 3. 所有元素自动初始化为类型的零值
//
// 使用场景：
// - make([]int, 5)      - 需要立即使用5个元素的切片
// - make([]int, 0, 100) - 需要预留容量，但暂时不需要元素（性能优化）
