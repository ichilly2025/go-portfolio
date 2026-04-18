package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("=== Channel（通道）===\n")

	// 1. 创建 Channel
	fmt.Println("1. 创建 Channel:")
	ch1 := make(chan int)       // 无缓冲
	ch2 := make(chan int, 5)    // 有缓冲（容量 5）
	fmt.Printf("   无缓冲: %T\n", ch1)
	fmt.Printf("   有缓冲: %T, 容量: %d\n", ch2, cap(ch2))

	// 2. 发送和接收
	fmt.Println("\n2. 发送和接收:")
	go func() {
		ch1 <- 42 // 发送
	}()
	value := <-ch1 // 接收
	fmt.Printf("   接收到: %d\n", value)

	// 3. 有缓冲 Channel
	fmt.Println("\n3. 有缓冲 Channel:")
	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	fmt.Printf("   发送了 3 个值\n")
	fmt.Printf("   接收: %d\n", <-ch2)
	fmt.Printf("   接收: %d\n", <-ch2)
	fmt.Printf("   接收: %d\n", <-ch2)

	// 4. 关闭 Channel
	fmt.Println("\n4. 关闭 Channel:")
	ch3 := make(chan int, 3)
	ch3 <- 1
	ch3 <- 2
	ch3 <- 3
	close(ch3)
	fmt.Println("   Channel 已关闭")

	// 从已关闭的 channel 接收
	for i := 0; i < 4; i++ {
		v, ok := <-ch3
		fmt.Printf("   接收: %d, ok: %t\n", v, ok)
	}

	// 5. range 遍历
	fmt.Println("\n5. range 遍历:")
	ch4 := make(chan int, 3)
	ch4 <- 10
	ch4 <- 20
	ch4 <- 30
	close(ch4)

	for v := range ch4 {
		fmt.Printf("   接收: %d\n", v)
	}

	// 6. 单向 Channel
	fmt.Println("\n6. 单向 Channel:")
	ch5 := make(chan int)

	// 只发送
	go func(ch chan<- int) {
		ch <- 100
		fmt.Println("   发送完成")
	}(ch5)

	// 只接收
	go func(ch <-chan int) {
		v := <-ch
		fmt.Printf("   接收: %d\n", v)
	}(ch5)

	time.Sleep(100 * time.Millisecond)

	// 7. select 多路复用
	fmt.Println("\n7. select 多路复用:")
	ch6 := make(chan string)
	ch7 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		ch6 <- "来自 ch6"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		ch7 <- "来自 ch7"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch6:
			fmt.Printf("   %s\n", msg)
		case msg := <-ch7:
			fmt.Printf("   %s\n", msg)
		case <-time.After(300 * time.Millisecond):
			fmt.Println("   超时")
		}
	}

	// 8. 注意事项
	fmt.Println("\n8. 注意事项:")
	fmt.Println("   - 向已关闭的 channel 发送会 panic")
	fmt.Println("   - 从已关闭的 channel 接收会返回零值")
	fmt.Println("   - 关闭已关闭的 channel 会 panic")
	fmt.Println("   - 无缓冲 channel 发送和接收必须同时准备好")
}
