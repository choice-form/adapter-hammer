package errno

var (
	ErrorAccessToken              = &APIError{HttpCode: 401, ErrCode: 401, Errmsg: "不合法的 access_token"}
	ErrorCreateCredential         = &APIError{HttpCode: 400, ErrCode: 400, Errmsg: "创建连接器失败，请检查参数是否正确"}
	ErrorUpdateCredential         = &APIError{HttpCode: 400, ErrCode: 400, Errmsg: "更新配置错误，请检查参数是否正确"}
	ErrorCreateOrUpdateCredential = &APIError{HttpCode: 400, ErrCode: 400, Errmsg: "创建或更新连接器失败，请检查参数是否正确"}
	ErrorDeleteCredential         = &APIError{HttpCode: 400, ErrCode: 403, Errmsg: "拒绝执行，此连接器删除失败"}
	ErrorNotFoundCredential       = &APIError{HttpCode: 404, ErrCode: 404, Errmsg: "连接器不存在"}
	ErrorParams                   = &APIError{HttpCode: 400, ErrCode: 400, Errmsg: "请求参数错误"}
	ErrorMethod                   = &APIError{HttpCode: 400, ErrCode: 400, Errmsg: "method 未找到"}
	ErrorOther                    = &APIError{ErrCode: 500, Errmsg: "未知错误"} // 服务自身错误
	// ErrorParams      = &APIError{ErrCode: 403, Errmsg: "没有调用该接口的权限"}

	ErrorInitSubscription = &APIError{HttpCode: 500, ErrCode: 511, Errmsg: "初始化事件订阅错误"} // 初始化事件订阅服务错误
	ErrorGetSubscription  = &APIError{HttpCode: 404, ErrCode: 512, Errmsg: "查询事件订阅错误"}  // 查询事件订阅错误
)
