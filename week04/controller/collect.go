package controller

import "github.com/gin-gonic/gin"

func RegisterCollect(e *gin.Engine) {
	a := e.Group("/collect")
	a.GET("info")
	a.POST("")
}
