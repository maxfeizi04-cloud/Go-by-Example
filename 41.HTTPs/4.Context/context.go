package main

import (
	"fmt"
	"net/http"
	"time"
)

// hello 是一个 HTTP 处理函数
// 演示了 Go 的 context（上下文）机制：当客户端断开连接时，服务器可以感知并中止处理
func hello(w http.ResponseWriter, req *http.Request) {

	// req.Context() 获取与当前请求关联的 context
	// context 会在以下情况被取消（cancel）：
	//   1. 客户端关闭了连接（如关闭浏览器、Ctrl+C 中断 curl）
	//   2. 请求超时（如果设置了超时中间件）
	//   3. 服务器主动调用 cancel()
	ctx := req.Context()

	fmt.Println("server: hello handler started")
	defer fmt.Println("server: hello handler ended")

	// select 同时监听两个 channel，哪个先就绪就执行哪个分支
	select {

	// 分支 1：正常处理路径
	// time.After(10 * time.Second) 返回一个 channel
	// 10 秒后该 channel 会收到一个时间值
	// 如果 10 秒内客户端没有断开，就执行这个分支，返回响应
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello\n")

	// 分支 2：客户端断开路径
	// ctx.Done() 返回一个 channel
	// 当 context 被取消时（如客户端断开），该 channel 会被关闭
	// 关闭的 channel 可以立即读取，所以 select 会进入这个分支
	case <-ctx.Done():

		// ctx.Err() 返回 context 被取消的原因：
		//   context.Canceled  —— 客户端主动断开
		//   context.DeadlineExceeded —— 超时
		err := ctx.Err()
		fmt.Println("server:", err)

		// 向客户端返回 500 Internal Server Error
		// 但实际上客户端已经断开，这个响应可能发不出去
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8090", nil)
}
