package main

import (
	"fmt"
	"os"
)

func main() {

	// 注册一个 defer 语句，正常情况下会在 main 函数返回前执行
	defer fmt.Println("!")

	// os.Exit 立即终止当前程序，以指定的退出码（3）退出
	//
	// ⚠️ 关键行为：
	//   os.Exit 不会触发 defer 语句的执行
	//   程序直接终止，上面的 defer fmt.Println("!") 会被跳过
	//   "!" 永远不会被打印
	//
	// 退出码（exit code）约定：
	//   0 —— 程序正常退出
	//   非 0 —— 程序异常退出，具体数值由程序自定义
	//   3 在这里只是一个示例值，没有特殊含义
	os.Exit(3)

	// ⚠️ os.Exit 之后的代码永远不会执行
	fmt.Println("这行永远不会打印")
}
