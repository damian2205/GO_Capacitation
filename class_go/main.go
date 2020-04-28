package main

import (
	"fmt"
)

func main() {
	type user struct {
		año              int
		nombre, apellido string
	}

	usuario := new(user)
	var año_ int
	var nombre_ string
	var apellido_ string
	fmt.Println("Ingresa tu año de nacimiento")
	fmt.Scan(&año_)
	fmt.Println("Ingresa tu nombre")
	fmt.Scan(&nombre_)
	fmt.Println("Ingresa tu apellido")
	fmt.Scan(&apellido_)
	usuario.año = año_
	usuario.nombre = nombre_
	usuario.apellido = apellido_

	fmt.Println(usuario)
}
