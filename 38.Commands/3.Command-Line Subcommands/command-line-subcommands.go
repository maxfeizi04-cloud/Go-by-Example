package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	// ===== 子命令 "foo" 的标志集 =====

	// flag.NewFlagSet 创建一个独立的标志集合，用于实现子命令
	// 第一个参数 "foo" 是该 FlagSet 的名称，用于错误信息中显示
	// 第二个参数 flag.ExitOnError 表示解析出错时自动调用 os.Exit(2) 退出程序
	// 与全局 flag 包分离，每个子命令拥有自己独立的一组标志
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)

	// 在 foo 子命令下定义布尔标志 -enable，默认值 false
	fooEnable := fooCmd.Bool("enable", false, "enable")

	// 在 foo 子命令下定义字符串标志 -name，默认值为空字符串
	fooName := fooCmd.String("name", "", "name")

	// ===== 子命令 "bar" 的标志集 =====

	// 为 bar 子命令创建另一套独立的标志集合
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)

	// 在 bar 子命令下定义整数标志 -level，默认值 0
	barLevel := barCmd.Int("level", 0, "level")

	// ===== 参数校验 =====

	// os.Args 至少需要 2 个元素：程序名 + 子命令名
	// 如果只输入了程序名（len < 2），说明用户没有指定子命令
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// ===== 根据子命令名称分发处理 =====

	// os.Args[1] 是用户输入的子命令名称（如 "foo" 或 "bar"）
	// os.Args[2:] 是子命令后面的所有参数（包含子命令自己的标志和非标志参数）
	switch os.Args[1] {

	case "foo":
		// 解析 foo 子命令的参数，从 os.Args[2:] 开始
		// 例如：prog foo -enable -name=gopher tail1
		//       解析 -enable 和 -name=gopher，tail1 留在 Args() 中
		fooCmd.Parse(os.Args[2:])

		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)
		fmt.Println("  name:", *fooName)
		// fooCmd.Args() 返回解析后剩余的非标志参数（尾巴参数）
		fmt.Println("  tail:", fooCmd.Args())

	case "bar":
		// 解析 bar 子命令的参数
		// 例如：prog bar -level=5 tail1 tail2
		barCmd.Parse(os.Args[2:])

		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)
		fmt.Println("  tail:", barCmd.Args())

	default:
		// 输入的子命令既不是 foo 也不是 bar，报错退出
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
