package controller

import "github.com/gin-gonic/gin"

type EnvioController interface {
	ObtenerDatos(c *gin.Context)
	InsertarDatos(c *gin.Context)
	EliminarDatos(c *gin.Context)
	ActualizarDatos(c *gin.Context)
}
