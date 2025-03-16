// 本段代码使用workerPool模式实现类似线程池的效果
package main

import (
	"fmt"
	"sync"
	"time"
)

type WorkerPool struct {
	workers   int
	taskQueue chan func()
	wg        sync.WaitGroup
}

func NewWorkerPool(workers int) *WorkerPool {
	return &WorkerPool{
		workers:   workers,
		taskQueue: make(chan func()),
	}
}
func (wp *WorkerPool) Start() {
	wp.wg.Add(wp.workers)
	for i := 0; i < wp.workers; i++ {
		go func() {
			defer wp.wg.Done()
			for task := range wp.taskQueue {
				task()
			}
		}()
	}
}
func (wp *WorkerPool) Submit(task func()) {
	wp.taskQueue <- task
}
func (wp *WorkerPool) Shutdown() {
	close(wp.taskQueue)
	wp.wg.Wait()
}
func main() {
	pool := NewWorkerPool(3) //创建3个worker
	pool.Start()
	for i := 0; i < 6; i++ {
		id := i
		pool.Submit(func() {
			fmt.Printf("Task %d is running\n", id)
			time.Sleep(1 * time.Second)
			fmt.Printf("Task %d completed\n", id)
		})
	}
	pool.Shutdown()
}
