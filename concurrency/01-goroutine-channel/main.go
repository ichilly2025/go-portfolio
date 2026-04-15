package main

import (
	"fmt"
	"time"
)

// 生产者-消费者模式
func producer(ch chan<- int) {
	fmt.Println("生产者：开始生产")
	for i := 1; i <= 5; i++ {
		fmt.Printf("生产者：生产 %d\n", i)
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
	fmt.Println("生产者：生产完成")
}

func consumer(ch <-chan int) {
	fmt.Println("消费者：开始消费")
	for num := range ch {
		fmt.Printf("消费者：消费 %d\n", num)
		time.Sleep(1 * time.Second)
	}
	fmt.Println("消费者：消费完成")
}

func main() {
	fmt.Println("=== Goroutine + Channel (CSP 模型) ===")
	fmt.Println("不要通过共享内存来通信，而要通过通信来共享内存\n")

	ch := make(chan int, 3) // 带缓冲的 channel

	go producer(ch)
	consumer(ch)

	fmt.Println("\n程序结束")
}
