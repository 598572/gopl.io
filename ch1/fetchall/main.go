// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time" // 时间处理包
)

// main 并发获取多个URL，统计时间和大小
// 【Go vs Java】并发模型对比
// Java:  ExecutorService + Future 或 CompletableFuture
// Go:    goroutine + channel（更轻量、更简洁）
func main() {
	// 【Go vs Java】记录开始时间
	// Java:  long start = System.currentTimeMillis();
	// Go:    start := time.Now()
	start := time.Now()

	// 【Go vs Java】创建channel
	// Java:  BlockingQueue<String> queue = new LinkedBlockingQueue<>();
	// Go:    ch := make(chan string)
	// 注意：channel是Go的核心并发原语，用于goroutine间通信
	ch := make(chan string)

	// 为每个URL启动一个goroutine
	for _, url := range os.Args[1:] {
		// 【Go vs Java】启动并发任务
		// Java:  executor.submit(() -> fetch(url, queue));
		// Go:    go fetch(url, ch)
		// 注意：go关键字启动goroutine，比Java线程轻量得多（几KB栈空间）
		go fetch(url, ch) // start a goroutine
	}

	// 等待所有goroutine完成
	// 【Go vs Java】从channel接收数据
	// Java:  for (int i = 0; i < urls.size(); i++) { queue.take(); }
	// Go:    for range os.Args[1:] { <-ch }
	// 注意：<-ch 从channel接收数据，会阻塞直到有数据
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}

	// 【Go vs Java】计算耗时
	// Java:  double elapsed = (System.currentTimeMillis() - start) / 1000.0;
	// Go:    time.Since(start).Seconds()
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

// fetch 获取URL内容并通过channel发送结果
// 【Go vs Java】channel方向
// Java:  无对应概念
// Go:    chan<- string 表示只能发送的channel（类型安全）
// 注意：chan<- 是send-only，<-chan 是receive-only
func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		// 【Go vs Java】发送数据到channel
		// Java:  queue.put(err.toString());
		// Go:    ch <- fmt.Sprint(err)
		// 注意：<- 是发送操作符，格式为 channel <- value
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	// 【Go vs Java】复制数据流
	// Java:  long nbytes = resp.body().transferTo(OutputStream.nullOutputStream());
	// Go:    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	// 注意：ioutil.Discard 类似 /dev/null，丢弃所有写入的数据
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources

	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	secs := time.Since(start).Seconds()

	// 【Go vs Java】格式化字符串
	// Java:  String.format("%.2fs  %7d  %s", secs, nbytes, url)
	// Go:    fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
	// 注意：%7d 表示右对齐，宽度为7
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
