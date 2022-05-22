package common

type PriorityType uint8

const (
	LowPriority     = PriorityType(0)
	HighPriority    = PriorityType(10)
	HighestPriority = PriorityType(100)
)
