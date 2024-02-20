package workers

import (
	"fmt"
	"sync"
	"time"
)

type Task interface {
	Process()
}

type SendEmailTask struct {
	Email   string
	Message string
	Subject string
}

func (t *SendEmailTask) Process() {
	fmt.Println("Sending email to:", t.Email)
	time.Sleep(2 * time.Second)
}

type ImageProcessingTask struct {
	ImageURL string
}

func (t *ImageProcessingTask) Process() {
	fmt.Println("Processing image:", t.ImageURL)
	time.Sleep(2 * time.Second)
}

// worker pool definition
type WorkerPool struct {
	Tasks       []Task
	Concurrency int
	taskChan    chan Task
	wg          sync.WaitGroup
}

// func to execute the worker pool
func (wp *WorkerPool) Worker() {
	for task := range wp.taskChan {
		task.Process()
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run() {
	wp.taskChan = make(chan Task, len(wp.Tasks))

	for range wp.Concurrency {
		go wp.Worker()
	}

	wp.wg.Add(len(wp.Tasks))

	for _, task := range wp.Tasks {
		wp.taskChan <- task
	}

	close(wp.taskChan)
	wp.wg.Wait()
}
