package worker

import (
	"github.com/fujiahui/talnet-challenge-payman/common"
	"github.com/fujiahui/talnet-challenge-payman/logger"
	"github.com/fujiahui/talnet-challenge-payman/worker/util"
)

type Scheduler struct {
	enablePriority bool
	pq             *util.JobPriorityQueue
}

func NewScheduler(enablePriority bool, cmp util.CmpHandler) *Scheduler {
	return &Scheduler{
		enablePriority: enablePriority,
		pq:             util.NewJobPriorityQueue(cmp),
	}
}

// Enqueue 入队
func (s *Scheduler) Enqueue(job *util.Job) {
	job.SetSleep()
	priority := job.Priority()
	if !s.enablePriority {
		priority = common.LowPriority
	}
	s.pq.PushBack(priority, job)
}

// Dequeue 出队
func (s *Scheduler) Dequeue(freePoint common.PointType) *util.Job {
	x := s.pq.PopFront()
	if x == nil {
		return nil
	}

	job := x.(*util.Job)
	if job.CurrTask().RemainPoint() > freePoint {
		priority := job.Priority()
		if !s.enablePriority {
			priority = common.LowPriority
		}
		s.pq.PushFront(priority, job)
		logger.Warnf("Scheduler.Dequeue return nil, because RemainPoint %d more than freePoint %d", job.CurrTask().RemainPoint(), freePoint)
		return nil
	}

	job.SetWait()
	return job
}
