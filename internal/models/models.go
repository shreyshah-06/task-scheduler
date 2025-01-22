package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        string
	Name      string
	Status    string
	CreatedAt time.Time
	updatedAt time.Time
}

func generateID() string {
	id:= uuid.New()
	return id.String()
}

func NewTask(name string) *Task{
	return &Task{
		ID: generateID(),
		Name : name,
		Status: "pending",
		CreatedAt: time.Now(),
		updatedAt: time.Now(),
	}
}

func (task *Task) GetStatus() string {
    return task.Status
}

func (task *Task) SetStatus(status string) {
	task.Status = status
	task.updatedAt = time.Now()
}

func (task *Task) UpdateTime(){
	task.updatedAt = time.Now()
}
