package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	// net.Listen 在指定端口上创建一个 TCP 监听器
	// 第一个参数 "tcp" 表示使用 TCP 协议
	// 第二个参数 ":8090" 表示监听所有网络接口的 8090 端口
	// 返回的 listener 是一个 net.Listener 接口，用于接受客户端连接
	listener, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal("Error listening:", err)
	}

	// 确保程序退出时关闭监听器，释放端口资源
	defer listener.Close()

	// ===== 主循环：持续等待并接受客户端连接 =====
	// Accept() 会阻塞，直到有新的客户端连接进来
	// 每次有新连接，就启动一个 goroutine 并发处理
	// 这是 Go 实现高并发 TCP 服务器的核心模式
	for {

		// Accept 接受一个新连接，返回一个 net.Conn 连接对象
		// 每个 Conn 代表一条客户端与服务器之间的 TCP 连接
		conn, err := listener.Accept()
		if err != nil {
			// 接受连接失败时打印错误，但不终止服务器，继续等待下一个连接
			log.Println("Error accepting conn:", err)
			continue
		}

		// go 关键字启动一个新的 goroutine 来处理该连接
		// 这样主循环可以立即回到 Accept() 等待下一个连接
		// 多个客户端可以同时连接，每个连接独立处理，互不阻塞
		go handleConnection(conn)
	}
}

// handleConnection 处理单个客户端连接
// 参数 conn 是客户端与服务器之间的一条 TCP 连接
func handleConnection(conn net.Conn) {

	// 函数结束时关闭连接，释放资源
	defer conn.Close()

	// 用 bufio.NewReader 包装连接，提供缓冲读取能力
	// 直接从 conn 读取也可以，但 bufio 提供了更方便的按行读取
	reader := bufio.NewReader(conn)

	// ReadString('\n') 从连接中读取数据，直到遇到换行符 '\n'
	// 这意味着客户端发送的每条消息必须以换行符结尾
	// 如果客户端断开连接而不发送数据，会返回 EOF 错误
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Read error: %v", err)
		return
	}

	// 处理收到的消息：
	//   1. TrimSpace 去掉末尾的 \n 和可能的空格
	//   2. ToUpper 将消息转为大写
	ackMsg := strings.ToUpper(strings.TrimSpace(message))

	// 构造响应消息，格式为 "ACK: 大写消息\n"
	response := fmt.Sprintf("ACK: %s\n", ackMsg)

	// conn.Write 将响应数据通过 TCP 连接发送给客户端
	// Write 接收 []byte 类型，所以用 []byte() 转换字符串
	_, err = conn.Write([]byte(response))
	if err != nil {
		log.Printf("Server write error: %v", err)
	}
}
