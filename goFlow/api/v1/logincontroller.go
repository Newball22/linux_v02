package v1

import (
	"github.com/gin-gonic/gin"
	"goFlow/model"
	"goFlow/utils/errmsg"
	"net/http"
)

func Login(c *gin.Context) {
	var data model.User
	var token string
	_ = c.ShouldBindJSON(&data) //注意这里是Json
	token, code = model.CheckLogin(&data)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    errmsg.GetErrMsg(code),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
		"token":  token,
	})
}
