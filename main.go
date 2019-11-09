package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Age      int    `form:"age" json:"age"`
}

func selfMiddleWare() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("before middleware")
		context.Next()
		fmt.Println("After middleWare")
	}
}
func AuthMiddleWare() gin.HandlerFunc {
    return func(context *gin.Context) {
		if cookie,err := context.Request.Cookie("session-id");err == nil {
			if cookie.Value == "123"{
				context.Next()
				return
			}
		}
		context.JSON(http.StatusUnauthorized,gin.H {
		    "error":"unauthed",
		})
	}
}
// gin框架demo
func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, name)
	})
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + "is" + action
		c.String(http.StatusOK, message)
	})
	r.GET("/welcome", func(c *gin.Context) {
		firstName := c.DefaultQuery("firstname", "gust")
		lastName := c.Query("lastname")
		c.String(http.StatusOK, "hello %s %s", firstName, lastName)
	})
	r.POST("/form-upload", func(context *gin.Context) {
		message := context.PostForm("message")
		nick := context.DefaultPostForm("nick", "anonymous")
		context.JSON(http.StatusOK, gin.H{
			"status": gin.H{
				"status_code": http.StatusOK,
				"status":      "ok",
			},
			"message": message,
			"nick":    nick,
		})
	})
	//同时使用查询字串和body参数发送数据给服务器
	r.PUT("/post", func(context *gin.Context) {
		id := context.Query("id")
		name := context.Query("name")
		page := context.PostForm("page")
		message := context.PostForm("message")
		context.JSON(http.StatusOK, gin.H{
			"status_code": http.StatusOK,
			"id":          id,
			"page":        page,
			"name":        name,
			"message":     message,
		})
	})
	//上传文件
	r.POST("file-upload", func(context *gin.Context) {
		name := context.PostForm("name")
		log.Println(name)
		file, header, err := context.Request.FormFile("upload")
		if err != nil {
			context.String(http.StatusBadRequest, "Bad Request")
			return
		}
		fileName := header.Filename;
		log.Println(fileName, file, err)
		out, err := os.Create(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
		context.String(http.StatusOK, "upload Sucessful!")
	})
	//上传多个文件
	r.POST("/multiFile-Upload", func(context *gin.Context) {
		//设置最大空间
		err := context.Request.ParseMultipartForm(200000)
		if err != nil {
			log.Fatal(err)
		}
		files := context.Request.MultipartForm.File["upload"]
		for i, _ := range files {
			file, err := files[i].Open()
			defer file.Close()
			if err != nil {
				log.Fatal(err)
			}
			out, err := os.Create(files[i].Filename)
			defer out.Close()
			if err != nil {
				log.Fatal(err)
			}

			_, err = io.Copy(out, file)
			if err != nil {
				log.Fatal(err)
			}
			context.String(http.StatusCreated, "upload successful")
		}
	})
	//参数绑定
	r.POST("/login", func(context *gin.Context) {
		var user User
		err := context.Bind(&user)
		if err != nil {
			log.Fatal(err)
		}
		context.String(http.StatusOK, user.Username, user.Password, user.Age)
	})
	//重定向
	r.GET("/redict/google", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https://google.com")
	})

	/*
	分组路由
	 */
	 v1 := r.Group("v1");
	v1.GET("/login", func(context *gin.Context) {
		context.String(http.StatusOK,"v1 login")
	})
	v2 := r.Group("v2");
	v2.GET("/login", func(context *gin.Context) {
		context.String(http.StatusOK,"v2 login")
	})
	//
	r.Use(selfMiddleWare())
	{
		r.GET("/testMidWare", func(context *gin.Context) {
			context.String(http.StatusOK, "")
		})
	}

	r.GET("/autoauth", func(context *gin.Context) {
	   cookie := &http.Cookie {
	       Name : "session-id",
	       Value :"123",
	       Path:"/",
	       HttpOnly : true,
	   }
	   http.SetCookie(context.Writer, cookie)
	   context.String(http.StatusOK,"login sucessful!")
	})

	r.GET("/testAuth",AuthMiddleWare(), func(context *gin.Context) {
		context.String(http.StatusOK, "")
	})

	r.GET("/sync", func(context *gin.Context) {
		time.Sleep(5 * time.Second)
		fmt.Println("Done! in path" + context.Request.URL.String())
	})

	//异步请求
	r.GET("/async", func(context *gin.Context) {
		copyedC := context.Copy();
		go func() {
			time.Sleep(5 * time.Second)
			fmt.Println("Done! in path" + copyedC.Request.URL.String())
		}()
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
