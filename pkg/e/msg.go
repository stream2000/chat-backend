package e

var MsgFlags = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",
	//ErrInvalidParams:         "请求参数错误",
	//ErrAuthCheckTokenFail:    "Token鉴权失败",
	//ErrAuthCheckTokenTimeout: "Token已超时",
	//ErrAuthToken:             "Token生成失败",
	//ErrAuth:                  "Token错误",
	//ErrGetRecordFailed:       "读取打卡记录错误",
	//ErrUserAlreadyExists:     "用户已经存在，请考虑重试",
	//ErrCreateUser:            "创建用户时出错",
	//ErrGetAllGroup:           "获取小组信息时出错",
}

// GetMsg get error information based on Code
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
