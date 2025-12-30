package main

import "fmt"

/*
映射字面量
映射的字面量和结构体类似，只不过必须有键名。
*/
// ========== 定义结构体类型 ==========
// 【Go vs Java】结构体定义
// Java:  class Vertex { double lat; double lon; }
// Go:    type Vertex struct { Lat, Long float64 }
type Vertex struct {
	Lat, Long float64 // 纬度（Latitude）和经度（Longitude）
}

// ========== Map 字面量初始化 ==========
// 【Go vs Java】Map 字面量
// Java:  Map<String, Vertex> m = Map.of(
//
//	    "Bell Labs", new Vertex(40.68433, -74.39967),
//	    "Google", new Vertex(37.42202, -122.08408)
//	);
//	或者：
//	Map<String, Vertex> m = new HashMap<>();
//	m.put("Bell Labs", new Vertex(...));
//	m.put("Google", new Vertex(...));
//
// Go:    var m = map[string]Vertex{ ... }
//
// Map 字面量语法说明：
// - map[KeyType]ValueType{ key1: value1, key2: value2, ... }
// - 可以一次性初始化多个键值对
// - 每个键值对用逗号分隔
// - 最后一行必须有逗号（即使是最后一项）
//
// 语法格式：
//
//	var m = map[string]Vertex{
//	    "键1": 值1,
//	    "键2": 值2,
//	}
//
// 注意：map 字面量会自动创建 map，不需要 make()
var m = map[string]Vertex{
	// 第一个键值对："Bell Labs" → Vertex{40.68433, -74.39967}
	"Bell Labs": Vertex{
		40.68433, -74.39967, // 贝尔实验室的坐标（纬度，经度）
	},
	// 第二个键值对："Google" → Vertex{37.42202, -122.08408}
	"Google": Vertex{
		37.42202, -122.08408, // 谷歌总部的坐标（纬度，经度）
	},
	// 注意：最后一行必须有逗号（即使只有两个元素）
}

func main() {
	// ========== 打印整个 map ==========
	// 【Go vs Java】打印 Map
	// Java:  System.out.println(m);  // 输出格式不友好
	// Go:    fmt.Println(m)  // 输出格式清晰：map[key:value ...]
	//
	// 输出格式：
	// map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
	//
	// 说明：
	// - map 的输出格式是 map[key1:value1 key2:value2]
	// - 结构体的输出格式是 {field1 field2}
	// - 组合起来就是：map[key:{field1 field2} ...]
	fmt.Println(m)
	// 输出：map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
}

// ========== Map 字面量的简化写法 ==========
// 如果值的类型已经在 map 类型中指定，可以省略类型名：
//
// 完整写法：
// var m = map[string]Vertex{
//     "Bell Labs": Vertex{40.68433, -74.39967},
//     "Google": Vertex{37.42202, -122.08408},
// }
//
// 简化写法（推荐）：
// var m = map[string]Vertex{
//     "Bell Labs": {40.68433, -74.39967},  // 省略 Vertex
//     "Google": {37.42202, -122.08408},    // 省略 Vertex
// }
//
// 编译器会自动推断值的类型为 Vertex

// ========== Map 字面量 vs make() 对比 ==========
// 方式1：字面量（适合初始化时就知道所有键值对）
//   var m = map[string]Vertex{ "key": value, ... }
//
// 方式2：make() + 逐个添加（适合动态添加）
//   m := make(map[string]Vertex)
//   m["key"] = value
//
// 选择建议：
// - 初始化时已知所有数据 → 使用字面量（简洁）
// - 需要动态添加 → 使用 make()（灵活）
