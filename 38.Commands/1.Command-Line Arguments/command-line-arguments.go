package main

import (
	"fmt"
	"os"
)

func main() {

	// os.Args 是一个字符串切片 (string slice)
	// 包含程序启动时传入的所有命令行参数
	// os.Args[0] 是程序自身的路径/名称，后续元素才是真正的参数
	argsWithProg := os.Args

	// 通过切片操作 [1:] 跳过第一个元素（程序名称）
	// 只保留用户传入的实际参数
	argsWithoutProg := os.Args[1:]

	// 取索引为 3 的参数（第 4 个元素，因为索引从 0 开始）
	// 注意：如果传入的参数不足 4 个，这里会触发 panic: index out of range
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
