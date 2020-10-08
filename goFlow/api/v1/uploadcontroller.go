package v1

import (
	"github.com/gin-gonic/gin"
	"goFlow/servers"
	"goFlow/utils/errmsg"
	"net/http"
)

func UploadData(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")

	fileSize := fileHeader.Size
	url, code := servers.UploadFile(file, fileSize)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    errmsg.GetErrMsg(code),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   url,
		"msg":    errmsg.GetErrMsg(code),
	})

}
