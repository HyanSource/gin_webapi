package main

//https://ithelp.ithome.com.tw/articles/10207504
//https://shofan.blogspot.com/2020/02/go-goland-gin-gonicapi-server.html
//中文文檔
//https://github.com/skyhee/gin-doc-cn

import (
	"github.com/HyanSource/gin_webapi/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	setupServer().Run(":8080")
}

//設置初始化以及路由 寫成方法是為了可以做test
func setupServer() *gin.Engine {

	r := gin.Default() //有日誌 以及恢復中間件
	// r := gin.New() //不帶日誌和中間件(request)

	r.GET("/ping", handler.PingGet)

	{
		r.GET("/h/:name", handler.HGet)
		r.POST("/h", handler.HPost)
		r.DELETE("/h", handler.HDelete)
	}

	r.POST("/Login", handler.LoginPost)

	return r
}
