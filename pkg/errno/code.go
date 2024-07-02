package errno

var (
	ErrorAccessToken              = &APIError{HttpCode: 401, ErrCode: 4011, Errmsg: "不合法的 access_token"}
	ErrorCreateCredential         = &APIError{HttpCode: 400, ErrCode: 4001, Errmsg: "创建连接器失败，请检查参数是否正确"}
	ErrorUpdateCredential         = &APIError{HttpCode: 400, ErrCode: 4002, Errmsg: "更新配置错误，请检查参数是否正确"}
	ErrorCreateOrUpdateCredential = &APIError{HttpCode: 400, ErrCode: 4003, Errmsg: "创建或更新连接器失败，请检查参数是否正确"}
	ErrorStreamClientExists       = &APIError{HttpCode: 400, ErrCode: 4004, Errmsg: "stream client已存在,app_key重复创建"}
	ErrorVerifyDingtalkClient_    = &APIError{HttpCode: 400, ErrCode: 4005, Errmsg: "验证钉钉客户端失败"}
	ErrorDeleteCredential         = &APIError{HttpCode: 400, ErrCode: 4031, Errmsg: "拒绝执行，此连接器删除失败"}
	ErrorNotFoundCredential       = &APIError{HttpCode: 404, ErrCode: 4041, Errmsg: "连接器不存在"}
	ErrorParams                   = &APIError{HttpCode: 400, ErrCode: 4005, Errmsg: "请求参数错误"}
	ErrorMethod                   = &APIError{HttpCode: 400, ErrCode: 4006, Errmsg: "method 未找到"}
	ErrorOther                    = &APIError{ErrCode: 500, Errmsg: "未知错误"} // 服务自身错误
	// ErrorParams      = &APIError{ErrCode: 403, Errmsg: "没有调用该接口的权限"}

	ErrorInitSubscription = &APIError{HttpCode: 500, ErrCode: 5001, Errmsg: "初始化事件订阅错误"} // 初始化事件订阅服务错误
	ErrorGetSubscription  = &APIError{HttpCode: 404, ErrCode: 4042, Errmsg: "查询事件订阅错误"}  // 查询事件订阅错误
)
