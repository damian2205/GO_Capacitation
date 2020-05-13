package controller

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"title": "HEY HOLA ES EL INDEX",
	})
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"mensaje": "PONG",
	})
}

func Post_message(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, gin.H{
		"message": string(value),
	})
}
