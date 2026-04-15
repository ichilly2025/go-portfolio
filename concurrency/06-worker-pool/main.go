package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker %d: 开始处理任务 %d\n", id, job)
		time.Sleep(time.Second) // 模拟耗时操作
		fmt.Printf("Worker %d: 完成任务 %d\n", id, job)
		results <- job * 2
	}
}

func main() {
	fmt.Println("=== Worker Pool (工作池) ===")
	fmt.Println("固定数量的 Worker 处理任务队列\n")

	const numWorkers = 3
	const numJobs = 9

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// 启动 3 个 worker
	for w := 1; w <= numWorkers; w++ {
		go worker(w, jobs, results)
	}

	// 发送 9 个任务
	fmt.Printf("发送 %d 个任务到队列\n\n", numJobs)
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// 收集结果
	fmt.Println("\n收集结果:")
	for a := 1; a <= numJobs; a++ {
		result := <-results
		fmt.Printf("结果 %d: %d\n", a, result)
	}

	fmt.Println("\n所有任务完成")
}
