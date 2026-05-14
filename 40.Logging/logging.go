package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"log/slog"
)

func main() {

	// ===== 标准 log 包 =====

	// 使用默认的 logger 输出一条消息
	// 默认格式：日期 时间 + 消息内容
	// 例如：2024/01/15 10:30:00 standard logger
	log.Println("standard logger")

	// SetFlags 修改日志输出的格式标志
	// log.LstdFlags   = log.Ldate | log.Ltime（默认的日期+时间）
	// log.Lmicroseconds = 在时间中追加微秒精度
	// 两者用 | 组合，同时启用
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Println("with micro")

	// log.Lshortfile = 在日志中显示调用处的文件名和行号
	// 例如：main.go:20: with file/line
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("with file/line")

	// ===== 自定义 Logger =====

	// log.New 创建一个新的 Logger 实例
	// 参数说明：
	//   os.Stdout    —— 输出目标，这里是标准输出
	//   "my:"        —— 每条日志的前缀
	//   log.LstdFlags —— 日志格式标志
	mylog := log.New(os.Stdout, "my:", log.LstdFlags)
	mylog.Println("from mylog")

	// SetPrefix 动态修改日志前缀
	mylog.SetPrefix("ohmy:")
	mylog.Println("from mylog")

	// ===== 输出到 Buffer 的 Logger =====

	// 创建一个内存缓冲区，日志将写入这里而非终端
	var buf bytes.Buffer

	// 将日志输出目标设为 buf（实现了 io.Writer 接口）
	// 日志内容会暂存在内存中，不会显示在屏幕上
	buflog := log.New(&buf, "buf:", log.LstdFlags)

	// 写入一条日志到缓冲区（屏幕上看不到）
	buflog.Println("hello")

	// 手动从缓冲区读取并打印日志内容
	// 输出：buf:2024/01/15 10:30:00 hello
	fmt.Print("from buflog:", buf.String())

	// ===== slog：结构化日志（Go 1.21+）=====

	// slog.NewJSONHandler 创建一个 JSON 格式的日志处理器
	// os.Stderr —— 输出到标准错误流（与标准输出分离，便于运维）
	// nil       —— 使用默认选项（如最低日志级别为 Info）
	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)

	// 用处理器创建一个结构化日志实例
	myslog := slog.New(jsonHandler)

	// Info 级别的结构化日志
	// 输出为 JSON 格式：
	// {"time":"2024-01-15T10:30:00Z","level":"INFO","msg":"hi there"}
	myslog.Info("hi there")

	// 带键值对的结构化日志
	// 键值对以成对方式传入："key", "val", "age", 25
	// 输出：
	// {"time":"2024-01-15T10:30:00Z","level":"INFO","msg":"hello again","key":"val","age":25}
	myslog.Info("hello again", "key", "val", "age", 25, "addr", "贵州")
}
