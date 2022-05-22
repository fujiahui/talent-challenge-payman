package worker

import (
	"github.com/fujiahui/talnet-challenge-payman/common"
	"github.com/fujiahui/talnet-challenge-payman/worker/util"
	"testing"
)

func TestActuator_Ticking(t *testing.T) {
	// timestamp <---> Job
	jobMap := make(map[int64][]*util.Job)
	jobMap[1] = []*util.Job{
		util.NewJob(1, 1, common.LowPriority, []*util.Task{
			util.NewTask(5),
			util.NewTask(6),
			util.NewTask(7),
		}),
	}

	jobMap[3] = []*util.Job{
		util.NewJob(2, 3, common.HighPriority, []*util.Task{
			util.NewTask(3),
			util.NewTask(5),
		}),
	}

	actuator := NewActuator(0, 10)

	for {
		waitJobs := actuator.Ticking(100)
		currTimestamp := actuator.CurrTimestamp()

		for _, job := range waitJobs {
			jobs, ok := jobMap[currTimestamp]
			if ok {
				jobs = append(jobs, job)
			} else {
				jobs = make([]*util.Job, 0, 16)
				jobs = append(jobs, job)
			}
			jobMap[currTimestamp] = jobs
		}

		if sleepJobs, ok := jobMap[currTimestamp]; ok {
			for _, job := range sleepJobs {
				actuator.Execute(job)
			}
			delete(jobMap, currTimestamp)
		}

		t.Log(actuator.String())

		if len(jobMap) == 0 && len(actuator.jobs) == 0 {
			t.Log("End TestActuator_Ticking")
			break
		}
	}
}
