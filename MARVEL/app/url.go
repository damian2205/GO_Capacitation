package app

import (
	ping "github.com/damian2205/GO_Capacitation/MARVEL/controller"

	"github.com/gin-gonic/gin"
)

func URL(app *Aplicacion) *gin.Engine {
	router := gin.Default()
	// mainrouter := router.Group("/")
	router.GET("/cosas", ping.Home)
	router.GET("/ping", ping.Ping)
	router.GET("/datos", app.envioC.ObtenerUser)
	router.POST("/datos", app.envioC.InsertarUser)
	router.DELETE("/datos", app.envioC.EliminarUser)
	router.PUT("/datos", app.envioC.ActualizarUser)
	// router.POST("/", ping.Post_message)
	return router
}
