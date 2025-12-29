// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 20.
//!+

// Server2 is a minimal "echo" and counter server.
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync" // 同步原语包（互斥锁、等待组等）
)

// 【Go vs Java】包级变量和互斥锁
// Java:  private static final Object lock = new Object();
//
//	private static int count = 0;
//
// Go:    var mu sync.Mutex
//
//	var count int
//
// 注意：Go没有private/public关键字，首字母小写表示包内可见
var mu sync.Mutex
var count int

// main 启动一个带计数器的HTTP服务器
func main() {
	// 【Go vs Java】注册多个路由
	// Java:  server.createContext("/", handler);
	//        server.createContext("/count", counter);
	// Go:    http.HandleFunc("/", handler)
	//        http.HandleFunc("/count", counter)
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Println("服务器启动在 http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler echoes the Path component of the requested URL.
func handler(w http.ResponseWriter, r *http.Request) {
	// 【Go vs Java】互斥锁
	// Java:  synchronized(lock) { count++; }
	// Go:    mu.Lock(); count++; mu.Unlock()
	// 注意：Go的锁是显式的，必须手动Lock和Unlock
	mu.Lock()
	count++
	mu.Unlock()
	// 注意：这里应该用 defer mu.Unlock() 确保一定会解锁

	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// counter echoes the number of calls so far.
func counter(w http.ResponseWriter, r *http.Request) {
	// 【Go vs Java】读取共享变量也需要加锁
	// Java:  synchronized(lock) { return count; }
	// Go:    mu.Lock(); ... mu.Unlock()
	// 注意：Go的HTTP服务器每个请求在独立的goroutine中处理，必须同步
	mu.Lock()
	fmt.Fprintf(w, "Count %d\n", count)
	mu.Unlock()
}

//!-
