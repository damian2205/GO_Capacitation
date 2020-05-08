package app

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	OB "github.com/damian2205/GO_Capacitation/MARVEL/repo"
	"github.com/gin-gonic/gin"
)

type Contacto struct {
	IDusurios  int    `json:"idusarios"`
	Nombre     string `json:"nombre"`
	Usuario    string `json:"usuario"`
	Contraseña string `json:"contraseña"`
}

func ObtenerDatos(c *gin.Context) {
	datos, err := OB.ObtenerUser()
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

func InsertarDatos(c *gin.Context) {
	var d Contacto
	err := c.BindJSON(&d)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	log.Println("User dto", d)
	response := OB.InsertarUser(d)
	// c.JSON(200, gin.H{"datos": "DB",
	// 	"message":   "Creado correctamente",
	// 	"resultado": todo.IDusurios,
	// })
	c.JSON(http.StatusOK, gin.H{"responseContent": response})
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
