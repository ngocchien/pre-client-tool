package curl

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type curl struct {
	userAgent string
}

func NewCurl(userAgent string) *curl {
	return &curl{
		userAgent: userAgent,
	}
}

type Params struct {
	Url     string
	Method  string
	Timeout int
	Queries map[string]interface{}
	Headers map[string]string
	Body    map[string]interface{}
}

func (c *curl) Execute(params Params) (bool, string) {
	url := params.Url
	if len(params.Queries) > 0 {
		queries := "&"
		for key, element := range params.Queries {
			queries += key + "=" + fmt.Sprint(element)
		}
		if strings.Contains("?", url) {
			url += queries
		} else {
			url += "?" + queries
		}
	}

	var req *http.Request
	if params.Body != nil {
		var body, _ = json.Marshal(params.Body)
		reqTemp, _ := http.NewRequest(params.Method, url, bytes.NewBuffer(body))
		req = reqTemp
	} else {
		reqTemp, _ := http.NewRequest(params.Method, url, nil)
		req = reqTemp
	}

	clientHttp := &http.Client{
		Timeout: 10 * time.Second,
	}

	if len(params.Headers) > 0 {
		for keyName, value := range params.Headers {
			req.Header.Set(keyName, value)
		}
	}

	resp, _ := clientHttp.Do(req)

	if resp == nil {
		return false, ""
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	return true, string(bodyBytes)
}
