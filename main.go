package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "你好，%v\n", "Golang!")
	})

	engine.GET("/news", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "新闻页 111\n")
	})

	engine.POST("/create", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "创建请求\n")
	})

	engine.PUT("/update", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "更新请求\n")
	})

	engine.DELETE("/delete", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "删除请求\n")
	})

	engine.Run(":8008")
}
