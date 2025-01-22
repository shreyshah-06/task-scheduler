package scheduler

import (
	"sync"

	"github.com/shreyshah-06/taskScheduler/internal/models"
)

type TaskScheduler struct {
	tasks []models.Task
	wg sync.WaitGroup
}

func NewTaskScheduler(tasks []models.Task) *TaskScheduler {
	return &TaskScheduler{
		tasks: tasks,
	}
}

func (ts *TaskScheduler) Start(){
	
	taskQueue:= make(chan models.Task, len(ts.tasks))

	for i:=1;i<=3;i++{
		worker := NewWorker(i, taskQueue, &ts.wg)
		worker.Start()
	}
}