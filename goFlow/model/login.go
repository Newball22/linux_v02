package model

import (
	"goFlow/middleware"
	"goFlow/utils/errmsg"
)

var (
	code  int
	token string
)

func CheckLogin(data *User) (string, int) {
	code = CheckData(data.Username, data.Password)

	if code != errmsg.SUCCESS {
		return "", code
	}
	token, code = middleware.SetToken(data.Username)
	return token, code

}

//登录验证
func CheckData(username, password string) int {
	var (
		user User
	)
	DB.Where("username=?", username).First(&user)

	if user.ID == 0 {
		return errmsg.ERROR_USER_NOT_EXIT
	}
	if ScryptPassword(password) != user.Password {
		return errmsg.ERROR_PASSWORD_WRONG
	}
	if user.Role != 1 {
		return errmsg.ERROR_USER_NOT_LIMIT
	}
	return errmsg.SUCCESS
}
