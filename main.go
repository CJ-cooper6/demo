package main

import (
	"github.com/CJ-cooper6/demo/scheduler"
)

func main() {
	s := scheduler.InitScheduler()
	go s.Producer()
	s.Start()
}
