package worker

import (
	"fmt"
	"github.com/fujiahui/talnet-challenge-payman/common"
	"github.com/fujiahui/talnet-challenge-payman/worker/util"
	"strings"
	"time"
)

// Actuator Job执行管理器
type Actuator struct {
	currTimestamp int64               // 当前时间戳
	jobs          map[int64]*util.Job // ID <--> *Job

	// 当前分配的任务，剩余的Points数
	capacity       common.PointType
	executingPoint common.PointType
}

func NewActuator(startTimestamp int64, capacity common.PointType) *Actuator {
	return &Actuator{
		currTimestamp:  startTimestamp,
		jobs:           make(map[int64]*util.Job),
		capacity:       capacity,
		executingPoint: 0,
	}
}

func (c *Actuator) ExecutingPoint() common.PointType {
	return c.executingPoint
}

func (c *Actuator) CurrTimestamp() int64 {
	return c.currTimestamp
}

func (c *Actuator) FreePoint() common.PointType {
	return c.capacity - c.executingPoint
}

// Ticking 滴答滴答向前一步步大胆的滴答
func (c *Actuator) Ticking(tick int) []*util.Job {
	time.Sleep(time.Duration(tick) * time.Millisecond)
	c.currTimestamp++

	ids := make([]int64, 0, 16)
	jobs := make([]*util.Job, 0, 16)
	for id, job := range c.jobs {
		c.executingPoint -= 1
		t := job.CurrTask()
		t.Ticking()
		if t.Finished() {
			ids = append(ids, id)
			job.NextTask(c.currTimestamp)
			if !job.Finished() {
				jobs = append(jobs, job)
			}
		}
	}
	for _, id := range ids {
		delete(c.jobs, id)
	}

	return jobs
}

// Execute 执行一个Job
func (c *Actuator) Execute(job *util.Job) {
	c.jobs[job.ID()] = job
	job.SetRunning()

	t := job.CurrTask()
	t.SetRunning()

	c.executingPoint += t.TaskPoint()
	return
}

func (c *Actuator) String() string {
	ss := make([]string, 0, 16)
	for _, job := range c.jobs {
		ss = append(ss, job.String())
	}

	tt := make([]string, 3, 4)
	currTimestamp := int(c.currTimestamp)
	for i := 2; i >= 0; i-- {
		tt[i] = fmt.Sprintf("%.2d", currTimestamp%60)
		currTimestamp /= 60
	}

	return fmt.Sprintf("%s | %s | %d",
		strings.Join(tt, ":"),
		strings.Join(ss, "|"),
		c.executingPoint)
}
