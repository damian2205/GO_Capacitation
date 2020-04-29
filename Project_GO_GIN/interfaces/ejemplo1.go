package main

import (
	"fmt"
)

type Persona struct {
	Nombre   string
	Apellido string
	Edad     uint8
}
type Programador struct {
	Persona
	Especialidad       string
	LenguajesFavoritos []string
}

func (p Persona) Saludar() {
	fmt.Println("Hola, mi nombre es "+p.Nombre+" "+p.Apellido+" y tengo %d \n", p.Edad)
}

func (p *Persona) Cumpleanios() {
	p.Edad += 1
}

func (p Programador) beberCafe() {
	fmt.Println("¡Me siento vivo!")
}

// func main() {
// 	p := Programador{
// 		Persona{"Miguel", "Solar", 23},
// 		"Juega lolsito",
// 		[]string{"GO", "Python", "JavaScript"},
// 	}
// 	p.Saludar()
// }

// func main() {
// 	p1 := Programador{
// 		Persona{"Orlando", "Olarte", 26},
// 		"Desarrollo Web",
// 		[]string{"Go", "Python", "JavaScript"},
// 	}
// 	p2 := Persona{"Daniel", "Herrera", 32}
// 	p3 := Persona{"Carmen", "Salazar", 20}
// 	p1.Saludar()
// 	p2.Saludar()
// 	p3.Saludar()
// }

// func main() {
// 	p1 := Programador{
// 		Persona{"Orlando", "Monteverde", 26},
// 		"Desarrollo Web",
// 		[]string{"Go", "Python", "JavaScript"},
// 	}
// 	p2 := Persona{"Daniel", "Herrera", 32}
// 	p3 := Persona{"Carmen", "Salazar", 20}

// 	vecinos := []Vecino{p1, p2, p3}
// 	for _, v := range vecinos {
// 		v.Saludar()
// 	}
// }

func main() {
	p1 := Programador{
		Persona{"Orlando", "Monteverde", 26},
		"Desarrollo Web",
		[]string{"Go", "Python", "JavaScript"},
	}
	p2 := Persona{"Daniel", "Herrera", 32}
	p3 := Persona{"Carmen", "Salazar", 20}

	var vecino Vecino
	vecino = p1
	fmt.Println(vecino)
	vecino = p2
	fmt.Println(vecino)
	vecino = p3
	fmt.Println(vecino)
}
