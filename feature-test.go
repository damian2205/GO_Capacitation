package main

import (
	"fmt"
	"strconv"
)

func main() {
	var num1, num2 int
	fmt.Println("Ingrese un numero")
	fmt.Scan(&num1)
	fmt.Println("Ingrese un numero")
	fmt.Scan(&num2)

	resultado := num1 + num2
	resultado1 := num1 - num2
	resultado3 := num1 / num2
	resultado2 := num1 * num2

	conver := strconv.Itoa(resultado)
	conver1 := strconv.Itoa(resultado1)
	conver3 := strconv.Itoa(resultado3)
	conver2 := strconv.Itoa(resultado2)

	fmt.Println("La suma es: " + conver)
	fmt.Println("La resta es: " + conver1)
	fmt.Println("La division es: " + conver3)
	fmt.Println("La multiplicacion es: " + conver2)

}
