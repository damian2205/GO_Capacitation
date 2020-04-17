package main

import "fmt"

func main()  {
	matriz := [][]int{{1,2},{1,2}}

	if matriz == nil{
		fmt.Printf("No hay nada")
	}else{
		fmt.Print(matriz)
	}
}