package server

import (
	"github.com/fujiahui/talnet-challenge-payman/common"
	"testing"
)

func TestDataHubServer_GetJobInfo(t *testing.T) {
	dirPath := "../warehouse/data/"
	svr := NewDataHubServer(dirPath)
	for i := common.TimestampType(0); i < 1000; i++ {
		jobArray := svr.GetJobInfo(i)
		if jobArray != nil {
			t.Log(jobArray.ToJsonString())
		}
	}
}
