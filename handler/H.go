package handler

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HGet(c *gin.Context) {

	n := c.Param("name")

	c.JSON(200, gin.H{
		"name": n,
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

func HDelete(c *gin.Context) {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		fmt.Println("Delete err:", err)
	}
	c.JSON(200, gin.H{
		"id": id,
	})
}
