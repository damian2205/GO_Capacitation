package controller

import "github.com/gin-gonic/gin"

type EnvioController interface {
	ObtenerUser(c *gin.Context)
	InsertarUser(c *gin.Context)
	EliminarUser(c *gin.Context)
	ActualizarUser(c *gin.Context)
}
