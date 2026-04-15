package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d: 收到取消信号，停止工作\n", id)
			return
		default:
			fmt.Printf("Worker %d: 正在工作...\n", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	fmt.Println("=== Context (上下文控制) ===")
	fmt.Println("超时控制和取消信号传播\n")

	// 创建一个 3 秒超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	// 启动 3 个 worker
	for i := 1; i <= 3; i++ {
		go worker(ctx, i)
	}

	// 等待 context 超时
	<-ctx.Done()
	fmt.Println("\n主程序：Context 超时，所有 worker 已停止")

	// 等待一下让输出完整
	time.Sleep(500 * time.Millisecond)
}
