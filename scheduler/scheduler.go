package scheduler

// Scheduler interface
type Scheduler interface {
	FilterNodes()
	Score()
	Pick()
}
