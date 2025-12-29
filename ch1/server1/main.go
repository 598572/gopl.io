// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 19.
//!+

// Server1 is a minimal "echo" server.
package main

import (
	"fmt"
	"log" // 日志包
	"net/http"
)

// main 启动一个简单的HTTP服务器
// 【Go vs Java】HTTP服务器对比
// Java:  HttpServer server = HttpServer.create(new InetSocketAddress(8000), 0);
//
//	server.createContext("/", handler);
//	server.start();
//
// Go:    http.HandleFunc("/", handler)
//
//	http.ListenAndServe("localhost:8000", nil)
//
// 注意：Go的HTTP服务器只需2行代码，非常简洁
func main() {
	// 【Go vs Java】注册路由处理器
	// Java:  server.createContext("/", handler);
	// Go:    http.HandleFunc("/", handler)
	// 注意：handler是一个函数，Go的函数是一等公民
	http.HandleFunc("/", handler) // each request calls handler

	// 【Go vs Java】启动服务器并监听
	// Java:  server.start();
	// Go:    log.Fatal(http.ListenAndServe("localhost:8000", nil))
	// 注意：ListenAndServe会阻塞，返回错误时log.Fatal退出程序
	log.Println("Server started on http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
// 【Go vs Java】HTTP处理器签名
// Java:  void handle(HttpExchange exchange)
// Go:    func handler(w http.ResponseWriter, r *http.Request)
// 注意：w用于写响应，r包含请求信息
func handler(w http.ResponseWriter, r *http.Request) {
	// 【Go vs Java】写HTTP响应
	// Java:  exchange.sendResponseHeaders(200, response.length());
	//        OutputStream os = exchange.getResponseBody();
	//        os.write(response.getBytes());
	// Go:    fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	// 注意：%q 格式化为带引号的字符串，w实现了io.Writer接口
	log.Println("请求路径：", r.URL.Path) // 打印请求路径
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

//!-
