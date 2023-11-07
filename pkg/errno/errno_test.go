package errno

import (
	"errors"
	"testing"
)

func TestAPIError_Error(t *testing.T) {
	t.Run("error has not Err", func(t *testing.T) {
		_e := &APIError{
			ErrCode: 200,
			Errmsg:  "success",
			Err:     nil,
		}
		t.Log(_e)
	})
}

func TestAPIError_String(t *testing.T) {
	t.Run("error has Err", func(t *testing.T) {
		_e := &APIError{
			ErrCode: 500,
			Errmsg:  "accessToken error",
			Err:     errors.New("not found accessToken"),
		}

		t.Log(_e)
	})

	t.Run("error has map Err", func(t *testing.T) {
		_e := &APIError{
			ErrCode: 500,
			Errmsg:  "accessToken error",
			Err: map[string]any{
				"code":      "InvalidAuthentication",
				"requestid": "48AD5A3F-713B-7689-AC2A-02C9CD045D03",
				"message":   "不合法的access_token",
			},
		}
		t.Log(_e)
	})
}
