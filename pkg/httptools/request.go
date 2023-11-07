package httptools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Response struct {
	Headers    http.Header
	StatusCode int
	Body       map[string]any
}

type Request struct {
	Method  string
	Url     string
	Input   map[string]any
	Headers map[string]string
}

func request(_req *http.Request) (*Response, error) {
	client := http.DefaultClient
	resp, err := client.Do(_req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	_resp := &Response{
		Headers:    resp.Header,
		StatusCode: resp.StatusCode,
		Body:       result,
	}

	return _resp, nil
}

func JsonPost(req *Request) (*Response, error) {
	req.Method = "POST"
	buf, _ := json.Marshal(req.Input)
	_b := bytes.NewBuffer(buf)
	_req, err := http.NewRequest(req.Method, req.Url, _b)
	if err != nil {
		return nil, err
	}

	_req.Header.Set("Content-Type", "application/json")

	if req.Headers != nil {
		for k, v := range req.Headers {
			_req.Header.Add(k, v)
		}
	}

	return request(_req)
}

func JsonGet(req *Request) (*Response, error) {
	req.Method = "GET"
	_req, err := http.NewRequest(req.Method, req.Url, nil)
	if err != nil {
		return nil, err
	}

	// add params
	if req.Input != nil {
		var _query url.Values
		for k, v := range req.Input {
			_query := _req.URL.Query()
			_query.Set(k, fmt.Sprint(v))
		}
		//如果参数中有中文参数,这个方法会进行URLEncode
		_req.URL.RawQuery = _query.Encode()
	}

	_req.Header.Set("Content-Type", "application/json")

	// add headers
	if req.Headers != nil {
		for k, v := range req.Headers {
			_req.Header.Add(k, v)
		}
	}

	return request(_req)
}
