// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 16.
//!+

// Fetch prints the content found at each specified URL.
package main

import (
	"fmt"
	"io/ioutil"
	"net/http" // HTTP客户端和服务器包
	"os"
)

// main 获取URL内容并打印到标准输出
func main() {
	// 遍历所有URL参数
	for _, url := range os.Args[1:] {
		// 【Go vs Java】HTTP GET请求
		// Java:  HttpClient client = HttpClient.newHttpClient();
		//        HttpResponse<String> resp = client.send(request, ...);
		// Go:    resp, err := http.Get(url)
		// 注意：Go的http.Get非常简洁，一行代码搞定
		resp, err := http.Get(url)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			// 【Go vs Java】退出程序
			// Java:  System.exit(1);
			// Go:    os.Exit(1)
			// 注意：0表示成功，非0表示错误
			os.Exit(1)
		}

		// 【Go vs Java】读取响应体
		// Java:  String body = resp.body();
		// Go:    b, err := ioutil.ReadAll(resp.Body)
		// 注意：resp.Body是io.Reader接口，需要用ReadAll读取全部内容
		b, err := ioutil.ReadAll(resp.Body)

		// 【Go vs Java】关闭资源
		// Java:  try-with-resources 自动关闭
		// Go:    必须手动调用 Close()
		// 注意：应该用 defer resp.Body.Close() 确保一定会关闭
		resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}

		// 【Go vs Java】打印字节数组
		// Java:  System.out.print(new String(bytes));
		// Go:    fmt.Printf("%s", b)
		// 注意：%s可以直接格式化[]byte为字符串
		fmt.Printf("%s", b)
	}
}

//!-
