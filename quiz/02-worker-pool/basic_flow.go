package main

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID int
}

type Result struct {
	WorkerID int
	TaskID   int
	Value    int
}

type WorkerPool struct {
	wg            sync.WaitGroup
	workers       int
	numTasks      int
	taskChannel   chan Task
	resultChannel chan Result
}

func NewWorkerPool(workers, numTasks int) *WorkerPool {
	return &WorkerPool{
		workers:       workers,
		numTasks:      numTasks,
		taskChannel:   make(chan Task, numTasks),
		resultChannel: make(chan Result, numTasks),
	}
}

func (pool *WorkerPool) ProcessTasks() {
	for i := 0; i < pool.workers; i++ {
		pool.wg.Add(1)

		// Handle tasks
		go func(workerID int) {
			defer pool.wg.Done()
			for task := range pool.taskChannel {
				time.Sleep(time.Millisecond * 100)
				pool.resultChannel <- Result{
					WorkerID: workerID,
					TaskID:   task.ID,
					Value:    task.ID * 2,
				}
			}
		}(i)
	}
}

func (pool *WorkerPool) SubmitTask(task Task) {
	pool.taskChannel <- task
}

func (pool *WorkerPool) Close() {
	close(pool.taskChannel)
	go func() {
		pool.wg.Wait()
		close(pool.resultChannel)
	}()
}

func (pool *WorkerPool) GetResults() <-chan Result {
	return pool.resultChannel
}

func main() {
	pool := NewWorkerPool(3, 10)

	// 1. start 3 worker routine
	// cosumer for task channel
	// producer for result channel
	pool.ProcessTasks()

	// 2. submit 10 tasks
	// producer for task channel
	for i := 0; i < 10; i++ {
		pool.SubmitTask(Task{ID: i})
	}

	// 3. close task channel,
	// wait until workers are done,
	// close result channel
	pool.Close()

	// 4. read all results from result channel
	// consumer for result channel
	for result := range pool.GetResults() {
		fmt.Printf("[worker#%d, task#%d]result: %d\n",
			result.WorkerID, result.TaskID, result.Value)
	}
}

// func main() {
// 	workers := 3
// 	numTasks := 10
// 	var wg sync.WaitGroup
// 	taskChannel := make(chan int, numTasks)
// 	resultChannel := make(chan Result, numTasks)

// 	for i := 0; i < workers; i++ {
// 		wg.Add(1)

// 		// Handle tasks
// 		go func(workerID int) {
// 			defer wg.Done()
// 			for task := range taskChannel {
// 				time.Sleep(time.Millisecond * 100)
// 				resultChannel <- Result{
// 					WorkerID: workerID,
// 					TaskID: task,
// 					Value: task * 2,
// 				}
// 			}
// 		}(i)
// 	}

// 	// Add tasks
// 	for i := 0; i < numTasks; i++ {
// 		taskChannel <- i
// 	}

// 	// close tasks chanel once producer is done
// 	close(taskChannel)

// 	// start another go routine to wait and close result channel
// 	go func() {
// 		wg.Wait()
// 		close(resultChannel)
// 	}()

// 	// print result
// 	for result := range resultChannel {
// 		fmt.Printf("[%d] task %d result: %d\n", result.WorkerID, result.TaskID, result.Value)
// 	}
// }
