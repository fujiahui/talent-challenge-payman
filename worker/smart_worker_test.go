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

const (
	TimeOut = int64(20000)
)

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

	time.Sleep(time.Duration(TimeOut) * time.Millisecond)
	cancel()
	wg.Wait()
}

// Task 2.1
func TestNewWorkerWithCapacity_0(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "../warehouse/data/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := common.TimestampType(-1)
	capacity := common.PointType(0)
	logger.ChartLogger.Printf("Start TestNewWorkerWithCapacity capacity=%d", capacity)
	w := NewWorkerWithCapacity(startTimestamp, capacity)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(TimeOut) * time.Millisecond)
	cancel()
	wg.Wait()
}

func TestNewWorkerWithCapacity_6(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "../warehouse/data/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := common.TimestampType(-1)
	capacity := common.PointType(6)
	logger.ChartLogger.Printf("Start TestNewWorkerWithCapacity capacity=%d", capacity)
	w := NewWorkerWithCapacity(startTimestamp, capacity)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(TimeOut) * time.Millisecond)
	cancel()
	wg.Wait()
}

func TestNewWorkerWithCapacity_10(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "../warehouse/data/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := common.TimestampType(-1)
	capacity := common.PointType(10)
	logger.ChartLogger.Printf("Start TestNewWorkerWithCapacity capacity=%d", capacity)
	w := NewWorkerWithCapacity(startTimestamp, capacity)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(TimeOut) * time.Millisecond)
	cancel()
	wg.Wait()
}

func TestNewWorkerWithCapacity_15(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "../warehouse/data/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := common.TimestampType(-1)
	capacity := common.PointType(15)
	logger.ChartLogger.Printf("Start TestNewWorkerWithCapacity capacity=%d", capacity)
	w := NewWorkerWithCapacity(startTimestamp, capacity)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(TimeOut) * time.Millisecond)
	cancel()
	wg.Wait()
}

// Task 2.2
func TestNewWorkerWithSimplePriority_0(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "../warehouse/data/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := common.TimestampType(-1)
	capacity := common.PointType(0)
	logger.ChartLogger.Printf("Start TestNewWorkerWithSimplePriority capacity=%d", capacity)
	w := NewWorkerWithSimplePriority(startTimestamp, capacity)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(TimeOut) * time.Millisecond)
	cancel()
	wg.Wait()
}

func TestNewWorkerWithSimplePriority_6(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "../warehouse/data/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := common.TimestampType(-1)
	capacity := common.PointType(6)
	logger.ChartLogger.Printf("Start TestNewWorkerWithSimplePriority capacity=%d", capacity)
	w := NewWorkerWithSimplePriority(startTimestamp, capacity)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(TimeOut) * time.Millisecond)
	cancel()
	wg.Wait()
}

func TestNewWorkerWithSimplePriority_10(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "../warehouse/data/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := common.TimestampType(-1)
	capacity := common.PointType(10)
	logger.ChartLogger.Printf("Start TestNewWorkerWithSimplePriority capacity=%d", capacity)
	w := NewWorkerWithSimplePriority(startTimestamp, capacity)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(TimeOut) * time.Millisecond)
	cancel()
	wg.Wait()
}

func TestNewWorkerWithSimplePriority_15(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "../warehouse/data/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := common.TimestampType(-1)
	capacity := common.PointType(15)
	logger.ChartLogger.Printf("Start TestNewWorkerWithSimplePriority capacity=%d", capacity)
	w := NewWorkerWithSimplePriority(startTimestamp, capacity)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(TimeOut) * time.Millisecond)
	cancel()
	wg.Wait()
}

// Task 2.3
func TestNewWorkerWithSmartPriority_0(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "../warehouse/data/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := common.TimestampType(-1)
	capacity := common.PointType(0)
	logger.ChartLogger.Printf("Start TestNewWorkerWithSmartPriority capacity=%d", capacity)
	w := NewWorkerWithSmartPriority(startTimestamp, capacity)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(TimeOut) * time.Millisecond)
	cancel()
	wg.Wait()
}

