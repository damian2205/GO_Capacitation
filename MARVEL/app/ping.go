package app

import (
	"fmt"
	"io/ioutil"

	OB "github.com/damian2205/GO_Capacitation/MARVEL/repo"
	"github.com/gin-gonic/gin"
)

func ObtenerDatos(c *gin.Context) {
	datos, err := OB.ObtenerContactos()
	if err != nil {
		fmt.Printf("Error obteniendo contactos: %v", err)
		return
	}
	// contacto, err := json.Marshal(datos)
	// if err != nil {
	// 	fmt.Printf("Error codificando datos, %v", err)
	// }
	c.JSON(200, gin.H{"datos": datos})
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"mensaje": "PONG",
	})
}

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"title": "HEY HOLA ES EL INDEX",
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
