package main

import (
	"fmt"
	"sync"
)

// 单例模式
type Singleton struct {
	data string
}

var (
	instance *Singleton
	once     sync.Once
)

func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("初始化单例...")
		instance = &Singleton{data: "我是单例"}
	})
	return instance
}

func main() {
	fmt.Println("=== sync.Once (单次执行) ===")
	fmt.Println("确保初始化代码只执行一次（单例模式）\n")

	// 启动 10 个 goroutine 同时获取单例
	var wg sync.WaitGroup
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			s := GetInstance()
			fmt.Printf("Goroutine %d: 获取单例 -> %s (地址: %p)\n", id, s.data, s)
		}(i)
	}

	wg.Wait()
	fmt.Println("\n说明: 虽然有 10 个 goroutine，但初始化只执行了一次")
	fmt.Println("所有 goroutine 获取的是同一个实例（地址相同）")
}
