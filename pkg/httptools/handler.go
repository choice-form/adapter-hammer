package httptools

// 第三方接口响应接口
// 用于统一第三方返回错误
type APIErrorHandler interface {
	// ErrorHandler 对可能的错误进行统一封装
	//  - 要具有对access_token错误进行统一封装，如果出现access_token 过期的行为需要返回 error: ErrorAccessToken,以便在出现时程序会获取新的 access_token 并缓存起来，并重试
	ErrorHandler(error) error
	SetResponse(resp *Response)
}
