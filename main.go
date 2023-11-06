package main

import (
	"fmt"
	"sync"
	"time"
)

var TaskQueue = []Task{}

type Task struct {
	UserInfo    string
	RequestInfo string
}

type Rep struct {
}

type Worker struct {
	id        int
	host      string
	port      string
	isworking bool
}

type WorkPool struct {
	size    int
	workers []Worker
	wg      sync.WaitGroup
}

type Scheduler struct {
	WorkerPool WorkPool
	wg         sync.WaitGroup
}

func (s *Scheduler) start() {
	fmt.Printf("==============Scheduler 开始工作==============\n")
	for {
		for _, worker := range s.WorkerPool.workers {
			if !worker.isworking {
				task := s.getTask()
				go worker.ProcessRequest(task)
				break
			}
		}
	}
}

func (s *Scheduler) getTask() Task {
	task := TaskQueue[0]
	TaskQueue = TaskQueue[1:]
	return task
}

func (w *Worker) ProcessRequest(task Task) Rep {
	w.isworking = true
	fmt.Printf(" Worker%d UserInfo: %s RequestInfo：%s\n", w.id, task.UserInfo, task.RequestInfo)
	time.Sleep(2 * time.Second) // 模拟异步工作
	w.isworking = false
	return Rep{}
}

func main() {
	initTaskQueue()
	s := initScheduler()

	go func() {
		for i := 1; i <= 10; i++ {
			task := Task{
				UserInfo:    fmt.Sprintf(" %d", i),
				RequestInfo: fmt.Sprintf(" %d", i),
			}
			TaskQueue = append(TaskQueue, task)
		}
	}()
	time.Sleep(2 * time.Second)
	fmt.Print(TaskQueue)
	s.start()

}

func initTaskQueue() []Task {
	TaskQueue = make([]Task, 0)
	return TaskQueue
}

func initWorker(id int, host string, port string) Worker {
	return Worker{
		id:   id,
		host: host,
		port: port,
	}
}

func initWorkPool() WorkPool {
	return WorkPool{
		size: 2,
		workers: []Worker{
			initWorker(1, "127.0.0.1", "8080"),
			initWorker(2, "127.0.0.1", "8081"),
			initWorker(3, "127.0.0.1", "8082"),
		},
	}
}

func initScheduler() Scheduler {
	return Scheduler{
		WorkerPool: initWorkPool(),
		wg:         sync.WaitGroup{},
	}
}