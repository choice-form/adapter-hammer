package errno

import (
	"encoding/json"
	"strconv"
)

type APIError struct {
	HttpCode int
	ErrCode  int64  `json:"err_code,omitempty"`
	Errmsg   string `json:"errmsg,omitempty"`
	Err      any    `json:"err,omitempty"`
}

func (e *APIError) SetCode(code any) *APIError {
	if code != nil {
		switch code.(type) {
		case int:
			e.ErrCode = int64(code.(int))
		case int32:
			e.ErrCode = int64(code.(int32))
		case int64:
			e.ErrCode = int64(code.(int64))
		case float32:
			e.ErrCode = int64(code.(float32))
		case float64:
			e.ErrCode = int64(code.(float64))
		case string:
			_code, _err := strconv.Atoi(code.(string))
			if _err == nil {
				e.ErrCode = int64(_code)
			}
		}
	}
	return e
}

func (e *APIError) SetMessage(msg string) *APIError {
	if msg != "" {
		e.Errmsg = msg
	}
	return e
}

func (e *APIError) SetErr(err any) *APIError {
	e.Err = err
	return e
}

func (e *APIError) Error() string {
	if e.Err != nil {
		if _e, ok := e.Err.(error); ok {
			return _e.Error()
		}
		_j, _ := json.Marshal(e.Err)
		return string(_j)
	} else {
		return e.Errmsg
	}
}

func (e *APIError) String() string {
	return e.Error()
}

func (e *APIError) ToString() string {
	return e.Error()
}
