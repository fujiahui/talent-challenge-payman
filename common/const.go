package common

type PriorityType uint8
type PointType uint16
type TimestampType int64

const (
	LowPriority     = PriorityType(0)
	HighPriority    = PriorityType(10)
	HighestPriority = PriorityType(100)
)
