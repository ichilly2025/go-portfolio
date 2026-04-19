package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type ImageTask struct {
	ID       int
	Filename string
	Size     int // KB
}

type ImageResult struct {
	ID       int
	Filename string
	Success  bool
	Error    string
	Duration time.Duration
}

type WorkerPool interface {
	Submit(task ImageTask)
	Close()
	Results() <-chan ImageResult
}

type WorkerPoolService struct {
	tasks   chan ImageTask
	results chan ImageResult
	wg      sync.WaitGroup
	ctx     context.Context
	cancel  context.CancelFunc
}

func NewWorkerPoolService(numWorkers, bufferSize int) WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	service := &WorkerPoolService{
		tasks:   make(chan ImageTask, bufferSize),
		results: make(chan ImageResult, bufferSize),
		ctx:     ctx,
		cancel:  cancel,
	}

	// 启动固定数量的 Worker
	for i := 0; i < numWorkers; i++ {
		service.wg.Add(1)
		go service.worker(i + 1)
	}

	return service
}

func (service *WorkerPoolService) worker(workerID int) {
	defer service.wg.Done()

	for {
		select {
		case task, ok := <-service.tasks:
			if !ok {
				// tasks channel 已关闭，退出
				return
			}
			// 处理任务
			result := service.processTask(workerID, task)
			service.results <- result

		case <-service.ctx.Done():
			// Context 取消，退出
			return
		}
	}
}

func (service *WorkerPoolService) processTask(workerID int, task ImageTask) ImageResult {
	start := time.Now()

	// 错误检查：Size <= 0
	if task.Size <= 0 {
		return ImageResult{
			ID:       task.ID,
			Filename: task.Filename,
			Success:  false,
			Error:    "invalid image size",
			Duration: 0,
		}
	}

	// 错误检查：Size > 1000
	if task.Size > 1000 {
		return ImageResult{
			ID:       task.ID,
			Filename: task.Filename,
			Success:  false,
			Error:    "image too large",
			Duration: 0,
		}
	}

	// 模拟图片处理耗时
	fmt.Printf("[Worker %d] 开始处理 Task %d: %s (%dKB)\n", workerID, task.ID, task.Filename, task.Size)
	time.Sleep(time.Duration(task.Size) * time.Millisecond)

	duration := time.Since(start)
	fmt.Printf("[Worker %d] 完成 Task %d (耗时: %v)\n", workerID, task.ID, duration)

	return ImageResult{
		ID:       task.ID,
		Filename: task.Filename,
		Success:  true,
		Error:    "",
		Duration: duration,
	}
}

func (service *WorkerPoolService) Submit(task ImageTask) {
	service.tasks <- task
}

func (service *WorkerPoolService) Close() {
	close(service.tasks)  // 关闭任务 channel
	service.wg.Wait()     // 等待所有 Worker 完成
	close(service.results) // 关闭结果 channel
}

func (service *WorkerPoolService) Results() <-chan ImageResult {
	return service.results
}

func main2() {
	fmt.Println("=== 图片处理 Worker Pool ===\n")

	workers := 3
	jobs := 5

	// 创建 Worker Pool
	fmt.Printf("创建 Worker Pool (%d 个 Worker)\n\n", workers)
	workerPool := NewWorkerPoolService(workers, jobs)

	// 提交任务
	fmt.Println("提交任务:")
	tasks := []ImageTask{
		{ID: 1, Filename: "photo1.jpg", Size: 100},
		{ID: 2, Filename: "photo2.jpg", Size: 200},
		{ID: 3, Filename: "photo3.jpg", Size: 150},
		{ID: 4, Filename: "photo4.jpg", Size: 300},
		{ID: 5, Filename: "photo5.jpg", Size: 250},
	}

	for _, task := range tasks {
		fmt.Printf("  Task %d: %s (%dKB)\n", task.ID, task.Filename, task.Size)
		workerPool.Submit(task)
	}

	fmt.Println("\n处理中...")

	// 在后台收集结果
	go func() {
		// 等待所有任务完成后关闭
		time.Sleep(time.Second * 2)
		workerPool.Close()
	}()

	// 收集结果
	fmt.Println("\n结果统计:")
	successCount := 0
	failCount := 0
	totalDuration := time.Duration(0)

	for result := range workerPool.Results() {
		if result.Success {
			successCount++
			totalDuration += result.Duration
		} else {
			failCount++
			fmt.Printf("  ❌ Task %d 失败: %s\n", result.ID, result.Error)
		}
	}

	fmt.Printf("  总任务数: %d\n", len(tasks))
	fmt.Printf("  成功: %d\n", successCount)
	fmt.Printf("  失败: %d\n", failCount)
	if successCount > 0 {
		avgDuration := totalDuration / time.Duration(successCount)
		fmt.Printf("  平均耗时: %v\n", avgDuration)
	}

	fmt.Println("\n所有任务完成！")
}

// 测试错误处理
func testErrorHandling() {
	fmt.Println("\n=== 测试错误处理 ===\n")

	workerPool := NewWorkerPoolService(2, 10)

	// 测试无效 Size
	workerPool.Submit(ImageTask{ID: 1, Filename: "invalid.jpg", Size: 0})
	workerPool.Submit(ImageTask{ID: 2, Filename: "negative.jpg", Size: -10})

	// 测试超大文件
	workerPool.Submit(ImageTask{ID: 3, Filename: "huge.jpg", Size: 2000})

	// 正常文件
	workerPool.Submit(ImageTask{ID: 4, Filename: "normal.jpg", Size: 100})

	go func() {
		time.Sleep(time.Second)
		workerPool.Close()
	}()

	for result := range workerPool.Results() {
		if result.Success {
			fmt.Printf("✅ Task %d: %s 处理成功 (耗时: %v)\n", result.ID, result.Filename, result.Duration)
		} else {
			fmt.Printf("❌ Task %d: %s 处理失败 - %s\n", result.ID, result.Filename, result.Error)
		}
	}
}

// 测试并发性能
func testConcurrency() {
	fmt.Println("\n=== 测试并发性能 ===\n")

	workers := 3
	numTasks := 9

	workerPool := NewWorkerPoolService(workers, numTasks)

	start := time.Now()

	// 提交 9 个任务，每个 100ms
	for i := 0; i < numTasks; i++ {
		workerPool.Submit(ImageTask{
			ID:       i + 1,
			Filename: fmt.Sprintf("task-%d.jpg", i+1),
			Size:     100,
		})
	}

	go func() {
		time.Sleep(time.Second * 2)
		workerPool.Close()
	}()

	count := 0
	for range workerPool.Results() {
		count++
	}

	elapsed := time.Since(start)

	fmt.Printf("处理 %d 个任务 (每个 100ms)\n", numTasks)
	fmt.Printf("使用 %d 个 Worker\n", workers)
	fmt.Printf("总耗时: %v\n", elapsed)
	fmt.Printf("串行需要: %v\n", time.Duration(numTasks*100)*time.Millisecond)
	fmt.Printf("并发加速: %.2fx\n", float64(numTasks*100)/float64(elapsed.Milliseconds()))
}

// 如果想运行测试，取消注释：
// func main() {
// 	testErrorHandling()
// 	testConcurrency()
// }
