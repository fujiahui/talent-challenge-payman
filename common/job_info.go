package common

import (
	"bufio"
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

type JobInfo struct {
	ID       int64        `json:"JobID"`
	Created  int64        `json:"Created"`
	Priority PriorityType `json:"Priority"`
	Tasks    []uint16     `json:"Tasks"`
}

func NewJobInfo(filename string) *JobInfo {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	job := JobInfo{
		Tasks: make([]uint16, 0, 8),
	}
	var sectionName string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		if line[0] == '[' && line[len(line)-1] == ']' {
			sectionName = line[1 : len(line)-1]
			sectionName = strings.TrimSpace(sectionName)
		} else {
			switch sectionName {
			case "JobID":
				id, err := strconv.ParseInt(line, 0, 64)
				if err != nil {
					return nil
				}
				job.ID = id
			case "Created":
				tt := strings.Split(line, ":")
				created := int64(0)
				for _, t := range tt {
					it, err := strconv.ParseInt(t, 0, 64)
					if err != nil {
						return nil
					}
					created *= 60
					created += it
				}
				job.Created = created
			case "Priority":
				if strings.EqualFold(line, "Low") {
					job.Priority = LowPriority
					continue
				} else if strings.EqualFold(line, "High") {
					job.Priority = HighPriority
					continue
				}
				priority, err := strconv.ParseUint(line, 0, 8)
				if err != nil || priority < 0 || priority > 100 {
					return nil
				}
				job.Priority = PriorityType(priority)
			case "Tasks":
				task, err := strconv.ParseUint(line, 0, 16)
				if err != nil {
					continue
				}
				job.Tasks = append(job.Tasks, uint16(task))
			}
		}
	}

	return &job
}

func (info *JobInfo) ToJsonString() string {
	b, _ := json.Marshal(info)
	return string(b)
}

func (info *JobInfo) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &info)
}

type JobInfoArray struct {
	JobInfos []*JobInfo `json:"JobInfos"`
}

func NewJobInfoArray() *JobInfoArray {
	return &JobInfoArray{
		JobInfos: make([]*JobInfo, 0, 4),
	}
}

func (array *JobInfoArray) ToJsonString() string {
	b, _ := json.Marshal(array)
	return string(b)
}

func (array *JobInfoArray) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &array)
}
