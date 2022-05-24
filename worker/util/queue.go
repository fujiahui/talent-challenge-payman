package util

import (
	"container/list"
	"github.com/fujiahui/talnet-challenge-payman/common"
	"github.com/fujiahui/talnet-challenge-payman/logger"
)

type JobQueue struct {
	priority common.PriorityType
	queue    *list.List // 同一优先级进行排序
	index    int
}

func NewJobQueue(priority common.PriorityType) *JobQueue {
	return &JobQueue{
		priority: priority,
		queue:    list.New(),
	}
}

func (q *JobQueue) Front() any {
	e := q.queue.Front()
	if e == nil {
		return nil
	}
	return e.Value
}

func (q *JobQueue) Len() int {
	return q.queue.Len()
}

func (q *JobQueue) PushBack(x any) {
	q.queue.PushBack(x)
}

func (q *JobQueue) PushFront(x any) {
	q.queue.PushFront(x)
}

func (q *JobQueue) PopFront() any {
	e := q.queue.Front()
	if e == nil {
		logger.Errorf("JobQueue.PushFront return nil")
		return nil
	}
	q.queue.Remove(e)
	return e.Value
}
