package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d: 开始工作\n", id)
	time.Sleep(time.Duration(id) * 500 * time.Millisecond)
	fmt.Printf("Worker %d: 工作完成\n", id)
}

func main() {
	fmt.Println("=== WaitGroup (等待组) ===")
	fmt.Println("等待多个 Goroutine 完成\n")

	var wg sync.WaitGroup

	// 启动 5 个 worker
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	fmt.Println("主程序：等待所有 worker 完成...")
	wg.Wait()
	fmt.Println("\n主程序：所有 worker 已完成")
}
