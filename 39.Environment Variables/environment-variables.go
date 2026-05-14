package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// os.Setenv 设置环境变量
	// 第一个参数是变量名，第二个参数是变量值
	// 设置后当前进程及其子进程都可以读取到该变量
	os.Setenv("FOO", "1")

	// os.Getenv 读取环境变量的值
	// 如果变量存在，返回其值；如果不存在，返回空字符串 ""
	// FOO 刚刚被设置为 "1"，所以这里输出 "1"
	fmt.Println("FOO:", os.Getenv("FOO"))

	// BAR 从未被设置，GetEnv 返回空字符串 ""
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()

	// os.Environ() 返回当前进程中所有环境变量的切片
	// 每个元素格式为 "KEY=VALUE" 的字符串
	// 例如 "PATH=C:\Go\bin;C:\Windows;..."
	for _, e := range os.Environ() {
		// strings.SplitN 按 "=" 分割，最多分成 2 部分
		// 第二个参数 2 表示最多切分一次，防止 VALUE 中含有 "=" 被错误分割
		// 例如 "PATH=C:\Go=bin" 只会在第一个 "=" 处切分
		//   pair[0] = "PATH"
		//   pair[1] = "C:\Go=bin"
		pair := strings.SplitN(e, "=", 2)

		// 只打印环境变量的名称（不打印值）
		fmt.Println(pair[0])
	}
}
