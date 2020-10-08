package v1

import (
	"github.com/gin-gonic/gin"
	"goFlow/model"
	"goFlow/utils/errmsg"
	"net/http"
	"strconv"
)

//新增分类
func AddCate(c *gin.Context) {
	var cate model.Category
	_ = c.ShouldBindJSON(&cate)
	code = model.CheckCate(cate.Name)
	if code == errmsg.SUCCESS {
		model.AddACate(&cate)
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   cate,
		"msg":    errmsg.GetErrMsg(code),
	})

}

//查询分类
func GetCate(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1 //gorm就会取消分页
	}
	if pageNum == 0 {
		pageNum = -1
	}
	id, _ := strconv.Atoi(c.Param("id"))
	cate, code, total := model.GetCateS(id, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   cate,
		"msg":    errmsg.GetErrMsg(code),
		"total":  total,
	})
}

//编辑分类
func EditCate(c *gin.Context) {
	var cate model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&cate)
	code = model.EditACate(id, &cate)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})

}

//删除分类
func DeleteCate(c *gin.Context) {
	var cate model.Category
	_ = c.ShouldBindJSON(&cate)
	code = model.CheckCate(cate.Name)
	id, _ := strconv.Atoi(c.Param("id"))
	if code == errmsg.SUCCESS {
		code = errmsg.ERROR
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    errmsg.GetErrMsg(code),
		})
		return
	}
	code = model.DeleteACate(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}
