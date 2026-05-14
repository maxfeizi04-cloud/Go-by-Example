package main

import (
	"flag"
	"fmt"
)

func main() {

	// 定义一个名为 "word" 的字符串命令行标志
	// 参数说明：
	//   "word"  —— 命令行中使用的标志名称（-word）
	//   "foo"   —— 默认值，用户未指定时使用
	//   "a string" —— 帮助说明文字，通过 -h 查看
	// 返回值是一个 *string 指针，需要通过 *wordPtr 解引用获取实际值
	wordPtr := flag.String("word", "foo", "a string")

	// 定义整数标志 -numb，默认值 42
	numbPtr := flag.Int("numb", 42, "an int")

	// 定义布尔标志 -fork，默认值 false
	// 布尔标志可以不传值：-fork 即表示 true
	forkPtr := flag.Bool("fork", false, "a bool")

	// 另一种定义方式：将标志绑定到已有的变量上
	// 参数顺序：绑定的变量指针、标志名、默认值、帮助文字
	// 这种方式不需要解引用，直接使用 svar 变量即可
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// 解析命令行参数
	// 必须在所有标志定义之后调用，否则无法识别标志
	// 解析后，flag.String/Int/Bool 返回的指针才会指向实际传入的值
	flag.Parse()

	// 输出每个标志解析后的值
	// 注意前三个需要通过 * 解引用指针
	fmt.Println("word:", *wordPtr)
	fmt.Println("numb:", *numbPtr)
	fmt.Println("fork:", *forkPtr)
	fmt.Println("svar:", svar)

	// flag.Args() 返回标志解析完之后剩余的非标志参数（即 tail）
	// 例如 go run main.go -word=hello extra1 extra2 中，tail 为 [extra1 extra2]
	fmt.Println("tail:", flag.Args())
}
