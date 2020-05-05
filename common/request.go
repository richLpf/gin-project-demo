package common

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//PublicGet get请求
func PublicGet(url string, key string, value string) (str string, err error) {
	req, _ := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	if key != "" && value != "" {
		req.Header.Set(key, value)
	}
	req.Header.Set("Content-Type", "application/json")
	req.URL.RawQuery = q.Encode()
	resp, err := (&http.Client{}).Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	str = string(body)
	return str, err
}

// PublicPost post请求
func PublicPost(url string, data interface{}, headerParams ...string) (content interface{}, err error) {
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	q := req.URL.Query()
	for index, param := range headerParams {
		if index == 0 {
			req.Header.Add("content-type", param)
		}
		if index == 1 {
			req.Header.Add("remote_user", param)
		}
	}
	req.URL.RawQuery = q.Encode()
	if err != nil {
		return content, err
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 0}
	resp, err := client.Do(req)
	if err != nil {
		return content, err
	}
	defer resp.Body.Close()
	result, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal([]byte(string(result)), &content)
	return content, err
}
