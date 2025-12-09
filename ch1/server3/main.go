// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 21.

// Server3 is an "echo" server that displays request parameters.
package main

import (
	"fmt"
	"log"
	"net/http"
)

// main 启动一个显示请求详细信息的HTTP服务器
func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// !+handler
// handler echoes the HTTP request.
func handler(w http.ResponseWriter, r *http.Request) {
	// 【Go vs Java】打印请求行
	// Java:  String requestLine = request.getMethod() + " " + request.getRequestURI();
	// Go:    fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	// 注意：r.Method是字符串（GET/POST等），r.URL是*url.URL类型
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)

	// 【Go vs Java】遍历请求头
	// Java:  for (Map.Entry<String, List<String>> entry : headers.entrySet())
	// Go:    for k, v := range r.Header
	// 注意：r.Header是map[string][]string类型，一个key可以有多个value
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}

	// 【Go vs Java】获取Host和客户端地址
	// Java:  String host = request.getHeader("Host");
	//        String remoteAddr = request.getRemoteAddr();
	// Go:    r.Host, r.RemoteAddr (直接访问字段)
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)

	// 【Go vs Java】解析表单参数
	// Java:  Map<String, String[]> params = request.getParameterMap();
	// Go:    r.ParseForm(); r.Form
	// 注意：必须先调用ParseForm()，然后才能访问r.Form
	if err := r.ParseForm(); err != nil {
		// 【Go vs Java】if语句中的短变量声明
		// Java:  Error err = r.parseForm(); if (err != null) {...}
		// Go:    if err := r.ParseForm(); err != nil {...}
		// 注意：err的作用域仅在if块内
		log.Print(err)
	}

	// 【Go vs Java】遍历表单参数
	// Java:  for (Map.Entry<String, String[]> entry : params.entrySet())
	// Go:    for k, v := range r.Form
	// 注意：r.Form也是map[string][]string类型
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

//!-handler
