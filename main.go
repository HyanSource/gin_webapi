package main

//https://ithelp.ithome.com.tw/articles/10207504
//https://shofan.blogspot.com/2020/02/go-goland-gin-gonicapi-server.html
//中文文檔
//https://github.com/skyhee/gin-doc-cn

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //有日誌 以及恢復中間件
	// r := gin.New() //不帶日誌和中間件(request)
	r.GET("/ping", PingGet)

	{
		r.GET("/h", HGet)
		r.POST("/h", HPost)
	}

	r.POST("/Login", LoginPost)

	r.Run(":8080")
}

func PingGet(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func HGet(c *gin.Context) {

	c.JSON(200, gin.H{
		"n": c.Param("n"),
	})
}

func HPost(c *gin.Context) {

	//獲取參數
	id := c.PostForm("id")
	//獲取參數 可設置默認值
	name := c.DefaultPostForm("name", "Hyan")
	money := c.DefaultPostForm("money", "0")

	c.JSON(200, gin.H{
		"id":    id,
		"name":  name,
		"money": money,
	})
}

//加上binding字段 如果值是空的 則會不成功並返回錯誤
type LoginData struct {
	Username string `from:"username" json:"username" binding:"required"`
	Password string `from:"password" json:"passowrd" binding:"required"`
}

func LoginPost(c *gin.Context) {
	var json LoginData
	if err := c.BindJSON(json); err != nil {
		if json.Username == "Hyan" && json.Password == "123" {
			c.JSON(200, gin.H{
				"status": "OK",
			})
		} else {
			c.JSON(200, gin.H{
				"status": "not OK",
			})
		}
	}
}
