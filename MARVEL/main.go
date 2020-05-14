package main

import (
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/damian2205/GO_Capacitation/MARVEL/app"
	"github.com/gin-gonic/gin"
)

//Variable para saber si se encuentra en modo de testeo
var TestingMode bool

func main() {
	appplicacion := app.BuildAplicacion()
	engine := app.URL(appplicacion)
	StartServer(engine, TestingMode)
}

// Inicia el servidor

func StartServer(engine *gin.Engine, TestingMode bool) {
	if TestingMode {
		servertest := httptest.NewServer(engine)
		log.Print(servertest)
		return
	}
	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine,
	}
	srv.ListenAndServe()
}
