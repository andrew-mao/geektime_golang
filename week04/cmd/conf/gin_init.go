package conf

import (
	"github.com/gin-gonic/gin"
	"mblog/controller"
)

type Gin struct {
	Listen string `toml:"listen"`
	JWT    bool   `toml:"jwt"`
	Engine *gin.Engine
}

//gin初始化
func (g *Gin) GinInit() {
	gin.SetMode(gin.ReleaseMode)
	g.Engine = gin.New()
	g.register()
}

//gin启动
func (g *Gin) GinRun() {
	g.Engine.Run(g.Listen)
}

//gin 注册路由
//遵循RESTful规范
func (g *Gin) register() {
	//文章相关路由，增删改查
	controller.RegisterArticle(g.Engine)
	////评论相关路由，增删改查
	//controller.RegisterComment(g.Engine)
	////阅读记录相关路由，增查
	//controller.RegisterCollect(g.Engine)
	//标签相关路由，增删改查
	//管理相关路由，
}

//全局中间件加载
func (g *Gin) middleware() {
	if g.JWT {
		g.Engine.Use()
	}
}
