package main

import (
	"fmt"
	"time"
)

// worker: 消费者，从 jobs channel 读取任务，将结果发送到 results channel
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs { // 持续读取直到 jobs 被关闭
		fmt.Println("Worker ", id, "started job", j)
		time.Sleep(time.Second) // 模拟耗时操作
		fmt.Println("Worker ", id, "finished job", j)
		results <- j * 2 // 返回结果（job编号 * 2）
	}
}

func main() {
	const numJobs = 5                  // 5个任务
	jobs := make(chan int, numJobs)    // 有缓冲的任务队列
	results := make(chan int, numJobs) // 有缓冲的结果队列

	// 启动3个 worker goroutines
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 发送5个任务到 jobs channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // 关闭 jobs，通知 workers 不会再有新任务

	// 等待接收所有结果
	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
