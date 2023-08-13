package node

// Node represents the physical node in which workers run the tasks
type Node struct {
	Name            string
	IP              string
	Cores           int
	Memory          int
	MemoryAllocated int
	Disk            int
	DiskAllocated   int
	Role            string
	TaskCount       int
}
