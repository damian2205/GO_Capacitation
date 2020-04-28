package main

import (
	"fmt"
	"strconv"
)

func main() {
	for_if()
	sichi()
}

func for_if() {
	var num1, num2 int
	fmt.Println("Digite un numero para comenzar la cuenta")
	fmt.Scan(&num1)
	fmt.Println("Digite un numero para maximo de cuenta")
	fmt.Scan(&num2)
	fmt.Println("\n")

	for i := num1; i <= num2; i++ {
		if i%2 != 1 {
			a := strconv.Itoa(i)
			fmt.Println("Par  " + a + "\n")
		} else if i%2 != 0 {
			b := strconv.Itoa(i)
			fmt.Println("Impar  " + b + "\n")
		}
	}
}

func sichi() {
	var edad int
	fmt.Println("Ingresa tu edad para verificar que es mayor de edad")
	fmt.Scan(&edad)

	switch {
	case edad > 18:
		fmt.Println("Eres mayor de edad")
		fallthrough //Permite que la ejecucion de los casos sea simultanea,sin esta solo hayun break

	case edad > 22:
		fmt.Println("Edad perfecta")

	default:
		fmt.Println("No tienes edad adecuada")
	}
}
