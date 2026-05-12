package main

import (
	"fmt"
	"sync"
	"time"
)

// worker 是一个模拟耗时任务的函数。
// 参数 id 用于标识不同的 worker。
// 实际场景中，这里可以是网络请求、文件读写等 I/O 操作。
func worker(id int) {
	fmt.Printf("Worker %d starting\n", id)

	// 模拟耗时任务，休眠 1 秒
	time.Sleep(1 * time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	// 声明一个 WaitGroup，用于等待所有 goroutine 完成
	var wg sync.WaitGroup

	// 循环启动 5 个 worker 协程
	for i := 1; i <= 5; i++ {
		// wg.Go() 是 Go 1.24 新增的便捷方法，等价于：
		//   wg.Add(1)
		//   go func() {
		//       defer wg.Done()
		//       worker(i)
		//   }()
		// 它会自动管理计数器的增减，并在新 goroutine 中执行传入的函数。
		wg.Go(func() {
			worker(i)
		})
	}

	// 阻塞 main 协程，直到所有 worker 执行完毕（WaitGroup 计数器归零）
	wg.Wait()
}
