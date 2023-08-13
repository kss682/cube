package tasks

import (
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

// State to define the state of tasks
type State int

// Constants defining different states of tasks
const (
	Pending State = iota
	Scheduled
	Runnings
	Completed
	Failed
)

// Task struct represents what identifies as a task
type Task struct {
	ID            uuid.UUID
	Name          string
	State         State
	Image         string
	Memory        string
	Disk          string
	ExposedPorts  nat.PortSet
	RestartPolicy string
	StartTime     time.Time
	EndTime       time.Time
}

// TaskEvent represents the event that moves a task from one state to another (still not really clear the purpose of this Task)
type TaskEvent struct {
	ID        uuid.UUID
	State     State
	Timestamp time.Time
	Task      Task
}
