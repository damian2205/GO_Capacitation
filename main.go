package main

import (
	"fmt"
)

func main() {
	type user struct {
		edad             int
		nombre, apellido string
	}

	usuario := new(user)
	var edad_ int
	var nombre_ string
	var apellido_ string
	fmt.Println("Ingresa tu edad")
	fmt.Scan(&edad_)
	fmt.Println("Ingresa tu nombre")
	fmt.Scan(&nombre_)
	fmt.Println("Ingresa tu apellido")
	fmt.Scan(&apellido_)
	usuario.edad = edad_
	usuario.nombre = nombre_
	usuario.apellido = apellido_

	fmt.Println(usuario)
}
