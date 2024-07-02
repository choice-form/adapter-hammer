package errno

import "errors"

// event-subscription
var ErrorNotFoundStreamClient = errors.New("事件订阅stream client未找到,请检查是否设置了webhook")
var ErrorStreamClientExist = errors.New("事件订阅stream client已存在,app_key重复创建")
