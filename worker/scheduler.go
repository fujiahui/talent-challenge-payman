package worker

import (
	util2 "github.com/fujiahui/talnet-challenge-payman/worker/util"
)

type Scheduler struct {
	pq *util2.JobPriorityQueue
}

func NewScheduler(enablePriority bool, cmp util2.CmpHandler) *Scheduler {
	return &Scheduler{
		pq: util2.NewJobPriorityQueue(enablePriority, cmp),
	}
}

// Enqueue 入队
func (s *Scheduler) Enqueue(job *util2.Job) {
	job.SetSleep()
	s.pq.PushBack(job)
}

// Dequeue 出队
func (s *Scheduler) Dequeue(freePoint uint16) *util2.Job {
	job := s.pq.PopFront(freePoint)
	if job == nil {
		return nil
	}
	job.SetWait()
	return job
}
