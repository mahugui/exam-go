package httpbase

var MsgFlags = map[int]string {
	SUCCESS:                "操作成功",
	ERROR:                  "操作失败",
	INVALID_PARAMS:         "请求参数错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok{
		return msg
	}
	return MsgFlags[ERROR]
}