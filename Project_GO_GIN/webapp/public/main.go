package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	// "../interfaces"
)

type Contacto struct {
	idusurios  string `json:"id, omitempty"`
	nombre     string `json:"nombre, omitempty"`
	usuario    string `json:"usuario, omitempty"`
	contraseña string `json:"contraseña, omitempty"`
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

func ObtenerUser(c *gin.Context) ([]Contacto, error) {
	contactos := []Contacto{}
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

	var a Contacto

	for filas.Next() {
		err = filas.Scan(&a.idusurios, &a.nombre, &a.usuario, &a.contraseña)
		if err != nil {
			return nil, err
		}
		contactos = append(contactos, a)
	}
	// return contactos, nil
	c.JSON(200, gin.H{
		"message": string(contactos),
	})
}

func CrearUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "HOLA CANRAL",
	})
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
	router := gin.Default()
	router.GET("/", Home_pagina)
	router.GET("/users", ObtenerUser)
	router.GET("/create", CrearUser)
	router.Run("127.0.0.1:8000")

}
