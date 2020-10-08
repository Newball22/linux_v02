package v1

import (
	"github.com/gin-gonic/gin"
	"goFlow/model"
	"goFlow/utils/errmsg"
	"goFlow/utils/validator"
	"net/http"
	"strconv"
)

var code int

//注册用户
func AddUser(c *gin.Context) {
	var user model.User
	var msg string
	_ = c.ShouldBindJSON(&user)
	msg, code = validator.Validate(&user)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    msg,
		})
		return
	}

	code = model.CheckUser(user.Username)
	if code == errmsg.SUCCESS {
		model.AddAUser(&user)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})

}

//查询单个用户
func GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	user, code := model.FindAUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   user,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//查询用户分页列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1 //gorm就会取消分页
	}
	if pageNum == 0 {
		pageNum = -1
	}
	userList, code, num := model.FindAllUsers(pageSize, pageNum)
	if code == errmsg.ERROR {
		code = errmsg.ERROR
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   userList,
		"msg":    errmsg.GetErrMsg(code),
		"total":  num,
	})

}

//编辑用户
func EditUser(c *gin.Context) {
	var user model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&user)
	code := model.CheckUser(user.Username)
	if code == errmsg.SUCCESS {
		model.EditAUser(id, &user)

	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})

}

//删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteAUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}
