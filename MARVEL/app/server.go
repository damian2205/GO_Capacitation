package app

import (
	"log"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

//Variable para saber si se encuentra en modo de testeo
var TestingMode bool

func Start() {
	router := URL()
	// router.Run("127.0.0.1:8000")
	StartServer(router, TestingMode)
}

// Inicia el servidor

func StartServer(engine *gin.Engine, TestingMode bool) {
	if TestingMode {
		servertest := httptest.NewServer(engine)
		log.Print(servertest)
		return
	}
	srv := &http.Server{
		Addr:    ":8000",
		Handler: engine,
	}
	srv.ListenAndServe()
}
