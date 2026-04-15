package main

import (
	"fmt"
	"sync"
	"time"
)

// 共享内存 + 锁模型
type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

func main() {
	fmt.Println("=== Mutex (互斥锁) ===")
	fmt.Println("共享内存 + 锁模型\n")

	counter := &Counter{}

	// 启动 10 个 goroutine 并发增加计数器
	for i := 0; i < 10; i++ {
		go func(id int) {
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
			fmt.Printf("Goroutine %d 完成\n", id)
		}(i)
	}

	time.Sleep(2 * time.Second)
	fmt.Printf("\n最终计数: %d (预期: 1000)\n", counter.Value())
}
