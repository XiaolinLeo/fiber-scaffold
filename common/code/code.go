package code
const (
	VERYTY_Token_NULL = 10001
	VERYTY_Token_Fail = 401
	SIGN_ERROR        = 20000
	ERROR             = 500
	SUCCESS             = 200
)

var MsgFlags = map[int]string{
	VERYTY_Token_NULL: "Token为空！",
	VERYTY_Token_Fail: "Token认证失败！",
	SIGN_ERROR:        "生成签名失败",
	ERROR: "Failed",
	SUCCESS: "Success",
}

func MsgInfo(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
