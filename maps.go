package main

import "fmt"

func main() {
	dias := make(map[string]int, 7)

	dias["Lunes"] = 1
	dias["Martes"] = 2
	dias["Miercoles"] = 3
	dias["Jueves"] = 4
	dias["Viernes"] = 5
	dias["Sabado"] = 6
	dias["Domingo"] = 7

	fmt.Println(dias, "\n")

	semana := map[string]int{
		"Lunes":     1,
		"Martes":    2,
		"Miercoles": 3,
		"Jueves":    4,
		"Viernes":   5,
		"Sabado":    6,
		"Domingo":   7,
	}

	// fmt.Println(semana["Martes"])
	delete(semana, "Domingo")
	fmt.Println(semana)

}
