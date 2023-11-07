package httptools

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/choice-form/adapter-hammer/pkg/errno"
)

func TestOKResponse(t *testing.T) {
	t.Run("test ok", func(t *testing.T) {
		data := map[string]any{
			"id":     "Forbidden.AccessDenied.AccessTokenPermissionDenied",
			"result": "ok",
		}
		_result := &APIResponse{
			ErrCode: 0,
			Errmsg:  "ok",
			Data:    data,
		}
		_b, _ := json.Marshal(_result)
		t.Log(string(_b))
	})
}

func TestErrResponse(t *testing.T) {
	t.Run("test error", func(t *testing.T) {
		_e := &errno.APIError{
			ErrCode: 403,
			Errmsg:  "fiber",
			Err:     nil,
		}
		_result := &APIResponse{
			ErrCode: _e.ErrCode,
			Errmsg:  _e.Errmsg,
			Err:     _e,
		}
		_b, _ := json.Marshal(_result)
		t.Log(string(_b))
	})

	t.Run("test error has Err", func(t *testing.T) {
		_e := &errno.APIError{
			ErrCode: 403,
			Errmsg:  "fiber",
			Err:     errors.New("not found params"),
		}
		_result := &APIResponse{
			ErrCode: _e.ErrCode,
			Errmsg:  _e.Errmsg,
			Err:     _e,
		}
		_b, _ := json.Marshal(_result)
		t.Log(string(_b))
	})

	t.Run("test error has ErrMap", func(t *testing.T) {
		_e := &errno.APIError{
			ErrCode: 403,
			Errmsg:  "fiber",
			Err: map[string]any{
				"code":      "InvalidAuthentication",
				"requestid": "48AD5A3F-713B-7689-AC2A-02C9CD045D03",
				"message":   "不合法的access_token",
			},
		}
		_result := &APIResponse{
			ErrCode: _e.ErrCode,
			Errmsg:  _e.Errmsg,
			Err:     _e,
		}
		_b, _ := json.Marshal(_result)
		t.Log(string(_b))
	})
}
