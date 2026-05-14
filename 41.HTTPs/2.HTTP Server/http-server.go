package main

import (
	"fmt"
	"net/http"
)

// hello 是一个 HTTP 处理函数（Handler）
// 参数说明：
//
//	w     http.ResponseWriter —— 用于向客户端写入响应内容
//	req   *http.Request       —— 包含客户端请求的所有信息（URL、方法、头部、参数等）
//
// 当用户访问 /hello 时，向客户端返回 "hello\n" 这段文本
func hello(w http.ResponseWriter, req *http.Request) {

	// Fprintf 将格式化的字符串写入 w（即 HTTP 响应体）
	// 客户端（浏览器或 curl）会收到这段纯文本
	fmt.Fprintf(w, "hello\n")
}

// headers 是另一个 HTTP 处理函数
// 当用户访问 /headers 时，将请求中的所有 HTTP 头部信息返回给客户端
func headers(w http.ResponseWriter, req *http.Request) {

	// req.Header 是一个 map[string][]string 类型
	//   key   —— 头部名称，如 "Content-Type"、"User-Agent"
	//   value —— 字符串切片，因为同一个头部可以出现多次
	//
	// 双层循环：外层遍历每个头部名称，内层遍历该名称下的每个值
	for name, headers := range req.Header {
		for _, h := range headers {
			// 以 "名称: 值" 的格式写入响应
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	// http.HandleFunc 注册路由：将 URL 路径与处理函数绑定
	// "/hello"   → 当请求路径匹配 /hello 时，调用 hello 函数
	// "/headers" → 当请求路径匹配 /headers 时，调用 headers 函数
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	// http.ListenAndServe 启动 HTTP 服务器
	// 第一个参数 ":8090" —— 监听本机所有网卡的 8090 端口
	// 第二个参数 nil      —— 使用 DefaultServeMux（即上面 HandleFunc 注册的默认路由器）
	// 该函数会阻塞运行，持续监听请求直到程序被终止
	// 如果启动失败（如端口被占用），返回错误；这里忽略了返回值
	http.ListenAndServe(":8090", nil)
}
