package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/sayhello/", sayHello)

	r.Run()
}

func sayHello(c *gin.Context) {
	data := c.Query("name")
	pData := c.Params
	fmt.Println(pData)

	res := gin.H{
		"message": "hello " + data,
	}

	c.JSON(200, res)

}
