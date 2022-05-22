package worker

import (
	"github.com/fujiahui/talnet-challenge-payman/common"
	"github.com/fujiahui/talnet-challenge-payman/worker/util"
)

type Scheduler struct {
	pq *util.JobPriorityQueue
}

func NewScheduler(enablePriority bool, cmp util.CmpHandler) *Scheduler {
	return &Scheduler{
		pq: util.NewJobPriorityQueue(enablePriority, cmp),
	}
}

// Enqueue 入队
func (s *Scheduler) Enqueue(job *util.Job) {
	job.SetSleep()
	s.pq.PushBack(job)
}

// Dequeue 出队
func (s *Scheduler) Dequeue(freePoint common.PointType) *util.Job {
	job := s.pq.PopFront(freePoint)
	if job == nil {
		return nil
	}
	job.SetWait()
	return job
}
