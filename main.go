package main

import (
	"fmt"
	"sync"
	"time"
)

var TaskQueue chan Task

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
		for i := 0; i < len(s.WorkerPool.workers); i++ {
			worker := s.WorkerPool.workers[i]
			if !worker.isworking {
				select {
				case task := <-TaskQueue:
					go worker.ProcessRequest(task)
				default:
					// 任务队列为空，暂时不分配任务
				}
			} else {
				fmt.Println("在工作啦！")
			}
		}
	}
}

func (w *Worker) ProcessRequest(task Task) Rep {
	w.isworking = true
	time.Sleep(5 * time.Second) // 模拟异步工作
	fmt.Printf(" Worker%d UserInfo: %s RequestInfo：%s\n", w.id, task.UserInfo, task.RequestInfo)
	w.isworking = false
	return Rep{}
}

func initTaskQueue() chan Task {
	TaskQueue = make(chan Task, 5)
	return TaskQueue
}

func initWorker(id int, host string, port string, isworking bool) Worker {
	return Worker{
		id:        id,
		host:      host,
		port:      port,
		isworking: isworking,
	}
}

func initWorkPool() WorkPool {
	return WorkPool{
		size: 2,
		workers: []Worker{
			initWorker(1, "127.0.0.1", "8080", false),
			initWorker(2, "127.0.0.1", "8081", false),
			//initWorker(3, "127.0.0.1", "8082", false),
		},
	}
}

func initScheduler() Scheduler {
	return Scheduler{
		WorkerPool: initWorkPool(),
		wg:         sync.WaitGroup{},
	}
}

func main() {
	initTaskQueue()
	s := initScheduler()
	go func() {
		for i := 1; i <= 50; i++ {
			task := Task{
				UserInfo:    fmt.Sprintf(" %d", i),
				RequestInfo: fmt.Sprintf(" %d", i),
			}
			TaskQueue <- task
			//time.Sleep(2 * time.Second)
		}
	}()
	s.start()
}
