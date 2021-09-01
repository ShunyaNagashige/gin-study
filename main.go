package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/post", postfunc)
	r.Run(":10000")
}

type InputCompany struct {
	UserName string
}

func postfunc(c *gin.Context) {
	var hoge InputCompany
	c.BindJSON(&hoge)
	fmt.Println(hoge.UserName)

	c.JSON(200, gin.H{
		"message": hoge.UserName,
	})
}
