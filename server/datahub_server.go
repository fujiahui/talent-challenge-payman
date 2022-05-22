package server

import (
	"github.com/fujiahui/talnet-challenge-payman/common"
	"io/ioutil"
	"path"
	"strings"
)

type DataHubServer struct {
	dirPath    string
	jobInfoMap map[common.TimestampType]*common.JobInfoArray
}

func NewDataHubServer(dirPath string) *DataHubServer {
	return newDataHubServer(dirPath)
}

func newDataHubServer(dirPath string) *DataHubServer {
	server := &DataHubServer{
		dirPath:    dirPath,
		jobInfoMap: make(map[common.TimestampType]*common.JobInfoArray),
	}

	filePaths := getFilePaths(dirPath)
	for _, filePath := range filePaths {
		jobInfo := common.NewJobInfo(filePath)
		if jobInfo == nil {
			continue
		}

		jobArray, ok := server.jobInfoMap[jobInfo.Created]
		if !ok {
			jobArray = common.NewJobInfoArray()
			server.jobInfoMap[jobInfo.Created] = jobArray
		}
		jobArray.JobInfos = append(jobArray.JobInfos, jobInfo)
	}

	return server
}

func (s *DataHubServer) GetJobInfo(created common.TimestampType) *common.JobInfoArray {
	jobArray, ok := s.jobInfoMap[created]
	if !ok {
		return nil
	}

	return jobArray
}

func getFilePaths(dirPath string) []string {
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		return nil
	}

	filenames := make([]string, 0, 64)
	for _, file := range dir {
		if file.IsDir() {
			continue
		}

		if strings.HasSuffix(file.Name(), ".job") {
			filenames = append(filenames, path.Join(dirPath, file.Name()))
		}
	}

	return filenames
}
