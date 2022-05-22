package util

import (
	"fmt"
	"github.com/fujiahui/talnet-challenge-payman/common"
)

type JobStatusType uint8

const (
	JobCreated  = JobStatusType(0)
	JobWait     = JobStatusType(1)
	JobRunning  = JobStatusType(2)
	JobSleep    = JobStatusType(4)
	JobFinished = JobStatusType(8)
)

func BaseCmp(x any, y any) bool {
	j1, j2 := x.(*Job), y.(*Job)
	t1, t2 := j1.CurrTask(), j2.CurrTask()
	if t1.ExpectedTimestamp() < t2.ExpectedTimestamp() {
		return true
	} else if t1.ExpectedTimestamp() > t2.ExpectedTimestamp() {
		return false
	}

	return j1.ID() < j2.ID()
}

func SimpleCmp(x any, y any) bool {
	// 1. 按照期待开始时间进行比较排序
	j1, j2 := x.(*Job), y.(*Job)
	t1, t2 := j1.CurrTask(), j2.CurrTask()
	if t1.ExpectedTimestamp() < t2.ExpectedTimestamp() {
		return true
	} else if t1.ExpectedTimestamp() > t2.ExpectedTimestamp() {
		return false
	}

	// 2. 按照优先级进行比较排序 高优先级先执行
	if j1.Priority() > j2.Priority() {
		return true
	} else if j1.Priority() < j2.Priority() {
		return false
	}

	return j1.ID() < j2.ID()
}

func SmartCmp(x any, y any) bool {
	// 1. 按照Job优先级加权重进行比较排序
	j1, j2 := x.(*Job), y.(*Job)
	if j1.Weight() < j2.Weight() {
		return true
	} else if j1.Weight() > j2.Weight() {
		return false
	}

	// 2. 按照优先级进行比较排序 高优先级先执行
	if j1.Priority() > j2.Priority() {
		return true
	} else if j1.Priority() < j2.Priority() {
		return false
	}

	return j1.ID() < j2.ID()
}

type Job struct {
	id       common.JobIDType
	created  common.TimestampType
	priority common.PriorityType

	curr   int
	tasks  []*Task
	status JobStatusType
}

func NewJob(id common.JobIDType, created common.TimestampType, priority common.PriorityType, tasks []*Task) *Job {
	for i, t := range tasks {
		t.SetJobID(id)
		t.SetTaskID(i + 1)
		t.SetExpectedTime(created)
	}

	return &Job{
		id:       id,
		created:  created,
		priority: priority,
		curr:     0,
		tasks:    tasks,
		status:   JobCreated,
	}
}

func NewJobFromCommon(info *common.JobInfo) *Job {

	tasks := make([]*Task, 0, 16)
	for i, point := range info.Tasks {
		task := NewTask(point)
		task.SetJobID(info.ID)
		task.SetTaskID(i + 1)
		task.SetExpectedTime(info.Created)
		tasks = append(tasks, task)
	}

	return &Job{
		id:       info.ID,
		created:  info.Created,
		priority: info.Priority,
		curr:     0,
		tasks:    tasks,
		status:   JobCreated,
	}
}

func (j *Job) ID() common.JobIDType {
	return j.id
}

func (j *Job) Weight() int64 {
	return int64(j.CurrTask().ExpectedTimestamp()) - int64(j.priority)
}

func (j *Job) Priority() common.PriorityType {
	return j.priority
}

func (j *Job) EnableTaskSpeed() {
	for _, t := range j.tasks {
		t.EnableSpeed()
	}
}

func (j *Job) DisableTaskSpeed() {
	for _, t := range j.tasks {
		t.DisableSpeed()
	}
}

// CurrTask 返回当前正在执行 或 待执行的Task
func (j *Job) CurrTask() *Task {
	/*
		1. 如果Job处于等待或休眠状态, 则返回Job下一个需要执行的Task
		2. 如果Job处于正在运行状态, 则返回当前正在运行的Task
	*/
	if j.curr >= len(j.tasks) {
		return nil
	}
	return j.tasks[j.curr]
}

// NextTask 顺序执行Job中的Tasks
func (j *Job) NextTask(currTimestamp common.TimestampType) {
	j.curr++
	if j.curr == len(j.tasks) {
		j.status = JobFinished
		return
	}
	j.CurrTask().SetExpectedTime(currTimestamp)
	j.status = JobWait
	return
}

func (j *Job) Status() JobStatusType {
	return j.status
}

func (j *Job) Finished() bool {
	return j.status == JobFinished
}

func (j *Job) SetRunning() {
	j.status = JobRunning
}

func (j *Job) SetSleep() {
	j.status = JobSleep
}

func (j *Job) SetWait() {
	j.status = JobWait
}

func (j *Job) String() string {
	if j.curr == len(j.tasks) {
		return ""
	}

	t := j.CurrTask()
	return fmt.Sprintf("%d-%s", j.id, t.String())
}

func (j *Job) StringWithPriority() string {
	if j.curr == len(j.tasks) {
		return ""
	}

	t := j.CurrTask()
	return fmt.Sprintf("%d(%d)-%s", j.id, j.Weight(), t.String())
}
