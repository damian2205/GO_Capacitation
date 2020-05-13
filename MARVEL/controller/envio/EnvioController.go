package controller

import (
	"fmt"
	"log"
	"net/http"

	// "log"
	// "net/http"

	"github.com/damian2205/GO_Capacitation/MARVEL/controller"
	"github.com/damian2205/GO_Capacitation/MARVEL/dto"

	// "github.com/damian2205/GO_Capacitation/MARVEL/dto"
	"github.com/damian2205/GO_Capacitation/MARVEL/services"
	"github.com/gin-gonic/gin"
)

type envioController struct {
	envioS services.EnvioServices
}

func NewEnvioController(s services.EnvioServices) controller.EnvioController {
	return &envioController{s}
}

func (s *envioController) ObtenerDatos(c *gin.Context) {
	datos, err := s.envioS.ObtenerUser()
	if err != nil {
		fmt.Printf("Error obteniendo contactos: %v", err)
		return
	}
	c.JSON(200, gin.H{"datos": datos})
}

func (s *envioController) InsertarDatos(c *gin.Context) {
	var User dto.Contacto
	err := c.ShouldBindJSON(&User)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	log.Println("Usuario ingresado", User)
	// OB.InsertarUser(User)
	response := s.envioS.InsertarUser(User)
	c.JSON(http.StatusOK, gin.H{"Respuesta de ingreso": response})
}

func (s *envioController) EliminarDatos(c *gin.Context) {
	var IDuser dto.Contacto
	err := c.ShouldBindJSON(&IDuser)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	log.Println("Usuario iliminado", IDuser.IDusuarios)
	// OB.InsertarUser(ID)
	response := s.envioS.EliminarUser(IDuser)
	c.JSON(http.StatusOK, gin.H{"Respuesta de eliminacion": response})
}

func (s *envioController) ActualizarDatos(c *gin.Context) {
	var Update dto.Contacto
	err := c.ShouldBindJSON(&Update)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	log.Println("Usuario actualizado", Update.Nombre, Update.Usuario, Update.Contraseña)
	// OB.InsertarUser(Userr)
	response := s.envioS.ActualizarUser(Update)
	c.JSON(http.StatusOK, gin.H{"Respuesta de update": response})

}
