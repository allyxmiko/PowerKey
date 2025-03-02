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
	10001: "该请求中需要携带验证信息！",
	10002: "您没有权限执行此操作！",
	10003: "该请求缺少必要参数！",
	10004: "未查找到此设备！",
	10005: "操作执行出现错误！",
	10006: "设备未配置MAC地址，无法执行此操作！",
	10007: "未配置设备信息！",
	10008: "未查找到此用户！",
	10009: "密码错误！",
	10010: "更新密码失败！",
	10011: "更新Token失败！",
	10012: "添加设备失败！",
	10013: "路径参数错误！",
	10014: "删除设备失败！",
	10015: "更新设备信息失败！",
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
	UserNotFound
	PasswordNotMatch
	UpdatePasswordFailed
	UpdateTokenFailed
	DeviceAddFailed
	InvalidPathParam
	DeviceDeleteFailed
	DeviceUpdateFailed
)
