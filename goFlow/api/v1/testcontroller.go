package v1

import (
	"github.com/gin-gonic/gin"
	"goFlow/utils/errmsg"
	"net/http"
)

func TestNet(c *gin.Context) {
	code = errmsg.NETWORK_IS_OK
	c.JSON(http.StatusOK, gin.H{
		"staus": code,
		"msg":   errmsg.GetErrMsg(code),
	})
}
