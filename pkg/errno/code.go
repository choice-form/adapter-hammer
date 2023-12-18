package errno

var (
	ErrorAccessToken      = &APIError{ErrCode: 401, Errmsg: "不合法的 access_token"}
	ErrorCreateCredential = &APIError{ErrCode: 400, Errmsg: "创建连接器失败，请检查参数是否正确"}
	ErrorUpdateCredential = &APIError{ErrCode: 400, Errmsg: "更新配置错误，请检查参数是否正确"}
	ErrorDeleteCredential = &APIError{ErrCode: 403, Errmsg: "拒绝执行，此连接器删除失败"}
	ErrorParams           = &APIError{ErrCode: 400, Errmsg: "请求参数错误"}
	ErrorMethod           = &APIError{ErrCode: 404, Errmsg: "method 未找到"}
	// ErrorParams      = &APIError{ErrCode: 403, Errmsg: "没有调用该接口的权限"}
	ErrorOther = &APIError{ErrCode: 500, Errmsg: "未知错误"} // 服务自身错误

	ErrorInitSubscription = &APIError{ErrCode: 501, Errmsg: "初始化事件订阅错误"} // 初始化事件订阅服务错误
)
