package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(engine *gin.Engine) {

	adminRG := engine.Group("/admin")
	{
		adminRG.GET("/", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "admin/index.html", gin.H{
				"title": "后台主页",
			})
		})

		adminRG.GET("/news1", func(ctx *gin.Context) {
			a := &Article1{
				Title:   "新闻标题",
				Desc:    "新闻描述",
				Content: "新闻内容",
			}
			ctx.HTML(http.StatusOK, "admin/news1.html", gin.H{
				"news": a,
			})
		})
	}

}
