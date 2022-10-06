> Gin web框架学习笔记
> 本文章主要用于记录自己学习gin框架过程，根据gin官网例子，亲自动手写Demo，好记性不如烂笔头。多写没坏处。


# 什么是gin？
借助官网的一句话： Gin是一个使用Go语言开发的Web框架。 它提供类似Martini的API，但性能更佳，速度提升高达40倍。 如果你是性能和高效的追求者, 你会爱上 Gin。
# 官网
gin的官网地址是：https://gin-gonic.com/zh-cn/
有中文地址，对于国内像我这样英语不太好的人来说还是很友好的。
# 新建Go项目
我使用Goland工作作为go的开发工具。
![](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210041712157.png)

# 引入gin

配置代理，gin的安装包在github上，如果没设置goproxy可能无法下载，在goland中配置goproxy请看如下两个图
![](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210041715895.png)

![](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210041716000.png)

如果您安装go的时候没有设置GOPROXY可以在环境变量中设置。

```shell
export GOPROXY=https://goproxy.cn
```

进入项目目录执行gin的引入

```shell
➜  go-gin-notes git:(main) ✗ go get -u github.com/gin-gonic/gin
go: added github.com/gin-contrib/sse v0.1.0
go: added github.com/gin-gonic/gin v1.8.1
go: added github.com/go-playground/locales v0.14.0
go: added github.com/go-playground/universal-translator v0.18.0
go: added github.com/go-playground/validator/v10 v10.11.1
go: added github.com/goccy/go-json v0.9.11
go: added github.com/json-iterator/go v1.1.12
go: added github.com/leodido/go-urn v1.2.1
go: added github.com/mattn/go-isatty v0.0.16
go: added github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd
go: added github.com/modern-go/reflect2 v1.0.2
go: added github.com/pelletier/go-toml/v2 v2.0.5
go: added github.com/ugorji/go/codec v1.2.7
go: added golang.org/x/crypto v0.0.0-20220926161630-eccd6366d1be
go: added golang.org/x/net v0.0.0-20221002022538-bcab6841153b
go: added golang.org/x/sys v0.0.0-20220928140112-f11e5e49a4ec
go: added golang.org/x/text v0.3.7
go: added google.golang.org/protobuf v1.28.1
go: added gopkg.in/yaml.v2 v2.4.0
```

完毕后可以看到项目中modules文件go.mod已经发生了变化。

![image-20221004172526972](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210041725057.png)

依赖引入进来后，需要考虑如何启动应用，加下来定义程序入口main.go文件。

# 路由配置

项目根目录定义一个main.go的程序入口文件。

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//定义路由引擎
	r := gin.Default()
	//定义默认路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "该项目是gin框架的学习笔记",
		})
	})
	// 启动路由，默认端口是8080
	r.Run()
}
```

启动后，在浏览器上访问http://localhost:8080

![image-20221004173135299](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210041731383.png)

# 日志配置

默认gin是将日志记录打印在控制台，但是项目部署后，我们应该将日志记录在日志文件中（后期也可以使用其他插件将日志发送到ELK中，方便分析。）

还是在入口方法main.go中配置。

## 禁用控制台日志颜色

```go
// 禁用控制台日志颜色
gin.DisableConsoleColor()
```

禁用前

![image-20221004174139832](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210041741991.png)

禁用后

![image-20221004174202579](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210041742647.png)

前后对比可以看到控制台日志的背景颜色没了。

## 打印日志到文件

将日志文件打印到文件中，需要两步，第一步需要创建一个文件，第二个设置gin打印日志到文件。

```go
// 配置将日志打印到文件内
file, _ := os.Create("go-gin-notes.log")
gin.DefaultWriter = io.MultiWriter(file)
```

重新启动项目，可以看到项目下增加了一个go-gin-notes.log日志文件。

![image-20221004175026195](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210041750283.png)

## 错误日志文件配置

在上面配置的日志文件中不会打印错误日志。简单修改下默认路由，使用panic触发错误，请求接口使得其打印错误日志。

![image-20221004175408810](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210041754918.png)

可以看到上图中，控制台打印了错误，但是go-gin-notes.log中并没有错误日志（截图略），gin也有单独的对错误日志文件的配置。

```go
// 错误日志输出到文件配置
errFile, _ := os.Create("go-gin-notes-error.log")
gin.DefaultErrorWriter = io.MultiWriter(errFile)
```

再请求默认接口，然后再查看go-gin-notes-error.log日志。

![image-20221004175701622](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210041757712.png)

错误日志已经打印到了文件中。

**注意：**如果想将日志都放到一个文件中，只需要将DefaultWriter和DefaultErrorWriter指定到一个文件上即可。

## 同时输入日志到文件和控制台

使用上面的DefaultWriter和DefaultErrorWriter后，日志不会再输出到控制台。如果想两者都输出，可以使用如下代码：

```go
// 如果想同时将日志输出到控制台（console）和文件，需要指定io.MultiWriter的第二个参数为os.Stdout
gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
gin.DefaultErrorWriter = io.MultiWriter(errFile, os.Stdout)
```

## 自定义路由DEBUG日志格式

默认的路由日志格式类似如下：

```text
[GIN-debug] GET    /                         --> main.main.func1 (3 handlers)
```

可以使用gin.DebugPrintRouteFunc方法来定义

```

