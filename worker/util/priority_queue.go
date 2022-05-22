package util

import (
	"container/heap"
	"container/list"
	"github.com/fujiahui/talnet-challenge-payman/common"
)

type CmpHandler func(j1 *Job, j2 *Job) bool

type JobQueue struct {
	priority common.PriorityType
	Queue    *list.List // 同一优先级进行排序
	index    int
}

func NewJobQueue(priority common.PriorityType) *JobQueue {
	return &JobQueue{
		priority: priority,
		Queue:    list.New(),
	}
}

func (q *JobQueue) Front() *Job {
	e := q.Queue.Front()
	if e == nil {
		return nil
	}
	return e.Value.(*Job)
}

func (q *JobQueue) Len() int {
	return q.Queue.Len()
}

func (q *JobQueue) PushBack(job *Job) {
	q.Queue.PushBack(job)
}

func (q *JobQueue) PopFront() *Job {
	e := q.Queue.Front()
	if e == nil {
		return nil
	}
	q.Queue.Remove(e)
	job := e.Value.(*Job)
	return job
}

type JobPriorityQueue struct {
	enablePriority    bool
	queue             []*JobQueue
	Priority2JobQueue map[common.PriorityType]*JobQueue

	cmp CmpHandler
}

func NewJobPriorityQueue(enablePriority bool, cmp CmpHandler) *JobPriorityQueue {

	pq := &JobPriorityQueue{
		enablePriority:    enablePriority,
		queue:             make([]*JobQueue, 0, 16),
		Priority2JobQueue: make(map[common.PriorityType]*JobQueue),
		cmp:               cmp,
	}
	heap.Init(pq)
	return pq
}

func (pq JobPriorityQueue) Len() int { return len(pq.queue) }

func (pq JobPriorityQueue) Less(i, j int) bool {
	j1 := pq.queue[i].Queue.Front().Value.(*Job)
	j2 := pq.queue[j].Queue.Front().Value.(*Job)

	return pq.cmp(j1, j2)
}

func (pq JobPriorityQueue) Swap(i, j int) {
	pq.queue[i], pq.queue[j] = pq.queue[j], pq.queue[i]
	pq.queue[i].index = i
	pq.queue[j].index = j
}

func (pq *JobPriorityQueue) Push(x interface{}) {
	n := len(pq.queue)
	item := x.(*JobQueue)
	item.index = n
	pq.queue = append(pq.queue, item)
	//
	priority := item.priority
	pq.Priority2JobQueue[priority] = item
}

func (pq *JobPriorityQueue) Pop() interface{} {
	old := pq.queue
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	pq.queue = old[0 : n-1]
	//
	priority := item.priority
	delete(pq.Priority2JobQueue, priority)
	return item
}

func (pq *JobPriorityQueue) PopFront(freePoint uint16) *Job {
	if pq.Len() == 0 {
		return nil
	}

	q := heap.Pop(pq).(*JobQueue)
	job := q.Front()
	if job.CurrTask().RemainPoint() > freePoint {
		heap.Push(pq, q)
		return nil
	}

	q.PopFront()
	if q.Len() > 0 {
		heap.Push(pq, q)
	}

	return job
}

func (pq *JobPriorityQueue) PushBack(job *Job) {
	priority := common.LowPriority
	if pq.enablePriority {
		priority = job.Priority()
	}

	if q, ok := pq.at(priority); !ok {
		q = NewJobQueue(priority)
		q.PushBack(job)
		heap.Push(pq, q)
	} else {
		q.PushBack(job)
	}
}

func (pq *JobPriorityQueue) at(priority common.PriorityType) (*JobQueue, bool) {
	q, ok := pq.Priority2JobQueue[priority]
	return q, ok
}
