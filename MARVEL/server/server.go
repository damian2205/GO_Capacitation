package main

import (
	ping "github.com/damian2205/GO_Capacitation/MARVEL/app"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/", ping.Home)
	r.GET("/datos", ping.ObtenerDatos)
	r.GET("/ping", ping.Ping)
	r.POST("/", ping.Post_message)
	r.Run("127.0.0.1:8000") // listen and server :8080
}
