package main

import "fmt"

func main_suma() {
	var num1, num2 int
	fmt.Println("Ingrese un numero")
	fmt.Scan(&num1)
	fmt.Println("Ingrese un numero")
	fmt.Scan(&num2)

	resultado := num1 + num2

	fmt.Println(resultado)
}