func TestNewWorkerWithSmartPriority_6(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "../warehouse/data/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := common.TimestampType(-1)
	capacity := common.PointType(6)
	logger.ChartLogger.Printf("Start TestNewWorkerWithSmartPriority capacity=%d", capacity)
	w := NewWorkerWithSmartPriority(startTimestamp, capacity)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(TimeOut) * time.Millisecond)
	cancel()
	wg.Wait()
}

func TestNewWorkerWithSmartPriority_10(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "../warehouse/data/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := common.TimestampType(-1)
	capacity := common.PointType(10)
	logger.ChartLogger.Printf("Start TestNewWorkerWithSmartPriority capacity=%d", capacity)
	w := NewWorkerWithSmartPriority(startTimestamp, capacity)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(TimeOut) * time.Millisecond)
	cancel()
	wg.Wait()
}

func TestNewWorkerWithSmartPriority_15(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "../warehouse/data/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := common.TimestampType(-1)
	capacity := common.PointType(15)
	logger.ChartLogger.Printf("Start TestNewWorkerWithSmartPriority capacity=%d", capacity)
	w := NewWorkerWithSmartPriority(startTimestamp, capacity)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(TimeOut) * time.Millisecond)
	cancel()
	wg.Wait()
}

// Task 3.1
func TestNewWorkerWithNumPriority_0(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "../warehouse/data_num_priority/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := common.TimestampType(-1)
	capacity := common.PointType(0)
	logger.ChartLogger.Printf("Start TestNewWorkerWithNumPriority capacity=%d", capacity)
	w := NewWorkerWithSmartPriority(startTimestamp, capacity)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(TimeOut) * time.Millisecond)
	cancel()
	wg.Wait()
}

func TestNewWorkerWithNumPriority_6(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "../warehouse/data_num_priority/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := common.TimestampType(-1)
	capacity := common.PointType(6)
	logger.ChartLogger.Printf("Start TestNewWorkerWithNumPriority capacity=%d", capacity)
	w := NewWorkerWithSmartPriority(startTimestamp, capacity)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(TimeOut) * time.Millisecond)
	cancel()
	wg.Wait()
}

func TestNewWorkerWithNumPriority_10(t *testing.T) {
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

	time.Sleep(time.Duration(TimeOut) * time.Millisecond)
	cancel()
	wg.Wait()
}

func TestNewWorkerWithNumPriority_15(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "../warehouse/data_num_priority/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := common.TimestampType(-1)
	capacity := common.PointType(15)
	logger.ChartLogger.Printf("Start TestNewWorkerWithNumPriority capacity=%d", capacity)
	w := NewWorkerWithSmartPriority(startTimestamp, capacity)
	wg.Add(1)
	go func() {
		defer wg.Done()
		// w.Start(ctx, getJobArray)
		w.Start(ctx, dataHub.GetJobInfo)
	}()

	time.Sleep(time.Duration(TimeOut) * time.Millisecond)
	cancel()
	wg.Wait()
}

// Task 3.3
func TestNewWorkerWithTaskSpeed_0(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "../warehouse/data_num_priority/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := common.TimestampType(-1)
	capacity := common.PointType(0)
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

	time.Sleep(time.Duration(TimeOut) * time.Millisecond)
	cancel()
	wg.Wait()
}

func TestNewWorkerWithTaskSpeed_6(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "../warehouse/data_num_priority/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := common.TimestampType(-1)
	capacity := common.PointType(6)
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

	time.Sleep(time.Duration(TimeOut) * time.Millisecond)
	cancel()
	wg.Wait()
}

func TestNewWorkerWithTaskSpeed_10(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "../warehouse/data_num_priority/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := common.TimestampType(-1)
	capacity := common.PointType(10)
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

	time.Sleep(time.Duration(TimeOut) * time.Millisecond)
	cancel()
	wg.Wait()
}

func TestNewWorkerWithTaskSpeed_15(t *testing.T) {
	root := context.Background()
	ctx, cancel := context.WithCancel(root)
	wg := &sync.WaitGroup{}

	dirPath := "../warehouse/data_num_priority/"
	dataHub := server.NewDataHubServer(dirPath)

	startTimestamp := common.TimestampType(-1)
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

	time.Sleep(time.Duration(TimeOut) * time.Millisecond)
	cancel()
	wg.Wait()
}
