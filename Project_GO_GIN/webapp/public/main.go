package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Home_pagina(c *gin.Context) {
	c.JSON(200, gin.H{
		"title": "HEY HOLA ES EL INDEX",
	})
}



func CrearUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HOLA CANRAL",
	})
}
func main() {

	// SERVER
	router := gin.Default()
	router.GET("/", Home_pagina)
	router.GET("/users", ObtenerUser)
	router.GET("/create", CrearUser)
	router.Run("127.0.0.1:8000")

}
