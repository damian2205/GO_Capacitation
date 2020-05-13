package app

import (
	controller "github.com/damian2205/GO_Capacitation/MARVEL/controller"
	c "github.com/damian2205/GO_Capacitation/MARVEL/controller/envio"
	b "github.com/damian2205/GO_Capacitation/MARVEL/db"
	s "github.com/damian2205/GO_Capacitation/MARVEL/services/envio"
)

type Aplicacion struct {
	envioC controller.EnvioController
}

//inicializacion de contructores
func BuildAplicacion() *Aplicacion {

	db := b.NewDB()

	envioServices := s.NewEnvioServices(db)
	return &Aplicacion{envioC: c.NewEnvioController(envioServices)}
}
