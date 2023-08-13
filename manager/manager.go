package manager

import (
	"cube/tasks"
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

// Manager represents the manager's identity
type Manager struct {
	Pending       queue.Queue
	TaskDB        map[string][]tasks.Task
	EventDB       map[string][]tasks.TaskEvent
	Worker        []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

// SelectWorker method resonsible for checking the requirements of a task
func (m *Manager) SelectWorker() {
	fmt.Println("I will select appropriate worker")
}

// UpdateTask method to keep track of tasks and their states
func (m *Manager) UpdateTask() {
	fmt.Println("I will update task")
}

// SendWork method to send task to worker
func (m *Manager) SendWork() {
	fmt.Println("I will send work to workers")
}