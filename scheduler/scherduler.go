package scheduler

import (
	"fmt"
	"github.com/CJ-cooper6/demo/task"
	"github.com/CJ-cooper6/demo/worker"
)

type Scheduler struct {
	WorkerPool worker.WorkPool
	TaskQueue  chan task.Task
}

func InitScheduler() Scheduler {
	return Scheduler{
		WorkerPool: worker.InitWorkPool(),
		TaskQueue:  task.InitTaskQueue(),
	}
}

func (s *Scheduler) Producer() {
	for i := 1; i <= 25; i++ {
		task := task.Task{
			UserInfo:    fmt.Sprintf(" %d", i),
			RequestInfo: fmt.Sprintf(" %d", i),
		}
		s.TaskQueue <- task
	}
}

func (s *Scheduler) Start() {
	fmt.Printf("==============scheduler 开始工作==============\n")
	for {
		for i := 0; i < len(s.WorkerPool.Workers); i++ {
			worker := s.WorkerPool.Workers[i]
			select {
			case <-worker.Isworking:
				task := <-s.TaskQueue
				go worker.ProcessRequest(task)
			default:
				continue
			}
		}
	}
}
