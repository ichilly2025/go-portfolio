package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Select (多路复用) ===")
	fmt.Println("监听多个 Channel\n")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// Goroutine 1: 1 秒后发送消息
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "来自 Channel 1 的消息"
	}()

	// Goroutine 2: 2 秒后发送消息
	go func() {
		time.Sleep(2 * time.Second)
		ch2 <- "来自 Channel 2 的消息"
	}()

	// 使用 select 监听多个 channel
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Println("收到:", msg1)
		case msg2 := <-ch2:
			fmt.Println("收到:", msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("超时：3 秒内没有收到消息")
		}
	}

	fmt.Println("\n程序结束")
}
