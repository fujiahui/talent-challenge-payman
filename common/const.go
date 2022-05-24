package common

type (
	JobIDType     int64
	TimestampType int64
	PriorityType  uint8
	PointType     uint16
)

const (
	MaxCapacity = PointType(uint32(1<<16) - 1)
)

const (
	LowPriority     = PriorityType(0)
	HighPriority    = PriorityType(10)
	HighestPriority = PriorityType(100)
)
