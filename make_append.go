package main 

import "fmt"

func main()  {

	slice := make([]int,3,5)	//Make permite agregar un parametro despues de declarar el tipo de dato es la cantidad, segundo parametro es la capacidad
	slice = append(slice, 5)	//Append permite ingresar mas valores al slice ya declarado

	fmt.Println(slice)
	
}