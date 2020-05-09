package controller

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/damian2205/GO_Capacitation/MARVEL/dto"
	OB "github.com/damian2205/GO_Capacitation/MARVEL/services"
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
	var User dto.Contacto
	err := c.ShouldBindJSON(&User)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	log.Println("Usuario ingresado", User)
	// OB.InsertarUser(User)
	response := OB.InsertarUser(User)
	c.JSON(http.StatusOK, gin.H{"Respuesta de ingreso": response})
}

func EliminarDatos(c *gin.Context) {
	var IDuser dto.Contacto
	err := c.ShouldBindJSON(&IDuser)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	log.Println("Usuario iliminado", IDuser.IDusuarios)
	// OB.InsertarUser(ID)
	response := OB.EliminarUser(IDuser)
	c.JSON(http.StatusOK, gin.H{"Respuesta de eliminacion": response})
}

func ActualizarDatos(c *gin.Context) {
	var Update dto.Contacto
	err := c.ShouldBindJSON(&Update)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	// log.Println("Usuario actualizado", Update)
	// OB.InsertarUser(Userr)
	response := OB.ActualizarUser(Update)
	c.JSON(http.StatusOK, gin.H{"Respuesta de update": response})

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
