package v1

import (
	"github.com/gin-gonic/gin"
	"goFlow/model"
	"goFlow/utils/errmsg"
	"net/http"
	"strconv"
)

//添加文章
func AddArticle(c *gin.Context) {
	var art model.Article
	_ = c.ShouldBindJSON(&art)
	code = model.AddArt(&art)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})

}

//查询一篇文章
func GetArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	art, code := model.FindArt(id)
	if code == errmsg.ERROR {
		code = errmsg.ERROR_ART_NOT_EXIT
		c.JSON(http.StatusOK, gin.H{
			"status": code,
			"msg":    errmsg.GetErrMsg(code),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   art,
		"msg":    errmsg.GetErrMsg(code),
	})
}

//查询分类下的文章列表
func GetCateArticle(c *gin.Context) {
	cid, _ := strconv.Atoi(c.Param("cid"))
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1 //gorm就会取消分页
	}
	if pageNum == 0 {
		pageNum = -1
	}
	artList, code, total := model.FindAllCateArt(cid, pageSize, pageNum)
	if code == errmsg.ERROR {
		code = errmsg.ERROR
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   artList,
		"msg":    errmsg.GetErrMsg(code),
		"total":  total,
	})
}

//查询所有文章分页列表
func GetArticles(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	if pageSize == 0 {
		pageSize = -1 //gorm就会取消分页
	}
	if pageNum == 0 {
		pageNum = -1
	}
	artList, code, total := model.FindAllArt(pageSize, pageNum)
	if code == errmsg.ERROR {
		code = errmsg.ERROR
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"data":   artList,
		"msg":    errmsg.GetErrMsg(code),
		"total":  total,
	})

}

//编辑文章
func EditArticle(c *gin.Context) {
	var art model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&art)
	code = model.EditArt(id, &art)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})

}

//删除文章
func DeleteArticle(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteArt(id)
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    errmsg.GetErrMsg(code),
	})
}
