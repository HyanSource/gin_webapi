package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

//加上binding字段 如果值是空的 則會不成功並返回錯誤
type LoginData struct {
	Username string `from:"username" json:"username" binding:"required"`
	Password string `from:"password" json:"password" binding:"required"`
}

func LoginPost(c *gin.Context) {
	var json LoginData

	if err := c.BindJSON(&json); err == nil {
		fmt.Println(json)
		if json.Username == "Hyan" && json.Password == "123" {
			c.JSON(200, gin.H{
				"username": json.Username,
				"password": json.Password,
			})
		} else {
			c.JSON(200, gin.H{
				"username": json.Username,
				"password": json.Password,
			})
		}
	}

}
