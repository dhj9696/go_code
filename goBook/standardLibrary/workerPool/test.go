package main

import (
	"fmt"
	"sync"
)

// 任务结构体
type Task struct {
	ID int
	// 可以添加更多的任务相关字段
}

// 线程池结构体
type ThreadPool struct {
	workers chan struct{}
	tasks   chan Task
	wg      sync.WaitGroup
}

// 创建线程池
func NewThreadPool(size int, taskQueueSize int) *ThreadPool {
	return &ThreadPool{
		workers: make(chan struct{}, size),
		tasks:   make(chan Task, taskQueueSize),
	}
}

// 启动线程池
func (tp *ThreadPool) Start() {
	go func() {
		for task := range tp.tasks {
			tp.workers <- struct{}{}
			tp.wg.Add(1)
			go func(t Task) {
				defer tp.wg.Done()
				defer func() { <-tp.workers }()
				// 执行任务
				fmt.Printf("Processing task %d\n", t.ID)
			}(task)
		}
	}()
}

// 提交任务
func (tp *ThreadPool) Submit(task Task) {
	tp.tasks <- task
}

// 关闭线程池
func (tp *ThreadPool) Shutdown() {
	close(tp.tasks)
	tp.wg.Wait()
}

func main() {
	// 创建一个包含 3 个工作线程，任务队列大小为 5 的线程池
	tp := NewThreadPool(3, 5)
	// 启动线程池
	tp.Start()

	// 提交 10 个任务
	for i := 0; i < 10; i++ {
		tp.Submit(Task{ID: i})
	}

	// 关闭线程池
	tp.Shutdown()
}
