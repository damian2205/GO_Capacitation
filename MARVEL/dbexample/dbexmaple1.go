/*package dbexample

import (
	"database/sql"
	"fmt"
)

func TraerDB() {
	db, err := sql.Open("mysql", "root:12345(localhost:3306)/prueba")

	if err != nil {
		fmt.Println("error1")
	}

	dato, err2 := db.Query("select * from usuarios")

	if err2 != nil {
		fmt.Println("error2")
	}

	for dato.Next() {
		var nombre, usuario string
		err3 := dato.Scan(&nombre, &usuario)
		if err3 != nil {
			fmt.Println("error3")
		}
		fmt.Println(nombre + " " + usuario)
	}
}
*/