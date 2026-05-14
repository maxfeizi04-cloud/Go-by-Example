package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {

	// http.Get 发送一个 HTTP GET 请求到指定 URL
	// 返回值：
	//   resp —— *http.Response 结构体，包含状态码、响应头、响应体等
	//   err  —— 如果请求失败（如网络不通、DNS 解析失败），err 不为 nil
	// 注意：resp.Body 实现了 io.ReadCloser 接口，必须手动关闭
	resp, err := http.Get("https://baidu.com")
	if err != nil {
		panic(err)
	}

	// defer 确保在 main 函数结束时关闭响应体
	// 不关闭会导致连接无法被复用，造成连接泄漏
	defer resp.Body.Close()

	// resp.Status 是状态码的文本表示，例如 "200 OK"
	fmt.Println("Response status:", resp.Status)

	// bufio.NewScanner 创建一个文本扫描器，逐行读取 resp.Body
	// resp.Body 是一个 io.Reader，Scanner 会按行（\n）拆分数据流
	scanner := bufio.NewScanner(resp.Body)

	// scanner.Scan() 每次调用读取一行
	// && i < 5：限制只读取前 5 行后停止
	// scanner.Text() 返回当前行的文本内容（不含换行符）
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}

	// 检查扫描过程中是否发生了错误（如网络中断、数据读取失败）
	// 正常结束（EOF）不算错误，Scan() 返回 false 且 Err() 返回 nil
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
