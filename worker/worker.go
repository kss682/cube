package worker

import (
	"cube/tasks"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

// Worker represents the
type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]tasks.Task
	TaskCount int
}

// RunTask method to run a task
func (w *Worker) RunTask() {

}

// StartTask method to start a task
func (w *Worker) StartTask() {

}

// StopTask method to stop a task
func (w *Worker) StopTask() {

}

// CollectStats method to collect stats
func (w *Worker) CollectStats() {

}
