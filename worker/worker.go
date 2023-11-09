package worker

import (
	"fmt"
	"github.com/CJ-cooper6/demo/task"
	"time"
)

type Worker struct {
	Id        int
	host      string
	port      string
	Isworking chan struct{} //判断Worker是否工作，
}

type WorkPool struct {
	size    int
	Workers []Worker
}

func InitWorker(id int, host string, port string) Worker {
	w := Worker{
		Id:        id,
		host:      host,
		port:      port,
		Isworking: make(chan struct{}, 1),
	}
	w.Isworking <- struct{}{}
	return w
}

func InitWorkPool() WorkPool {
	return WorkPool{
		size: 2,
		Workers: []Worker{
			InitWorker(1, "127.0.0.1", "8080"),
			InitWorker(2, "127.0.0.1", "8081"),
			InitWorker(3, "127.0.0.1", "8082"),
			InitWorker(4, "127.0.0.1", "8082"),
			InitWorker(5, "127.0.0.1", "8082"),
		},
	}
}
func (w *Worker) ProcessRequest(task task.Task) task.Rep {
	defer func() {
		w.Isworking <- struct{}{}
	}()
	time.Sleep(2 * time.Second) // 模拟异步工作
	fmt.Printf(" worker%d UserInfo: %s RequestInfo：%s\n", w.Id, task.UserInfo, task.RequestInfo)
	return struct{}{}
}
