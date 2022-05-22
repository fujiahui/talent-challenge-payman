package worker

import (
	"context"
	"fmt"
	"github.com/fujiahui/talnet-challenge-payman/common"
	"github.com/fujiahui/talnet-challenge-payman/worker/util"
	"sync"
)

type handler func(created int64) *common.JobInfoArray

type SmartWorker struct {
	// Job调度管理器
	scheduler *Scheduler
	// Job执行管理器
	actuator *Actuator

	speedFlag bool
}

// NewBaseWorker Task 1.2
func NewBaseWorker(startTimestamp int64) *SmartWorker {
	capacity := uint16(1 << 15) // 16位最大整数 == 不限制容量
	return &SmartWorker{
		scheduler: NewScheduler(false, util.BaseCmp),
		actuator:  NewActuator(startTimestamp, capacity),
	}
}

// NewWorkerWithCapacity Task 2.1
func NewWorkerWithCapacity(startTimestamp int64, capacity uint16) *SmartWorker {
	return &SmartWorker{
		scheduler: NewScheduler(false, util.BaseCmp),
		actuator:  NewActuator(startTimestamp, capacity),
	}
}

// NewWorkerWithSimplePriority Task 2.2
func NewWorkerWithSimplePriority(startTimestamp int64, capacity uint16) *SmartWorker {
	return &SmartWorker{
		scheduler: NewScheduler(true, util.SimpleCmp),
		actuator:  NewActuator(startTimestamp, capacity),
	}
}

// NewWorkerWithSmartPriority Task 2.3
func NewWorkerWithSmartPriority(startTimestamp int64, capacity uint16) *SmartWorker {
	return &SmartWorker{
		scheduler: NewScheduler(true, util.SmartCmp),
		actuator:  NewActuator(startTimestamp, capacity),
	}
}

func (w *SmartWorker) EnableTaskSpeed() {
	w.speedFlag = true
}

func (w *SmartWorker) DisableTaskSpeed() {
	w.speedFlag = false
}

func (w *SmartWorker) Start(ctx context.Context, h handler) {
	wg := &sync.WaitGroup{}

	tick := 100
	for {
		// 0. 按照时间进行迭代图标
		jobs := w.actuator.Ticking(tick)

		// 1. 每隔1ms / 1s 获取一批次的job列表
		if jobArray := h(w.actuator.CurrTimestamp()); jobArray != nil {
			for _, info := range jobArray.JobInfos {
				job := util.NewJobFromCommon(info)
				if job == nil {
					continue
				}

				if w.speedFlag {
					job.EnableTaskSpeed()
				} else {
					job.DisableTaskSpeed()
				}
				jobs = append(jobs, job)
			}
		}

		// 2. 把jobs放入优先队列中
		for _, job := range jobs {
			w.scheduler.Enqueue(job)
		}

		// 3. 分发Task
		for {
			freePoint := w.actuator.FreePoint()
			job := w.scheduler.Dequeue(freePoint)
			if job == nil {
				break
			}
			w.actuator.Execute(job)

			wg.Add(1)
			go func(t *util.Task) {
				defer wg.Done()
				// t.Running(tick)
			}(job.CurrTask())

		}

		fmt.Println(w.actuator.String())
		if ctx.Err() != nil {
			break
		}

	}

	wg.Wait()
}
