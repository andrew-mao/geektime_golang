package controller

import "github.com/gin-gonic/gin"

func RegisterComment(e *gin.Engine) {
	a := e.Group("/comment")
	a.POST("comment")
	a.GET("comment")
	a.PUT("comment")
	a.DELETE("comment")

}
