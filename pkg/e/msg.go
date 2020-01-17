package e

var MsgFlags = map[int]string{
	SUCCESS:                 "ok",
	ERROR:                   "fail",
	ErrInvalidParams:        "请求参数错误",
	ErrUnKnownInternalError: "未知服务器内部错误，请稍后重试",

	ErrInvalidBasicAuthParam: "Basic认证参数错误",
	ErrBasicAuthFailed:       "未通过Basic认证",

	ErrInvalidBearerAuthParams: "Bearer认证参数错误",
	ErrAuthCheckTokenFail:      "Token鉴权失败",
	ErrAuthCheckTokenTimeout:   "Token已超时",
	ErrJwtAuth:                 "Token生成失败",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
