package httpUtil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}

func Put(url string, data interface{}, contentType string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	jsonStr, _ := json.Marshal(data)
	fmt.Println(string(jsonStr))
	payload := strings.NewReader(string(jsonStr))
	req, _ := http.NewRequest("PUT", url, payload)
	req.Header.Add("Content-Type", contentType)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}

func Delete(url string) string {
	// 超时时间：5秒
	client := &http.Client{Timeout: 5 * time.Second}
	req, _ := http.NewRequest("DELETE", url, nil)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}
