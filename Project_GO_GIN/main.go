package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	// "../interfaces"
)

func Home_pagina(c *gin.Context) {
	c.JSON(200, gin.H{
		"title": "HEY HOLA ES EL INDEX",
	})
}
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Pong",
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
func main() {
	r := gin.Default()
	r.GET("/", Home_pagina)
	r.GET("/ping", Ping)
	r.POST("/", Post_message)
	r.Run("127.0.0.1:8000") // listen and server :8080
}