```

修改完毕后后日志格式变为如下：

```text
2022/10/04 18:39:05 itlab1024.com GET / main.main.func2 3
```

## 自定义日志格式

路由请求日志也可以自定义，自定该日志需要我们自己定义引擎，而不是使用gin.Default()。

首先看看默认的格式

```text
[GIN] 2022/10/04 - 18:49:30 | 200 |     122.018µs |             ::1 | GET      "/"
```

根据如下代码修改。

```go
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	// 路由引擎是可以自定义的
	r := gin.New()
	r.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		// 你的自定义格式
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC1123),
			params.Method,
			params.Path,
			params.Request.Proto,
			params.StatusCode,
			params.Latency,
			params.Request.UserAgent(),
			params.ErrorMessage,
		)
	}))
	r.Use(gin.Recovery())
	//定义默认路由
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "该项目是gin框架的学习笔记",
		})
	})
	// 启动路由，默认端口是8080
	r.Run()
	// 也可以修改为其他的端口，比如8000
	//r.Run(":8000")
}
```

修改后

```text
::1 - [Tue, 04 Oct 2022 18:44:44 CST] "GET / HTTP/1.1 200 97.171µs "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/16.0 Safari/605.1.15" "
```

# 响应结构

## JSON

需要注意的是，JSON 使用 unicode 替换特殊 HTML 字符，例如 < 变为 \ u003c。

```go
// JSON
r.GET("/JSON", func(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "code": http.StatusOK,
    "msg":  "<p>该项目是gin框架的学习笔记</p>",
  })
})
```

运行结果如下：

```shell
➜  go-gin-notes git:(main) ✗ curl http://localhost:8080/JSON
{"code":200,"msg":"\u003cp\u003e该项目是gin框架的学习笔记\u003c/p\u003e"}
```

可以看到被转义了。

## AsciiJSON

使用 `AsciiJSON` 方法可以生成只包含 ASCII 字符的 JSON 格式数据，对于非 ASCII 字符会进行转义

```go
r.GET("/AsciiJSON", func(c *gin.Context) {
		c.AsciiJSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "该项目是gin框架的学习笔记",
		})
	})
```

查看结果：

```shell
➜  go-gin-notes git:(main) ✗ curl http://localhost:8080/AsciiJSON           
{"code":200,"msg":"\u8be5\u9879\u76ee\u662fgin\u6846\u67b6\u7684\u5b66\u4e60\u7b14\u8bb0"}
```

可以看到非ASCII字符被转义了。

## PureJSON

对比JSON，PureJSON不会进行转义，而是完全按照字面量进行序列化。

```go
r.GET("/PureJSON", func(c *gin.Context) {
		c.PureJSON(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "<p>该项目是gin框架的学习笔记</p>",
		})
	})
