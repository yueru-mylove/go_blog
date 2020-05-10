package e

var MsgFlags  = map[int]string {
	SUCCESS : "ok",
	ERROR : "fail",
	INVALID_PARAMS : "请求参数错误",
	ERROR_EXIST_TAG : "标签已经存在",
	ERROR_NOT_EXISTS_TAG : "标签不存在",
	ERROR_NOT_EXISTS_ARTICLE : "文章不存在",
	ERROR_AUTH_CHECK_TOKEN_FAIL : "token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "token已经失效",
	ERROR_AUTH_TOKEN : "token生成失败",
	ERROR_AUTH : "token错误",
}

func GetMsg(code int) string  {

	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}