package httptools

import (
	"encoding/json"
	"net/http"

	"github.com/choice-form/adapter-hammer/pkg/errno"
	"github.com/gin-gonic/gin"
)

type Data map[string]any

type APIResponse struct {
	ErrCode int64  `json:"err_code,omitempty"`
	Errmsg  string `json:"errmsg,omitempty"`
	Err     any    `json:"error,omitempty"`
	Data
}

func (e *APIResponse) MarshalJSON() ([]byte, error) {
	_resp := map[string]any{
		"errcode": e.ErrCode,
		"errmsg":  e.Errmsg,
	}
	// 处理错误信息
	if e.Err != nil {
		if _e, ok := e.Err.(*errno.APIError); ok {
			_resp["error"] = _e.Error()
		} else {
			_j, _ := json.Marshal(e.Err)
			_resp["error"] = string(_j)
		}
		return json.Marshal(_resp)
	}

	// 处理正确信息
	for k, v := range e.Data {
		_resp[k] = v
	}
	if e.Err != nil {
		// 这里后台打印错误信息
	}
	return json.Marshal(_resp)
}

func OKResponse(c *gin.Context, data map[string]any) {
	c.JSON(http.StatusOK, &APIResponse{
		ErrCode: 0,
		Errmsg:  "ok",
		Data:    data,
	})
}

func ErrResponse(c *gin.Context, err error) {
	if _err, ok := err.(*errno.APIError); ok {
		if _e, ok := _err.Err.(error); ok {
			c.JSON(http.StatusOK, &APIResponse{
				ErrCode: _err.ErrCode,
				Errmsg:  _err.Errmsg,
				Err:     _e.Error(),
			})
		} else {
			c.JSON(http.StatusOK, &APIResponse{
				ErrCode: _err.ErrCode,
				Errmsg:  _err.Errmsg,
				Err:     _err.Err,
			})
		}
		return
	}

	c.JSON(http.StatusOK, &APIResponse{
		ErrCode: 500,
		Errmsg:  err.Error(),
		Err:     err,
	})

}
