package main

import "fmt"

func main() {
	nombres := []string{
		"Pluma",
		"Gay",
		"Pluma",
		"Gay",
	}

	for index, nombre := range nombres {
		fmt.Println(nombre, index)
	}

	for _, nombre := range nombres {
		fmt.Println(nombre)
	}
	for index, _ := range nombres {
		fmt.Println(index)
	}
	for index, _ := range nombres {
		fmt.Println(index)
	}

	cadena := "PUTO EL QUE LO LEA"

	for index, letra := range cadena {
		fmt.Printf("La letra %q esta en el index %d\n", letra, index)
	}
}
