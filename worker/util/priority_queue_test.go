package util

import (
	"github.com/fujiahui/talnet-challenge-payman/common"
	"testing"
)

func TestJobQueue(t *testing.T) {

	jobs := []*Job{
		NewJob(1, 1, common.LowPriority, []*Task{
			NewTask(5),
			NewTask(6),
			NewTask(7),
		}),
		NewJob(2, 3, common.LowPriority, []*Task{
			NewTask(3),
			NewTask(5),
		}),
	}

	q := NewJobQueue(common.LowPriority)
	for _, job := range jobs {
		q.PushBack(job)
	}

	for q.Len() > 0 {
		job := q.PopFront()
		t.Log(job.String())
	}
}

func TestJobPriorityQueueWithNothing(t *testing.T) {
	//
	jobs := []*Job{
		NewJob(1, 6, common.LowPriority, []*Task{
			NewTask(6),
			NewTask(7),
		}),
		NewJob(2, 6, common.HighPriority, []*Task{
			NewTask(5),
		}),
	}

	pq := NewJobPriorityQueue(false, BaseCmp)
	for _, job := range jobs {
		pq.PushBack(job)
	}

	freePoint := uint16(11)
	for pq.Len() > 0 {
		job := pq.PopFront(freePoint)
		if job == nil {
			break
		}
		t.Log(job.String())
		freePoint -= job.CurrTask().RemainPoint()
	}
}

func TestJobPriorityQueueWithBaseCmp(t *testing.T) {
	//
	jobs := []*Job{
		NewJob(1, 7, common.LowPriority, []*Task{
			NewTask(6),
			NewTask(7),
		}),
		NewJob(2, 6, common.HighPriority, []*Task{
			NewTask(5),
		}),
	}

	pq := NewJobPriorityQueue(true, BaseCmp)
	for _, job := range jobs {
		pq.PushBack(job)
	}

	freePoint := uint16(11)
	for pq.Len() > 0 {
		job := pq.PopFront(freePoint)
		if job == nil {
			break
		}
		t.Log(job.String())
		freePoint -= job.CurrTask().RemainPoint()
	}
}

// 不同的创建时间 不同的优先级
func TestJobPriorityQueueWithSimpleCmp_1(t *testing.T) {
	jobs := []*Job{
		NewJob(1, 15, common.LowPriority, []*Task{
			NewTask(6),
			NewTask(7),
		}),
		NewJob(2, 20, common.HighPriority, []*Task{
			NewTask(5),
		}),
	}

	pq := NewJobPriorityQueue(true, SimpleCmp)
	for _, job := range jobs {
		pq.PushBack(job)
	}

	freePoint := uint16(20)
	for pq.Len() > 0 {
		job := pq.PopFront(freePoint)
		if job == nil {
			break
		}
		t.Log(job.String())
		freePoint -= job.CurrTask().RemainPoint()
	}
}

// 相同的创建时间 不同的优先级
func TestJobPriorityQueueWithSimpleCmp_2(t *testing.T) {
	jobs := []*Job{
		NewJob(1, 20, common.LowPriority, []*Task{
			NewTask(6),
			NewTask(7),
		}),
		NewJob(2, 20, common.HighPriority, []*Task{
			NewTask(5),
		}),
	}

	pq := NewJobPriorityQueue(true, SimpleCmp)
	for _, job := range jobs {
		pq.PushBack(job)
	}

	freePoint := uint16(20)
	for pq.Len() > 0 {
		job := pq.PopFront(freePoint)
		if job == nil {
			break
		}
		t.Log(job.String())
		freePoint -= job.CurrTask().RemainPoint()
	}
}

// 相同的创建时间 相同的优先级
func TestJobPriorityQueueWithSimpleCmp_3(t *testing.T) {
	jobs := []*Job{
		NewJob(1, 20, common.HighPriority, []*Task{
			NewTask(6),
			NewTask(7),
		}),
		NewJob(2, 20, common.HighPriority, []*Task{
			NewTask(5),
		}),
	}

	pq := NewJobPriorityQueue(true, SimpleCmp)
	for _, job := range jobs {
		pq.PushBack(job)
	}

	freePoint := uint16(20)
	for pq.Len() > 0 {
		job := pq.PopFront(freePoint)
		if job == nil {
			break
		}
		t.Log(job.String())
		freePoint -= job.CurrTask().RemainPoint()
	}
}

// 不同权重weight 不同的优先级
func TestJobPriorityQueueWithSmartCmp_11(t *testing.T) {
	jobs := []*Job{
		NewJob(1, 11, common.LowPriority, []*Task{
			NewTask(6),
			NewTask(7),
		}),
		NewJob(2, 20, common.HighPriority, []*Task{
			NewTask(5),
		}),
	}

	pq := NewJobPriorityQueue(true, SmartCmp)
	for _, job := range jobs {
		pq.PushBack(job)
	}

	freePoint := uint16(20)
	for pq.Len() > 0 {
		job := pq.PopFront(freePoint)
		if job == nil {
			break
		}
		t.Log(job.String())
		freePoint -= job.CurrTask().RemainPoint()
	}
}

// 不同权重weight 不同的优先级
func TestJobPriorityQueueWithSmartCmp_12(t *testing.T) {
	jobs := []*Job{
		NewJob(1, 9, common.LowPriority, []*Task{
			NewTask(6),
			NewTask(7),
		}),
		NewJob(2, 20, common.HighPriority, []*Task{
			NewTask(5),
		}),
	}

	pq := NewJobPriorityQueue(true, SmartCmp)
	for _, job := range jobs {
		pq.PushBack(job)
	}

	freePoint := uint16(20)
	for pq.Len() > 0 {
		job := pq.PopFront(freePoint)
		if job == nil {
			break
		}
		t.Log(job.String())
		freePoint -= job.CurrTask().RemainPoint()
	}
}

// 相同权重weight 不同的优先级
func TestJobPriorityQueueWithSmartCmp_2(t *testing.T) {
	jobs := []*Job{
		NewJob(1, 10, common.LowPriority, []*Task{
			NewTask(6),
			NewTask(7),
		}),
		NewJob(2, 20, common.HighPriority, []*Task{
			NewTask(5),
		}),
	}

	pq := NewJobPriorityQueue(true, SmartCmp)
	for _, job := range jobs {
		pq.PushBack(job)
	}

	freePoint := uint16(20)
	for pq.Len() > 0 {
		job := pq.PopFront(freePoint)
		if job == nil {
			break
		}
		t.Log(job.String())
		freePoint -= job.CurrTask().RemainPoint()
	}
}

// 相同权重weight 想同的优先级
func TestJobPriorityQueueWithSmartCmp_3(t *testing.T) {
	jobs := []*Job{
		NewJob(1, 10, common.LowPriority, []*Task{
			NewTask(6),
			NewTask(7),
		}),
		NewJob(2, 10, common.LowPriority, []*Task{
			NewTask(5),
		}),
	}

	pq := NewJobPriorityQueue(true, SmartCmp)
	for _, job := range jobs {
		pq.PushBack(job)
	}

	freePoint := uint16(20)
	for pq.Len() > 0 {
		job := pq.PopFront(freePoint)
		if job == nil {
			break
		}
		t.Log(job.String())
		freePoint -= job.CurrTask().RemainPoint()
	}
}
