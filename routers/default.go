package routers

import (
	"encoding/xml"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultRoutersInit(engine *gin.Engine) {

	defaultRG := engine.Group("/")
	{
		// engine.GET("/", func(ctx *gin.Context) {
		// 	ctx.String(http.StatusOK, "我是首页！你好，%v\n", "Golang!")
		// })
		defaultRG.GET("/", func(ctx *gin.Context) {

			ctx.HTML(http.StatusOK, "default/index.html", gin.H{
				"title": "我是前台数据",
				"msg":   "我是msg",
				"score": 50,
				"hobby": []string{"吃饭", "睡觉", "写代码"},
				"newsList": []interface{}{
					&Article1{
						Title:   "标题111",
						Desc:    "描述111",
						Content: "内容111",
					},
					&Article1{
						Title:   "标题222",
						Desc:    "描述222",
						Content: "内容222",
					},
				},
				"emptySlice": []string{},
				"news": &Article1{
					Title:   "标题333",
					Desc:    "描述333",
					Content: "内容333",
				},
				"date": 1629423555,
			})

		})

		defaultRG.GET("/news1", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "新闻页 111\n")
		})

		defaultRG.GET("/news2", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "default/news2.html", gin.H{
				"title": "我是后台数据",
			})
		})

		defaultRG.GET("/news3", func(ctx *gin.Context) {
			a := &Article1{
				Title:   "新闻标题",
				Desc:    "新闻描述",
				Content: "新闻内容",
			}
			ctx.HTML(http.StatusOK, "default/news3.html", gin.H{
				"news": a,
			})
		})

		defaultRG.GET("/goods1", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "default/goods1.html", gin.H{
				"name":  "PS5",
				"price": 200,
			})
		})

		////////// json //////////

		defaultRG.GET("/json1", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, map[string]interface{}{
				"success: ": true,
				"message: ": "你好啊！ json1",
			})
		})

		defaultRG.GET("/json2", func(ctx *gin.Context) {
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

		defaultRG.GET("/json3", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, &Article1{
				Title:   "标题3",
				Desc:    "描述",
				Content: "内容",
			})
		})

		defaultRG.GET("/jsonp", func(ctx *gin.Context) {
			ctx.JSONP(http.StatusOK, &Article1{
				Title:   "标题 jsonp",
				Desc:    "描述",
				Content: "内容",
			})
		})

		// 获取Get请求的传值
		// http://127.0.0.1:8008/json4?username=zhang&age=18
		// 127.0.0.1:8008/json4?username=zhang&age=18&page=100
		defaultRG.GET("/json4", func(ctx *gin.Context) {

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
		defaultRG.GET("/article", func(ctx *gin.Context) {
			id := ctx.DefaultQuery("id", "1")
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "新闻详情",
				"id":  id,
			})
		})

		// post
		defaultRG.GET("/user1", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "default/user1.html", gin.H{})
		})

		// 获取Post请求的传值 表单数据
		defaultRG.POST("/doAddUser1", func(ctx *gin.Context) {
			username := ctx.PostForm("username")
			password := ctx.PostForm("password")
			age := ctx.DefaultPostForm("age", "20")

			ctx.JSON(http.StatusOK, gin.H{
				"username": username,
				"password": password,
				"age":      age,
			})
		})

		// 获取Get Post传递的数据绑定到结构体
		// http://127.0.0.1:8008/getUser?username=zhangsan&password=123
		defaultRG.GET("/getUser", func(ctx *gin.Context) {
			user := &UserInfo{}
			fmt.Printf("user: %p | %p | %v\n", &user, user, user)
			if err := ctx.ShouldBind(user); err == nil {
				fmt.Printf("user: %p | %p | %v\n", &user, user, user)
				ctx.JSON(http.StatusOK, user)
			} else {
				ctx.JSON(http.StatusOK, gin.H{
					"err": err.Error(),
				})
			}
		})

		// 获取Get Post传递的数据绑定到结构体
		defaultRG.GET("/user2", func(ctx *gin.Context) {
			ctx.HTML(http.StatusOK, "default/user2.html", gin.H{})
		})

		defaultRG.POST("/doAddUser2", func(ctx *gin.Context) {
			user := &UserInfo{}
			fmt.Printf("user: %p | %p | %v\n", &user, user, user)
			if err := ctx.ShouldBind(user); err == nil {
				fmt.Printf("user: %p | %p | %v\n", &user, user, user)
				ctx.JSON(http.StatusOK, user)
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"err": err.Error(),
				})
			}
		})

		// xml
		defaultRG.GET("/xml1", func(ctx *gin.Context) {
			ctx.JSONP(http.StatusOK, gin.H{
				"success: ": true,
				"message: ": "你好啊！ xml",
			})
		})

		// 获取Post Xml数据
		defaultRG.POST("/xml2", func(ctx *gin.Context) {
			article := &Article2{}
			fmt.Printf("article: %p | %p | %v\n", &article, article, article)
			xmlRawData, _ := ctx.GetRawData()
			fmt.Printf("raw data: %v\n", xmlRawData)
			if err := xml.Unmarshal(xmlRawData, article); err == nil {
				fmt.Printf("article: %p | %p | %v\n", &article, article, article)
				ctx.JSON(http.StatusOK, article)
			} else {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"err": err.Error(),
				})
			}

		})

		////////// restful api //////////

		defaultRG.POST("/create", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "创建请求\n")
		})

		defaultRG.PUT("/update", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "更新请求\n")
		})

		defaultRG.DELETE("/delete", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "删除请求\n")
		})

		////////// 动态路由 //////////
		// /list1/123 /list1/456
		defaultRG.GET("/list1/:cid", func(ctx *gin.Context) {
			cid := ctx.Param("cid")
			ctx.String(200, "%v", cid)
		})

	}

}
