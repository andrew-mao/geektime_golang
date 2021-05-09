package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterArticle(e *gin.Engine) {
	art := e.Group("/artapi")
	art.POST("article", PostArticle)
	art.GET("article", GetArticle)
	art.PUT("article", PutArticle)
	art.DELETE("article", DelArticle)
}

//上传新文章
func PostArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data":   "上传完成",
		"status": 200,
	})
}

//获取文章
//根据文章id
func GetArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data":   "获取完成",
		"status": 200,
	})
}

//更新文章
//全量更新
func PutArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data":   "更新完成",
		"status": 200,
	})
}

//删除文章
//根据文章id，标记删除标记
func DelArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data":   "删除完成",
		"status": 200,
	})
}
