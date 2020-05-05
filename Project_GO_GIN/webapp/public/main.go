package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	// "../interfaces"
)

type Contacto struct {
	nombre, usuario, contraseña string
	idusurios                   int
}

func ObtenerDB() (db *sql.DB, e error) {
	usuario := "root"
	pass := "12345"
	host := "tcp(127.0.0.1:3306)"
	nombreDB := "prueba"
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreDB))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Home_pagina(c *gin.Context) {
	c.JSON(200, gin.H{
		"title": "HEY HOLA ES EL INDEX",
	})
}
func ObtenerUser(w http.ResponseWriter, r *http.Request) {
	db, err := ObtenerDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	filas, err := db.Query("SELECT idusuarios, nombre, usuario, contraseña FROM usuarios")

	if err != nil {
		return nil, err
	}

	defer filas.Close()

	json.NewEncoder(w).Encode(filas)

	// // var c Contacto

	// for filas.Next() {
	// 	err = filas.Scan(&c.idusurios, &c.nombre, &c.usuario, &c.contraseña)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	contactos = append(contactos, c)
	// }
	// return contactos, nil
}
func CrearUser(c *gin.Context) {

}
func main() {
	// conexion DB
	db, err := ObtenerDB()
	if err != nil {
		fmt.Printf("Error obteniendo base de datos: %v", err)
		return
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		fmt.Printf("Error conectando: %v", err)
		return
	}
	fmt.Printf("Conectado correctamente\n")

	// SERVER
	r := gin.Default()
	r.GET("/", Home_pagina)
	r.GET("/users", ObtenerUser)
	r.POST("/create", CrearUser)
	r.Run("127.0.0.1:8000")

}
