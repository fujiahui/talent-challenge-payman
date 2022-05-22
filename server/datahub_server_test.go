package server

import (
	"github.com/fujiahui/talnet-challenge-payman/common"
	"testing"
)

func TestDataHubServer_GetJobInfo(t *testing.T) {
	dirPath := "/Users/fujiahui/go-workspace/talent-challenge-payman/warehouse/data/"
	svr := NewDataHubServer(dirPath)
	for i := common.TimestampType(0); i < 100; i++ {
		jobArray := svr.GetJobInfo(i)
		if jobArray != nil {
			t.Log(jobArray.ToJsonString())
		}
	}
}
