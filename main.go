package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	engine := gin.Default()

	// 加载HTML模板

	// engine.LoadHTMLFiles("templates/news.html")
	// engine.LoadHTMLFiles("templates/goods.html")
	// engine.LoadHTMLFiles("templates/goods.html", "templates/news.html")

	// engine.LoadHTMLFiles("templates/*") // 此用法不对
	engine.LoadHTMLGlob("templates/*")

	engine.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "你好，%v\n", "Golang!")
	})

	engine.GET("/news1", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "新闻页 111\n")
	})

	engine.GET("/news2", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "news.html", gin.H{
			"title": "我是后台数据",
		})
	})

	engine.GET("/goods", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "goods.html", gin.H{
			"name":  "PS5",
			"price": 200,
		})
	})

	engine.GET("/json1", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"success: ": true,
			"message: ": "你好啊！ json1",
		})
	})

	engine.GET("/json2", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"success: ": true,
			"message: ": "你好啊！ json2",
		})
	})

	// engine.GET("/json3", func(ctx *gin.Context) {

	// 	a := &Article{
	// 		Title:   "标题",
	// 		Desc:    "描述",
	// 		Content: "内容",
	// 	}
	// 	ctx.JSON(http.StatusOK, a)
	// })

	engine.GET("/json3", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, &Article{
			Title:   "标题3",
			Desc:    "描述",
			Content: "内容",
		})
	})

	engine.GET("/jsonp", func(ctx *gin.Context) {
		ctx.JSONP(http.StatusOK, &Article{
			Title:   "标题 jsonp",
			Desc:    "描述",
			Content: "内容",
		})
	})

	engine.GET("/xml", func(ctx *gin.Context) {
		ctx.JSONP(http.StatusOK, gin.H{
			"success: ": true,
			"message: ": "你好啊！ xml",
		})
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
