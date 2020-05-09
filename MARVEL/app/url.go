package app

import (
	ping "github.com/damian2205/GO_Capacitation/MARVEL/controller"
	"github.com/gin-gonic/gin"
)

func URL() *gin.Engine {
	router := gin.Default()
	router.GET("/", ping.Home)
	router.GET("/ping", ping.Ping)
	router.GET("/datos", ping.ObtenerDatos)
	router.POST("/datos", ping.InsertarDatos)
	router.DELETE("/datos", ping.EliminarDatos)
	router.PUT("/datos", ping.ActualizarDatos)
	// router.POST("/", ping.Post_message)
	return router
}
