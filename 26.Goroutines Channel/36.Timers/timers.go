package main

import (
	"fmt"
	"time"
)

func main() {
	// 1. 创建第一个定时器，2秒后触发
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C                   // 阻塞等待，直到定时器触发
	fmt.Println("Timer 1 fired") // 2秒后输出

	// 2. 创建第二个定时器，1秒后触发
	timer2 := time.NewTimer(1 * time.Second)

	// 3. 在 goroutine 中异步等待 timer2
	go func() {
		<-timer2.C                   // 这个通道接收会阻塞在 goroutine 中
		fmt.Println("Timer 2 fired") // 这行不会执行
	}()

	// 4. 立即停止 timer2
	stop := timer2.Stop()
	if stop {
		fmt.Println("Timer 2 stopped") // 输出这行
	}

	// 5. 等待2秒，让程序不立即退出
	time.Sleep(2 * time.Second)
}
