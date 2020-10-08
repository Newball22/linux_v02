package errmsg

//将这个错误信息处理抽离出来package

const (
	SUCCESS       = 200
	ERROR         = 500
	NETWORK_IS_OK = 4000
	/*code>1000:用户相关的错误*/
	ERROR_USERNAME_USED    = 1001
	ERROR_PASSWORD_WRONG   = 1002
	ERROR_USER_NOT_EXIT    = 1003
	ERROR_TOKEN_NOT_EXIT   = 1004
	ERROR_TOKEN_RUNTIME    = 1005
	ERROR_TOKEN_WRONG      = 1006
	ERROR_TOKEN_TYPE_WRONG = 1007
	ERROR_USER_NOT_LIMIT   = 1008
	ERROR_USERNAME_WRONG   = 1009

	/*code>2000:文章相关的错误*/
	ERROR_ART_NOT_EXIT = 2001

	/*code>3000:分类相关的错误*/
	ERROR_CATENAME_USED = 3001
)

var (
	codeMsg = map[int]string{
		SUCCESS:                "OK",
		ERROR:                  "FAIL",
		NETWORK_IS_OK:          "接口正常",
		ERROR_USERNAME_USED:    "用户名已注册",
		ERROR_PASSWORD_WRONG:   "密码错误",
		ERROR_USER_NOT_EXIT:    "用户不存在",
		ERROR_TOKEN_NOT_EXIT:   "token不存在",
		ERROR_TOKEN_RUNTIME:    "token超时",
		ERROR_TOKEN_WRONG:      "token验证错误",
		ERROR_TOKEN_TYPE_WRONG: "token格式错误",
		ERROR_ART_NOT_EXIT:     "文章不存在",
		ERROR_CATENAME_USED:    "分类已经存在",
		ERROR_USER_NOT_LIMIT:   "用户没有权限",
		ERROR_USERNAME_WRONG:   "用户名错误",
	}
)

func GetErrMsg(code int) string {
	return codeMsg[code]
}
