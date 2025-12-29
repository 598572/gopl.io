package main

import "fmt"

/*
*
map 映射
map 映射将键映射到值。

映射的零值为 nil 。nil 映射既没有键，也不能添加键。

make 函数会返回给定类型的映射，并将其初始化备用。
*/
// ========== 定义结构体类型 ==========
// 【Go vs Java】结构体定义
// Java:  class Vertex {
//            double lat;
//            double lon;
//            Vertex(double lat, double lon) { ... }
//        }
// Go:    type Vertex struct { Lat, Long float64 }
//
// 解释：
// - type Vertex struct 定义一个名为 Vertex 的结构体类型
// - Lat, Long float64 表示两个 float64 类型的字段
// - 这是紧凑写法，等价于：
//     type Vertex struct {
//         Lat  float64
//         Long float64
//     }
// - 字段名首字母大写表示导出（public），可以被外部包访问
type Vertex struct {
	Lat, Long float64 // 纬度（Latitude）和经度（Longitude）
}

// ========== 声明 map 变量 ==========
// 【Go vs Java】Map/字典声明
// Java:  Map<String, Vertex> m;
// Go:    var m map[string]Vertex
//
// 解释：
// - map[string]Vertex 表示一个 map 类型
//   - string 是键（key）的类型
//   - Vertex 是值（value）的类型
//
// - 此时 m 是 nil（零值），还不能使用
// - nil map 不能添加键值对，必须先初始化
var m map[string]Vertex

func main() {
	// ========== 使用 make 创建 map ==========
	// 【Go vs Java】创建 Map
	// Java:  Map<String, Vertex> m = new HashMap<>();
	// Go:    m = make(map[string]Vertex)
	//
	// make(map[string]Vertex) 的含义：
	// - 创建一个 map[string]Vertex 类型的 map
	// - 返回一个已初始化可用的 map（不是 nil）
	// - 现在可以向 map 中添加键值对了
	//
	// 注意：map 的零值是 nil，nil map 不能使用
	// - var m map[string]Vertex  → m 是 nil，不能使用
	// - m = make(map[string]Vertex) → m 是可用的空 map
	m = make(map[string]Vertex)

	// ========== 向 map 添加元素 ==========
	// 【Go vs Java】添加键值对
	// Java:  m.put("Bell Labs", new Vertex(40.68433, -74.39967));
	// Go:    m["Bell Labs"] = Vertex{40.68433, -74.39967}
	//
	// 语法说明：
	// - m[key] = value 向 map 添加或更新键值对
	// - "Bell Labs" 是键（string 类型）
	// - Vertex{40.68433, -74.39967} 是值（Vertex 结构体）
	//
	// 结构体字面量：
	// - Vertex{40.68433, -74.39967} 创建 Vertex 结构体实例
	// - 按字段定义的顺序赋值：40.68433 → Lat, -74.39967 → Long
	// - 也可以使用字段名：Vertex{Lat: 40.68433, Long: -74.39967}
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967, // 贝尔实验室的坐标（纬度，经度）
	}

	// ========== 从 map 读取元素 ==========
	// 【Go vs Java】读取值
	// Java:  Vertex v = m.get("Bell Labs");
	//        if (v == null) { ... }  // Java 返回 null 如果键不存在
	// Go:    v := m["Bell Labs"]
	//        if v, ok := m["Bell Labs"]; ok { ... }  // Go 可以检查键是否存在
	//
	// 语法说明：
	// - m[key] 从 map 中读取值
	// - 如果键存在，返回对应的值
	// - 如果键不存在，返回值类型的零值（这里是 Vertex 的零值）
	//
	// 注意：这里假设 "Bell Labs" 键存在，所以直接读取
	fmt.Println(m["Bell Labs"])
	// 输出：{40.68433 -74.39967}（结构体的默认打印格式）
}

// ========== Map 的核心概念总结 ==========
// 1. map 的零值是 nil，不能直接使用
// 2. 必须用 make() 创建 map 才能使用
// 3. map[key] = value 添加或更新键值对
// 4. value := map[key] 读取值（如果键不存在，返回零值）
// 5. value, ok := map[key] 可以检查键是否存在（ok 是 bool）
//
// 与 Java Map 的对比：
// - Java: 需要 new HashMap<>() 创建，get() 返回 null 如果不存在
// - Go:   需要 make(map[K]V) 创建，访问不存在的键返回零值（更安全）
