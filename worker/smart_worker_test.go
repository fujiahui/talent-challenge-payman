package worker

import (
	"context"
	"github.com/fujiahui/talnet-challenge-payman/common"
	"github.com/fujiahui/talnet-challenge-payman/logger"
	"github.com/fujiahui/talnet-challenge-payman/server"
	"sync"
	"testing"
	"time"
)

func getJobArray(created common.TimestampType) *common.JobInfoArray {
	jobArray1 := &common.JobInfoArray{
		JobInfos: []*common.JobInfo{
			{
				ID:       1,
				Created:  1,
				Priority: common.LowPriority,
				Tasks:    []common.PointType{5, 6, 7},
			},
		},
	}

	jobArray2 := &common.JobInfoArray{
		JobInfos: []*common.JobInfo{
			{
				ID:       2,
				Created:  3,
				Priority: common.HighPriority,
				Tasks:    []common.PointType{3, 5},
			},
		},
	}

	jobInfoMap := make(map[common.TimestampType]*common.JobInfoArray)
	jobInfoMap[1] = jobArray1
	jobInfoMap[3] = jobArray2

	if jobArray, ok := jobInfoMap[created]; ok {
		return jobArray
	}

	return nil
}

// Task 1.2
func TestNewBaseWorker(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "../warehouse/data/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := common.TimestampType(-1)
	logger.ChartLogger.Printf("Start TestNewBaseWorker")
	w := NewBaseWorker(startTimestamp)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(20000) * time.Millisecond)
	cancel()
	wg.Wait()
}

// Task 2.1
func TestNewWorkerWithCapacity(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "../warehouse/data/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := common.TimestampType(-1)
	// capacity := common.PointType(10)
	capacity := common.PointType(15)
	logger.ChartLogger.Printf("Start TestNewWorkerWithCapacity capacity=%d", capacity)
	w := NewWorkerWithCapacity(startTimestamp, capacity)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(20000) * time.Millisecond)
	cancel()
	wg.Wait()
}

// Task 2.2
func TestNewWorkerWithSimplePriority(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "../warehouse/data/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := common.TimestampType(-1)
	// capacity := common.PointType(10)
	capacity := common.PointType(6)
	logger.ChartLogger.Printf("Start TestNewWorkerWithSimplePriority capacity=%d", capacity)
	w := NewWorkerWithSimplePriority(startTimestamp, capacity)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(60000) * time.Millisecond)
	cancel()
	wg.Wait()
}

// Task 2.3
func TestNewWorkerWithSmartPriority(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "../warehouse/data/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := common.TimestampType(-1)
	capacity := common.PointType(10)
	// capacity := common.PointType(15)
	logger.ChartLogger.Printf("Start TestNewWorkerWithSmartPriority capacity=%d", capacity)
	w := NewWorkerWithSmartPriority(startTimestamp, capacity)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(5000) * time.Millisecond)
	cancel()
	wg.Wait()
}

// Task 3.1
func TestNewWorkerWithNumPriority(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "../warehouse/data_num_priority/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := common.TimestampType(-1)
	capacity := common.PointType(10)
	logger.ChartLogger.Printf("Start TestNewWorkerWithNumPriority capacity=%d", capacity)
	w := NewWorkerWithSmartPriority(startTimestamp, capacity)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(5000) * time.Millisecond)
	cancel()
	wg.Wait()
}

// Task 3.3
func TestNewWorkerWithTaskSpeed(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "../warehouse/data_num_priority/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := common.TimestampType(-1)
	// capacity := common.PointType(10)
	capacity := common.PointType(15)
	logger.ChartLogger.Printf("Start TestNewWorkerWithTaskSpeed capacity=%d", capacity)
	w := NewWorkerWithSmartPriority(startTimestamp, capacity)
	// w := NewBaseWorker(startTimestamp)
	w.EnableTaskSpeed() // 启用Task任务加速
	// w.DisableTaskSpeed()
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(20000) * time.Millisecond)
	cancel()
	wg.Wait()
}
