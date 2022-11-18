package main

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func UnixToTime(timestamp int) string {
	t := time.Unix(int64(timestamp), 0)
	return t.Format("2006-01-02 15:04:05")
}

func main() {
	engine := gin.Default()

	// 注意：要把自定义模板函数放在加载模板前
	engine.SetFuncMap(template.FuncMap{
		"UnixToTime": UnixToTime,
	})

	////////// 加载HTML模板 //////////

	// engine.LoadHTMLFiles("templates/news.html")
	// engine.LoadHTMLFiles("templates/goods.html")
	// engine.LoadHTMLFiles("templates/goods.html", "templates/news.html")

	// engine.LoadHTMLFiles("templates/*") // 此用法不对
	// engine.LoadHTMLGlob("templates/*")
	engine.LoadHTMLGlob("templates/**/*") // 加载templates子文件夹的内容

	////////// 配置静态web服务 //////////
	engine.Static("/static", "./static")

	////////// 处理请求 //////////

	// engine.GET("/", func(ctx *gin.Context) {
	// 	ctx.String(http.StatusOK, "我是首页！你好，%v\n", "Golang!")
	// })
	engine.GET("/", func(ctx *gin.Context) {

		ctx.HTML(http.StatusOK, "default/index.html", gin.H{
			"title": "我是前台数据",
			"msg":   "我是msg",
			"score": 50,
			"hobby": []string{"吃饭", "睡觉", "写代码"},
			"newsList": []interface{}{
				&Article{
					Title:   "标题111",
					Desc:    "描述111",
					Content: "内容111",
				},
				&Article{
					Title:   "标题222",
					Desc:    "描述222",
					Content: "内容222",
				},
			},
			"emptySlice": []string{},
			"news": &Article{
				Title:   "标题333",
				Desc:    "描述333",
				Content: "内容333",
			},
			"date": 1629423555,
		})

	})

	engine.GET("/news1", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "新闻页 111\n")
	})

	engine.GET("/news2", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "default/news2.html", gin.H{
			"title": "我是后台数据",
		})
	})

	engine.GET("/news3", func(ctx *gin.Context) {
		a := &Article{
			Title:   "新闻标题",
			Desc:    "新闻描述",
			Content: "新闻内容",
		}
		ctx.HTML(http.StatusOK, "default/news3.html", gin.H{
			"news": a,
		})
	})

	engine.GET("/goods1", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "default/goods1.html", gin.H{
			"name":  "PS5",
			"price": 200,
		})
	})

	////////// json //////////

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

	// 获取Get请求的传值
	engine.GET("/json4", func(ctx *gin.Context) {

		username := ctx.Query("username")
		age := ctx.Query("age")
		page := ctx.DefaultQuery("page", "1")
		ctx.JSON(http.StatusOK, gin.H{
			"username": username,
			"age":      age,
			"page":     page,
		})
	})

	// 获取Get请求的传值
	engine.GET("/article", func(ctx *gin.Context) {
		id := ctx.DefaultQuery("id", "1")
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "新闻详情",
			"id":  id,
		})
	})

	engine.GET("/xml", func(ctx *gin.Context) {
		ctx.JSONP(http.StatusOK, gin.H{
			"success: ": true,
			"message: ": "你好啊！ xml",
		})
	})

	////////// restful api //////////

	engine.POST("/create", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "创建请求\n")
	})

	engine.PUT("/update", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "更新请求\n")
	})

	engine.DELETE("/delete", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "删除请求\n")
	})

	////////// 后台 //////////

	engine.GET("/admin", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "后台主页",
		})
	})

	engine.GET("/admin/news1", func(ctx *gin.Context) {
		a := &Article{
			Title:   "新闻标题",
			Desc:    "新闻描述",
			Content: "新闻内容",
		}
		ctx.HTML(http.StatusOK, "admin/news1.html", gin.H{
			"news": a,
		})
	})

	engine.Run(":8008")
}
