package task

type Task struct { //任务类型
	UserInfo    string
	RequestInfo string
}
type Rep struct {
}

func InitTaskQueue() chan Task {
	TaskQueue := make(chan Task, 5)
	return TaskQueue
}
