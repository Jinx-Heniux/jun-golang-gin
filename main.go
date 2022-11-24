package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Article1 struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Article2 struct {
	Title   string `json:"title" xml:"title"`
	Content string `json:"content" xml:"content"`
}

type UserInfo struct {
	Username string `json:"user" form:"username"`
	Password string `json:"pw" form:"password"`
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

	engine.GET("/news1", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "新闻页 111\n")
	})

	engine.GET("/news2", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "default/news2.html", gin.H{
			"title": "我是后台数据",
		})
	})

	engine.GET("/news3", func(ctx *gin.Context) {
		a := &Article1{
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
		ctx.JSON(http.StatusOK, &Article1{
			Title:   "标题3",
			Desc:    "描述",
			Content: "内容",
		})
	})

	engine.GET("/jsonp", func(ctx *gin.Context) {
		ctx.JSONP(http.StatusOK, &Article1{
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

	// post
	engine.GET("/user1", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "default/user1.html", gin.H{})
	})

	// 获取Post请求的传值 表单数据
	engine.POST("/doAddUser1", func(ctx *gin.Context) {
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
	engine.GET("/getUser", func(ctx *gin.Context) {
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
	engine.GET("/user2", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "default/user2.html", gin.H{})
	})
	engine.POST("/doAddUser2", func(ctx *gin.Context) {
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
	engine.GET("/xml1", func(ctx *gin.Context) {
		ctx.JSONP(http.StatusOK, gin.H{
			"success: ": true,
			"message: ": "你好啊！ xml",
		})
	})

	// 获取Post Xml数据
	engine.POST("/xml2", func(ctx *gin.Context) {
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

	engine.POST("/create", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "创建请求\n")
	})

	engine.PUT("/update", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "更新请求\n")
	})

	engine.DELETE("/delete", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "删除请求\n")
	})

	////////// 动态路由 //////////
	// /list1/123 /list1/456
	engine.GET("/list1/:cid", func(ctx *gin.Context) {
		cid := ctx.Param("cid")
		ctx.String(200, "%v", cid)
	})

	////////// 后台 //////////

	engine.GET("/admin", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "admin/index.html", gin.H{
			"title": "后台主页",
		})
	})

	engine.GET("/admin/news1", func(ctx *gin.Context) {
		a := &Article1{
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
