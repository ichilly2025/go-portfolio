package main

import (
	"time"
	"sync"
	"fmt"
)

type ImageTask struct {
	ID int
	Filename string
	Size int // KB
}

type ImageResult struct {
	ID int
	Filename string
	Success bool
	Error string
	Duration time.Duration
}

type WorkerPool interface {
	Submit(task ImageTask)
	Close()
	Results() <- chan ImageResult
}

type WorkerPoolService struct {
	mu sync.Mutex
	tasks chan ImageTask
	results chan ImageResult
}

func NewWorkerPoolService(jobs int) WorkerPool {
	service := &WorkerPoolService{}
	service.tasks = make(chan ImageTask, jobs)
	service.results = make(chan ImageResult, jobs)
	return service
}

func (service *WorkerPoolService) Submit(task ImageTask) {
	// service.mu.Lock()
	// defer service.mu.Unlock()
	service.tasks <- task
}

func (service *WorkerPoolService) Close() {
	// service.mu.Lock()
	// defer service.mu.Unlock()
	// close(service.tasks)
	close(service.results) // wait until worker is done
}

func (service *WorkerPoolService) Results() <- chan ImageResult {
	// service.mu.RLock()
	// defer service.mu.RUnlock()
	return service.results
}

func my_main() {
	workers := 3
	jobs := 5

	// var workerPool WorkerPool
	workerService := NewWorkerPoolService(jobs)
	// workerPool = workerService
	var wg sync.WaitGroup

	// Create workers to handle tasks
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			svc := workerService.(*WorkerPoolService)
			for task := range svc.tasks {
				result := ImageResult{
					ID: task.ID,
					Filename: task.Filename,
					Success: true,
					Error: "",
					Duration: time.Duration(time.Millisecond * time.Duration(task.Size)),
				}
				svc.results <- result
			}
		}(i)
	}

	// submit tasks
	for i := 0; i < jobs; i++ {
		task:= ImageTask{
				ID: i,
				Filename: fmt.Sprintf("file-%02d.png", i+1),
				Size: 100 * (i + 1),
		}
		workerService.Submit(task)
	}

	// close task channel
	svc := workerService.(*WorkerPoolService)
	close(svc.tasks)

	// wait until all workers are done
	wg.Wait()

	// read results
	for i := 0; i < jobs; i++ {
		result := <- workerService.Results()
		fmt.Printf("image task result %d: name=%s, success=%t\n", i, result.Filename, result.Success)
	}

	// close results channel
	workerService.Close()
}