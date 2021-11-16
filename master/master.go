package master

import (
	"encoding/json"
	"fmt"
	"github.com/ngocchien/presearch-tool/curl"
	"net/http"
)

type Master struct {
	host string
}

type (
	Data struct {
	}
	Task struct {
		DailyTaskId string
		AccountId   string
		Cookie      string
	}
	apiResponse struct {
		Data struct {
			Rows []Task
		}
		Status bool
	}
)

const (
	TaskStatusDone = 3
)

var userAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 13_2_3 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.3 Mobile/15E148 Safari/604.1"

func NewMaster(host string) *Master {
	return &Master{
		host: host,
	}
}

func (m Master) GetTask(limit int) []Task {
	c := curl.NewCurl(userAgent)
	url := m.host + "/daily-task"
	status, body := c.Execute(curl.Params{
		Url:     url,
		Method:  http.MethodGet,
		Timeout: 10,
		Queries: map[string]interface{}{
			"limit": limit,
		},
		Headers: nil,
		Body:    nil,
	})

	if !status || body == "" {
		return []Task{}
	}

	var result apiResponse
	if err := json.Unmarshal([]byte(body), &result); err != nil { // Parse []byte to the go struct pointer
		fmt.Println("Can not unmarshal JSON")
	}
	return result.Data.Rows
}

func (m Master) UpdateTask(taskId string, data map[string]interface{}) bool {
	c := curl.NewCurl(userAgent)
	url := m.host + "/daily-task/" + taskId
	status, body := c.Execute(curl.Params{
		Url:     url,
		Method:  http.MethodPut,
		Timeout: 10,
		Queries: nil,
		Headers: nil,
		Body:    data,
	})

	if !status || body == "" {
		return false
	}

	return true
}
