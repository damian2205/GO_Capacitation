package main

import (
	"fmt"
)

func main() {
	type user struct {
		edad, año        int
		nombre, apellido string
	}

	usuario := new(user)
	var edad_ int
	var nombre_ string
	var apellido_ string
	var año_ int
	fmt.Println("Ingresa tu edad")
	fmt.Scan(&edad_)
	fmt.Println("Ingresa tu año de nacimiento")
	fmt.Scan(&edad_)
	fmt.Println("Ingresa tu nombre")
	fmt.Scan(&nombre_)
	fmt.Println("Ingresa tu apellido")
	fmt.Scan(&apellido_)
	usuario.edad = edad_
	usuario.año = año_
	usuario.nombre = nombre_
	usuario.apellido = apellido_

	fmt.Println(usuario)
}