```

运行结果：

```shell
➜  go-gin-notes git:(main) ✗ curl http://localhost:8080/PureJSON
{"code":200,"msg":"<p>该项目是gin框架的学习笔记</p>"}
```

可以看到p标签被原样输出。

## SecureJSON

使用 SecureJSON 防止 json 劫持。如果给定的结构是数组值，则默认预置 `"while(1),"` 到响应体。

非数组类型

```go
➜  go-gin-notes git:(main) ✗ curl http://localhost:8080/SecureJSON           
{"code":200,"msg":"\u003cp\u003e该项目是gin框架的学习笔记\u003c/p\u003e"}
```

结果和JSON方法一样。

数组类型

```go
// SecureJSON数组类型
r.GET("/SecureJSONOfArrayBody", func(c *gin.Context) {
  c.SecureJSON(http.StatusOK, []string{"tom", "jerry", "james"})
})
```

运行结果

```go
➜  go-gin-notes git:(main) ✗ curl http://localhost:8080/SecureJSONOfArrayBody
while(1);["tom","jerry","james"]
```

## JSONP

Jsonp(JSON with Padding) 是 json 的一种"使用模式"，可以让网页从别的域名（网站）那获取资料，即跨域读取数据。

为什么我们从不同的域（网站）访问数据需要一个特殊的技术( JSONP )呢？这是因为同源策略。

同源策略，它是由 Netscape 提出的一个著名的安全策略，现在所有支持 JavaScript 的浏览器都会使用这个策略。

```go
// JSONP
r.GET("/JSONP", func(c *gin.Context) {
  c.JSONP(http.StatusOK, gin.H{
    "code": http.StatusOK,
    "msg":  "<p>该项目是gin框架的学习笔记</p>",
  })
})
```

输出结果

```go
➜  go-gin-notes git:(main) ✗ curl http://localhost:8080/JSONP\?callback\=callbackFunc
callbackFunc({"code":200,"msg":"\u003cp\u003e该项目是gin框架的学习笔记\u003c/p\u003e"});
```

## XML

gin提供了xml方法来返回xml类型的数据

```go
r.GET("/xml", func(c *gin.Context) {
		c.XML(http.StatusOK, gin.H{
			"code": http.StatusOK,
			"msg":  "<p>该项目是gin框架的学习笔记</p>",
		})
	})
```

返回结果

```shell
➜  go-gin-notes git:(main) ✗ curl http://localhost:8080/xml                                        
<map><code>200</code><msg>&lt;p&gt;该项目是gin框架的学习笔记&lt;/p&gt;</msg></map>
```

## YAML

```go
// YAML
r.GET("/yaml", func(c *gin.Context) {
  c.YAML(http.StatusOK, gin.H{
    "code": http.StatusOK,
    "msg":  "<p>该项目是gin框架的学习笔记</p>",
  })
})
```

运行结果

```shell
➜  go-gin-notes git:(main) ✗ curl http://localhost:8080/yaml
code: 200
msg: <p>该项目是gin框架的学习笔记</p>
```

# HTML渲染

gin的html渲染跟原生的很类似，大部分就是遵循原生的用法。我们只需要定义html文件，在其中写模板语言要求的语法的代码即可。

加载html文件有两种方式， LoadHTMLGlob() 或者 LoadHTMLFiles()，前者可以使用表达式，比如templates/*就会匹配templates下的文件。templates/**/*能够匹配templates下的文件夹下的文件，比如templates/files/index.html。

```go
//HTML
// 加载html文件有两种方式， LoadHTMLGlob() 或者 LoadHTMLFiles()
r.LoadHTMLGlob("templates/*")
r.GET("/index", func(c *gin.Context) {
  c.HTML(http.StatusOK, "index.html", gin.H{
    "code": http.StatusOK,
    "msg":  "<p>该项目是gin框架的学习笔记</p>",
  })
})
```

templates下的index.html文件

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>HTML模板</title>
</head>
<body>
{{.}}
</body>
</html>
```

