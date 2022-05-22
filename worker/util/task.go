package util

import (
	"fmt"
	"time"
)

type TaskStatusType uint8

const (
	TaskWait     = TaskStatusType(1)
	TaskRunning  = TaskStatusType(2)
	TaskFinished = TaskStatusType(4)
)

type Task struct {
	jobID       int64
	taskID      int
	remainPoint uint16 `json:"remain_point"` // Task剩余要执行的Points数
	/*
		预计开始的时间
			1. 刚创建的Job的第一个Task的 ExpectedTime = Job.Created
			2. Job的其他Task 等于前一个Task结束时间 +1s
	*/
	expectedTimestamp int64
	status            TaskStatusType

	speedFlag bool
}

func NewTask(point uint16) *Task {
	return &Task{
		remainPoint: point,
		status:      TaskWait,
	}
}

func (t *Task) Ticking() {
	if t.speedFlag && t.remainPoint > 0 && t.remainPoint%2 == 0 {
		t.remainPoint -= 2
	} else {
		t.remainPoint -= 1
	}

	if t.remainPoint == 0 {
		t.status = TaskFinished
	}
}

func (t *Task) RemainPoint() uint16 {
	if t.speedFlag && t.remainPoint > 0 && t.remainPoint%2 == 0 {
		return t.remainPoint / 2
	}
	return t.remainPoint

}

func (t *Task) Status() TaskStatusType {
	return t.status
}

func (t *Task) EnableSpeed() {
	t.speedFlag = true
}

func (t *Task) DisableSpeed() {
	t.speedFlag = false
}

func (t *Task) SetJobID(jobID int64) {
	t.jobID = jobID
}

func (t *Task) SetTaskID(taskID int) {
	t.taskID = taskID
}

func (t *Task) SetRunning() {
	t.status = TaskRunning
}

func (t *Task) SetExpectedTime(currTimestamp int64) {
	t.expectedTimestamp = currTimestamp
}

func (t *Task) ExpectedTimestamp() int64 {
	return t.expectedTimestamp
}

func (t *Task) Finished() bool {
	return t.status == TaskFinished
}

func (t *Task) Running(tick int) {
	remainPoint := t.remainPoint
	for point := remainPoint; point > 0; point-- {
		time.Sleep(time.Duration(tick) * time.Millisecond)
		fmt.Printf("JobRunning %d-%d(%d)", t.jobID, t.taskID, point)
	}
}
