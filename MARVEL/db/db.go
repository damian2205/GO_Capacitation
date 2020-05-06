package db

import (
	"database/sql"
	"fmt"
)

func ObtenerBaseDeDatos() (db *sql.DB, e error) {
	usuario := "root"
	pass := "12345"
	host := "tcp(127.0.0.1:3306)"
	nombreBaseDeDatos := "prueba"
	// Debe tener la forma usuario:contraseña@host/nombreBaseDeDatos
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		return nil, err
	}
	return db, nil
}
