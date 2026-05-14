package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8090")
	if err != nil {
		fmt.Println("连接失败:", err)
		return
	}
	defer conn.Close()

	// 从终端读取用户输入
	fmt.Print("请输入消息: ")
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')

	// 发送给服务器
	conn.Write([]byte(input))

	// 读取服务器响应
	response, _ := bufio.NewReader(conn).ReadString('\n')
	fmt.Println("服务器回复:", response)
}
