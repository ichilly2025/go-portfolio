package main

import (
	"fmt"
	"sync"
	"time"
)

// 生成数字
func generate(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

// Worker：计算平方（模拟耗时操作）
func square(id int, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			fmt.Printf("  Worker %d: 处理 %d\n", id, n)
			time.Sleep(100 * time.Millisecond) // 模拟耗时
			out <- n * n
		}
		close(out)
	}()
	return out
}

// Fan-out：将一个输入 channel 分发给多个 worker
func fanOut(in <-chan int, numWorkers int) []<-chan int {
	channels := make([]<-chan int, numWorkers)
	for i := 0; i < numWorkers; i++ {
		channels[i] = square(i+1, in)
	}
	return channels
}

// Fan-in：合并多个 channel 到一个输出
func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for n := range c {
				out <- n
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	fmt.Println("=== Fan-out / Fan-in (扇出/扇入) ===")
	fmt.Println("一个输入，多个处理器，一个输出\n")

	// 生成数字
	in := generate(1, 2, 3, 4, 5, 6, 7, 8)

	// Fan-out: 启动 3 个 worker 并行处理
	fmt.Println("Fan-out: 将任务分发给 3 个 worker")
	workers := fanOut(in, 3)

	// Fan-in: 合并结果
	fmt.Println("Fan-in: 合并所有 worker 的结果\n")
	out := fanIn(workers...)

	// 输出结果
	fmt.Println("结果:")
	for result := range out {
		fmt.Printf("%d ", result)
	}
	fmt.Println("\n\n说明: 多个 worker 并行处理，结果顺序可能不同")
}