请求http://localhost:8080/index，结果如下：

![image-20221005193659772](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210051937170.png)

结果就是将我的map数据展示了出来。

## 模板文件中自动提示属性名

```go
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>HTML模板</title>
</head>
<body>
{{- /*gotype: go-gin-notes.Result*/ -}}
{{.Code}}
{{.Msg}}
</body>
</html>
```

gotype: 后的go-gin-notes.Result是结构体的名称。

## 路径不同名称相同的模板问题

实际开发中，可能会出现templates/category/index.html，templates/article/index.html这样的文件。也就是路径是不同的，但是最终的文件名是相同的，这就需要我们在文件的开头使用{{define "文件名"}}来声明，以{{end}}结尾。

首先看下不使用define定义的文件。

```go
//category的主页
r.LoadHTMLGlob("templates/**/*")
r.GET("/category/index", func(c *gin.Context) {
  c.HTML(http.StatusOK, "category/index.html", Result{Code: http.StatusOK, Msg: "<p>该项目是gin框架的学习笔记</p>"})
})

//category的主页
r.GET("/article/index", func(c *gin.Context) {
  c.HTML(http.StatusOK, "article/index.html", Result{Code: http.StatusOK, Msg: "<p>该项目是gin框架的学习笔记</p>"})
})
```

article/index.html

![image-20221006095209253](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210060952462.png)

category/index.html

![image-20221006095242335](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210060952430.png)

分别访问下category的主页和article的主页。

访问http://localhost:8080/category/index和http://localhost:8080/article/index，控制台会报错

![image-20221006095654167](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210060956374.png)

对于这种情况就需要使用{{define "文件名"}}来指定，修改两个html文件。

```html
{{define "article/index.html"}}
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
article的主页
</body>
</html>
{{end}}
```

category.html

```html
{{define "category/index.html"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
category的主页
</body>
</html>
{{end}}
```

## 自定义函数

可以自定义一个函数，在模板文件中使用，比如常用的日期格式化。

首先定义一个时间格式化的函数

```go
func DateFormat(time time.Time) string {
	return time.Format("2006 01 02")
}
```

然后在引擎中注册该方法

```go
// 注册时间格式化方法,特别注意要放到加载模板之前
r.SetFuncMap(template.FuncMap{
  "DateFormat": DateFormat,
})
```

然后在模板里就能使用该方法

```go
{{define "timeFmt/index.html"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
{{.Msg | DateFormat}}
</body>
</html>
{{end}}
```

运行结果如下：

![image-20221006110400123](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210061104308.png)

## 模板嵌套

模板嵌套使用{{template "模板名字"}}来实现，准备两个模板，header.html和index.html。需要将header嵌套到index页面中。

看下两个文件的定义

header.html

```html
{{define "nest/header.html"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
这是头部
</body>
</html>
{{end}}
```

index.html

```html
{{define "nest/index.html"}}
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>Title</title>
</head>
<body>
{{template "nest/header.html"}} <br/>
主页
</body>
</html>
{{end}}
```

控制器

```go
//模板嵌套
r.GET("/nest", func(c *gin.Context) {
   c.HTML(http.StatusOK, "nest/index.html", gin.H{"Code": http.StatusOK, "Msg": time.Date(2017, 07, 01, 0, 0, 0, 0, time.UTC)})
})
```

访问后查看界面结果

![image-20221006111753976](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210061117158.png)

# 表单

## 接收参数

首先来看下各种情况下如何接收参数。

## Query

获取Get请求中路径上的参数使用c.Query(key)方法，该方法如果取不到值（没传递），则会返回空字符串。他等价于c.Request.URL.Query().Get(key)

```go
r.GET("/login", func(c *gin.Context) {
		username := c.Query("username")
		pwd := c.Query("pwd")
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"pwd":      pwd,
		})
	})
```

请求结果

