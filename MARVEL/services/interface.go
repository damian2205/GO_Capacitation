package services

import (
	"github.com/damian2205/GO_Capacitation/MARVEL/dto"
)

type EnvioServices interface {
	ObtenerUser() (datos []dto.Contacto, e error)
	InsertarUser(d dto.Contacto) (e error)
	EliminarUser(d dto.Contacto) error
	ActualizarUser(d dto.Contacto) error
}
