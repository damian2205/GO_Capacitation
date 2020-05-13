package app

import (
	ping "github.com/damian2205/GO_Capacitation/MARVEL/controller"

	"github.com/gin-gonic/gin"
)

func URL(app *Aplicacion) *gin.Engine {
	router := gin.Default()
	mainrouter := router.Group("/")
	mainrouter.GET("/cosas", ping.Home)
	mainrouter.GET("/ping", ping.Ping)
	mainrouter.GET("/datos", app.envioC.ObtenerDatos)
	router.POST("/datos", app.envioC.InsertarDatos)
	router.DELETE("/datos", app.envioC.EliminarDatos)
	router.PUT("/datos", app.envioC.ActualizarDatos)
	// router.POST("/", ping.Post_message)
	return router
}