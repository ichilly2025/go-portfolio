package main

import (
	"fmt"
	"time"
)

// 信号量
type Semaphore chan struct{}

func NewSemaphore(max int) Semaphore {
	return make(Semaphore, max)
}

func (s Semaphore) Acquire() {
	s <- struct{}{}
}

func (s Semaphore) Release() {
	<-s
}

func main() {
	fmt.Println("=== Semaphore (信号量) ===")
	fmt.Println("限制并发数量\n")

	const maxConcurrency = 3
	const totalTasks = 10

	sem := NewSemaphore(maxConcurrency)

	fmt.Printf("最大并发数: %d\n", maxConcurrency)
	fmt.Printf("总任务数: %d\n\n", totalTasks)

	for i := 1; i <= totalTasks; i++ {
		sem.Acquire() // 获取信号量
		go func(id int) {
			defer sem.Release() // 释放信号量

			fmt.Printf("任务 %d: 开始执行\n", id)
			time.Sleep(1 * time.Second)
			fmt.Printf("任务 %d: 执行完成\n", id)
		}(i)
	}

	// 等待所有任务完成
	time.Sleep(5 * time.Second)
	fmt.Println("\n所有任务完成")
}