```shell
➜  go-gin-notes git:(main) ✗ http http://localhost:8080/login\?username\=itlabUsername\&pwd\=itlabPwd
HTTP/1.1 200 OK
Content-Length: 45
Content-Type: application/json; charset=utf-8
Date: Thu, 06 Oct 2022 05:20:57 GMT

{
    "pwd": "itlabPwd",
    "username": "itlabUsername"
}
```

## PostForm

PostForm能够获取multipart和urlencode的参数

```go
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
```

请求结果，为了能够清晰看到请求信息，看下图

![image-20221006132732209](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210061327366.png)

![image-20221006132822263](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210061328375.png)

可以看到能够正常获取，并且Query也能获取到Post请求url后的r参数，这跟其他语言大同小异。

# 参数绑定

## 绑定基本参数

上面的参数获取都是一个一个获取，当参数较多的情况下还是挺麻烦的，可以使用结构体来接收参数。

定义结构体

```go
type LoginForm struct {
   Username string
   Pwd string
}
```

控制器

```go
r.POST("/login2", func(c *gin.Context) {
		loginForm := LoginForm{}
		err := c.ShouldBind(&loginForm)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, loginForm)
	})
```

![image-20221006141856188](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210061420463.png)

从上图可以得出结论，ShouldBind()方法是无法丙丁路径参数的，如果需要绑定，则使用ShouldBindQuery。

另外上面的例子传递参数都是使用首字母大写，这可能并不规范，我们可以通过在结构体中设置。

```go
type LoginForm struct {
	Username string `form:"username"`
	Pwd      string `form:"pwd"`
	R        string `form:"r"`
}
```

form:"username"代表表单参数是username。

```go
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
```

这里我也绑定了url参数。

重新请求结果

![image-20221006142739819](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210061427016.png)

## 绑定JSON

使用shouldBindJSON方法。

```go
r.POST("/login3", func(c *gin.Context) {
		loginForm := LoginForm{}
		err := c.ShouldBindJSON(&loginForm)
		if err != nil {
			return
		}
		c.JSON(http.StatusOK, loginForm)
	})
```

![image-20221006143004731](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210061430914.png)

上面是绑定了一个结构体，如果绑定多个结构体，需要使用ShouldBindBodyWith，而不能使用ShouldBindJSON。

定义两个结构体。

```go
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
```

将请求参数绑定到LoginForm和LoginFormB上。

```go
r.POST("/login4", func(c *gin.Context) {
  loginForm := LoginForm{}
  loginFormB := LoginFormB{}
  err := c.ShouldBindJSON(&loginForm)
  if err != nil {
    log.Printf("%v", err)
    return
  }
  err = c.ShouldBindJSON(&loginFormB)
  if err != nil {
    log.Printf("%v", err)
    return
  }
  c.JSON(http.StatusOK, loginForm)
})
```

![image-20221006162132223](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210061621361.png)

应该使用ShouldBindBodyWith方法，重新修改。

```go
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
```

![image-20221006162337093](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210061623260.png)

这就没有问题了。

## 参数验证

可以在结构体中定义是否必传验证

```go
type LoginForm struct {
	Username string `form:"username" binding:"required"`
	Pwd      string `form:"pwd"`
	R        string `form:"r"`
}
```

请求会打印如下错误

![image-20221006143543080](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210061435222.png)

# 文件上传

## 单个文件

```go
// 上传文件，单个文件
r.POST("/upload", func(context *gin.Context) {
  file, _ := context.FormFile("file")
  context.SaveUploadedFile(file, "upload/a.png")
})
```

![image-20221006163015527](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210061630742.png)

## 多个文件

多个文件，传递一个数组，循环上传即可。

```go
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
```

结果如下：

![image-20221006163519932](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210061635100.png)

# 静态文件服务

上面上传的文件无法在服务器上访问，比如：http://localhost:8080/upload/a.png。

需要设置静态文件服务器方能够访问。

```go
r.Static("/upload", "upload")
```

除了Static方法还有StaticFS，StaticFile，功能类似。

重新访问图片

![image-20221006163819496](https://itlab1024-1256529903.cos.ap-beijing.myqcloud.com/202210061638708.png)

未完，待续。。。。。。
