package main

import (
	"fmt"
	"time"
)

type Task struct { //任务类型
	UserInfo    string
	RequestInfo string
}

var TaskQueue chan Task

type Rep struct {
}

func initTaskQueue() chan Task {
	TaskQueue = make(chan Task, 5)
	return TaskQueue
}

type Worker struct {
	id        int
	host      string
	port      string
	isworking chan struct{} //判断Worker是否工作，
}

type WorkPool struct {
	size    int
	workers []Worker
}

func initWorker(id int, host string, port string) Worker {
	w := Worker{
		id:        id,
		host:      host,
		port:      port,
		isworking: make(chan struct{}, 1),
	}
	w.isworking <- struct{}{}
	return w
}

func initWorkPool() WorkPool {
	return WorkPool{
		size: 2,
		workers: []Worker{
			initWorker(1, "127.0.0.1", "8080"),
			initWorker(2, "127.0.0.1", "8081"),
			initWorker(3, "127.0.0.1", "8082"),
			initWorker(4, "127.0.0.1", "8082"),
			initWorker(5, "127.0.0.1", "8082"),
		},
	}
}
func (w *Worker) ProcessRequest(task Task) Rep {
	time.Sleep(5 * time.Second) // 模拟异步工作
	fmt.Printf(" worker%d UserInfo: %s RequestInfo：%s\n", w.id, task.UserInfo, task.RequestInfo)
	w.isworking <- struct{}{}
	return Rep{}
}

type Scheduler struct {
	WorkerPool WorkPool
}

func initScheduler() Scheduler {
	return Scheduler{
		WorkerPool: initWorkPool(),
	}
}

func (s *Scheduler) start() {
	fmt.Printf("==============scheduler 开始工作==============\n")
	for {
		for i := 0; i < len(s.WorkerPool.workers); i++ {
			worker := s.WorkerPool.workers[i]
			select {
			case <-worker.isworking:
				select {
				case task := <-TaskQueue:
					go worker.ProcessRequest(task)
				default:
					// 任务队列为空，暂时不分配任务
				}
			default:
				continue
			}
		}
	}
}
func main() {
	initTaskQueue()
	s := initScheduler()
	go func() {
		for i := 1; i <= 25; i++ {
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
