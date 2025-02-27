package resp

type Resp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func NewResp(code int, msg string, data any) Resp {
	return Resp{
		code,
		msg,
		data,
	}
}

func Ok(msg string) Resp {
	return NewResp(OK, msg, nil)
}

func Data(data any) Resp {
	return NewResp(OK, "请求成功!", data)
}

func Err(code int) Resp {
	return NewResp(code, ErrorMsg[code], nil)
}

var ErrorMsg = map[int]string{
	1:     "当前设备不在线！",
	10001: "您没有权限执行此操作！",
	10002: "该请求中需要携带验证信息！",
	10003: "该请求缺少必要参数！",
	10004: "未查找到此设备！",
	10005: "操作执行出现错误！",
	10006: "设备未配置MAC地址，无法执行此操作！",
	10007: "未配置设备信息！",
}

const (
	OK            = 0
	DeviceOffline = 1
	TokenNotFound = iota + 9999
	Unauthorized
	InvalidQueryParam
	DeviceNotFound
	ActionExecuteFailed
	MacNotFound
	MissingConfiguration
)
