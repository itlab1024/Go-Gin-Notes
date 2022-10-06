package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"html/template"
	"log"
	"net/http"
	"time"
)

type Result struct {
	Code int
	Msg  string
}

func main() {
	// 日志配置
	// 禁用控制台日志颜色
	//gin.DisableConsoleColor()
	// 配置将日志打印到文件内
	//file, _ := os.Create("go-gin-notes.log")
	//gin.DefaultWriter = io.MultiWriter(file)
	// 错误日志输出到文件配置
	//errFile, _ := os.Create("go-gin-notes.log")
	//gin.DefaultErrorWriter = io.MultiWriter(errFile)
	//如果想同时将日志输出到控制台（console）和文件，需要指定io.MultiWriter的第二个参数为os.Stdout
	//gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
	//gin.DefaultErrorWriter = io.MultiWriter(errFile, os.Stdout)
	// 自定义路由日志格式
	//gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
	//	log.Printf("itlab1024.com %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
	//}
	//定义路由引擎
	r := gin.Default()
	// 路由引擎是可以自定义的
	// r := gin.New()
	//r.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
	//	// 你的自定义格式
	//	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
	//		params.ClientIP,
	//		params.TimeStamp.Format(time.RFC1123),
	//		params.Method,
	//		params.Path,
	//		params.Request.Proto,
	//		params.StatusCode,
	//		params.Latency,
	//		params.Request.UserAgent(),
	//		params.ErrorMessage,
	//	)
	//}))
	//r.Use(gin.Recovery())
	//定义默认路由
	r.GET("/", func(c *gin.Context) {
		//panic("测试错误")
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "该项目是gin框架的学习笔记",
		})
	})
	// JSON
	r.GET("/JSON", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "<p>该项目是gin框架的学习笔记</p>",
		})
	})

	// AsciiJSON
	r.GET("/AsciiJSON", func(c *gin.Context) {
		c.AsciiJSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "该项目是gin框架的学习笔记",
		})
	})
	//PureJSON
	r.GET("/PureJSON", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "<p>该项目是gin框架的学习笔记</p>",
		})
	})
	// SecureJSON非数组类型
	r.GET("/SecureJSON", func(c *gin.Context) {
		c.SecureJSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "<p>该项目是gin框架的学习笔记</p>",
		})
	})
	// SecureJSON数组类型
	r.GET("/SecureJSONOfArrayBody", func(c *gin.Context) {
		c.SecureJSON(http.StatusOK, []string{"tom", "jerry", "james"})
	})
	// JSONP
	r.GET("/JSONP", func(c *gin.Context) {
		c.JSONP(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "<p>该项目是gin框架的学习笔记</p>",
		})
	})
	//xml
	r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "<p>该项目是gin框架的学习笔记</p>",
		})
	})
	// YAML
	r.GET("/yaml", func(c *gin.Context) {
		c.YAML(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "<p>该项目是gin框架的学习笔记</p>",
		})
	})

	// protoBuf
	r.GET("/protoBuf", func(c *gin.Context) {
		reps := []int64{int64(1), int64(2)}
		label := "test"
		// protobuf 的具体定义写在 testdata/protoexample 文件中。
		data := &protoexample.Test{
			Label: &label,
			Reps:  reps,
		}
		// 请注意，数据在响应中变为二进制数据
		// 将输出被 protoexample.Test protobuf 序列化了的数据
		c.ProtoBuf(http.StatusOK, data)
	})

	//HTML
	// 加载html文件有两种方式， LoadHTMLGlob() 或者 LoadHTMLFiles()
	//r.LoadHTMLGlob("templates/*")
	//r.GET("/index", func(c *gin.Context) {
	//	c.HTML(http.StatusOK, "index.html", Result{Code: http.StatusOK, Msg: "<p>该项目是gin框架的学习笔记</p>"})
	//})
	// 注册时间格式化方法,特别注意要放到加载模板之前
	r.SetFuncMap(template.FuncMap{
		"DateFormat": DateFormat,
	})
	r.LoadHTMLGlob("templates/**/*")
	//category的主页
	r.GET("/category/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "category/index.html", Result{Code: http.StatusOK, Msg: "<p>该项目是gin框架的学习笔记</p>"})
	})

	//category的主页
	r.GET("/article/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "article/index.html", Result{Code: http.StatusOK, Msg: "<p>该项目是gin框架的学习笔记</p>"})
	})

	r.GET("/timeFmt", func(c *gin.Context) {
		c.HTML(http.StatusOK, "timeFmt/index.html", gin.H{"Code": http.StatusOK, "Msg": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC)})
	})

	//模板嵌套
	r.GET("/nest", func(c *gin.Context) {
		c.HTML(http.StatusOK, "nest/index.html", gin.H{"Code": http.StatusOK, "Msg": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC)})
	})

	// 参数接收
	r.GET("/login", func(c *gin.Context) {
		username := c.Query("username")
		pwd := c.Query("pwd")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"pwd":      pwd,
		})
	})

	r.POST("/login", func(c *gin.Context) {
		r := c.Query("r")
		username := c.PostForm("username")
		pwd := c.PostForm("pwd")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"pwd":      pwd,
			"r":        r,
		})
	})
	r.POST("/login2", func(c *gin.Context) {
		loginForm := LoginForm{}
		err := c.ShouldBind(&loginForm)
		if err != nil {
			return
		}
		// 同时绑定url的参数
		err = c.ShouldBindQuery(&loginForm)
		c.JSON(http.StatusOK, loginForm)
	})

	r.POST("/login3", func(c *gin.Context) {
		loginForm := LoginForm{}
		err := c.ShouldBindJSON(&loginForm)
		if err != nil {
			log.Printf("%v", err)
			return
		}
		c.JSON(http.StatusOK, loginForm)
	})

	r.POST("/login4", func(c *gin.Context) {
		loginForm := LoginForm{}
		loginFormB := LoginFormB{}
		err := c.ShouldBindBodyWith(&loginForm, binding.JSON)
		if err != nil {
			log.Printf("%v", err)
			return
		}
		err = c.ShouldBindBodyWith(&loginFormB, binding.JSON)
		if err != nil {
			log.Printf("%v", err)
			return
		}
		c.JSON(http.StatusOK, loginForm)
	})

	// 上传文件，单个文件
	r.POST("/upload", func(context *gin.Context) {
		file, _ := context.FormFile("file")
		context.SaveUploadedFile(file, "upload/a.png")
	})
	// 多个文件
	r.POST("/upload2", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["file[]"]

		for _, file := range files {
			log.Println(file.Filename)

			// 上传文件至指定目录
			c.SaveUploadedFile(file, "upload/"+file.Filename)
		}
	})

	// 静态文件服务
	r.Static("/upload", "upload")
	// 启动路由，默认端口是8080
	r.Run()
	// 也可以修改为其他的端口，比如8000
	//r.Run(":8000")
}

func DateFormat(time time.Time) string {
	return time.Format("2006 01 02")
}

type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Pwd      string `form:"pwd"`
	R        string `form:"r"`
}
type LoginFormB struct {
	Username string `form:"username" binding:"required"`
	Pwd      string `form:"pwd"`
	R        string `form:"r"`
}
