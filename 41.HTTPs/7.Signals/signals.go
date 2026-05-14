package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
)

func main() {

	// signal.NotifyContext 创建一个可被信号取消的 context
	// 参数说明：
	//   context.Background() —— 父 context，作为整个 context 树的根
	//   syscall.SIGINT       —— 用户按 Ctrl+C 时发送的中断信号
	//   syscall.SIGTERM      —— kill 命令默认发送的终止信号（优雅关闭）
	//
	// 返回值：
	//   ctx  —— 当收到上述信号时，ctx 会被自动取消（Done channel 关闭）
	//   stop —— 取消信号监听的函数，调用后不再监听信号
	//           即使不主动调用，defer stop() 也会在 main 结束时执行
	ctx, stop := signal.NotifyContext(
		context.Background(), syscall.SIGINT, syscall.SIGTERM)

	// defer stop() 确保函数退出时停止信号监听，释放相关资源
	// 即使程序是因为收到信号而退出，也需要调用 stop 清理注册
	defer stop()

	fmt.Println("awaiting signal")

	// <-ctx.Done() 阻塞当前 goroutine，等待 context 被取消
	// 只有收到 SIGINT 或 SIGTERM 信号时，ctx.Done() 的 channel 才会关闭
	// channel 关闭后，<- 操作立即返回，程序继续执行下面的代码
	<-ctx.Done()

	fmt.Println()

	// context.Cause(ctx) 返回 context 被取消的原因
	// 对于 signal.NotifyContext，返回的是触发取消的信号
	// 例如：
	//   按 Ctrl+C  → "signal: interrupt"
	//   kill <pid> → "signal: terminated"
	fmt.Println(context.Cause(ctx))

	fmt.Println("exiting")
}
