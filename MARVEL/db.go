package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func obtenerBaseDeDatos() (db *sql.DB, e error) {
	usuario := "root"
	pass := "12345"
	host := "tcp(127.0.0.1:3306)"
	nombreBaseDeDatos := "prueba"
	// Debe tener la forma usuario:contraseña@host/nombreBaseDeDatos
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		return nil, err
	}
	return db, nil
}

type Contacto struct {
	nombre, usuario, contraseña string
	ID                          int
}

func main() {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		fmt.Printf("Error obteniendo base de datos: %v", err)
		return
	}
	// Terminar conexión al terminar función
	defer db.Close()

	// Ahora vemos si tenemos conexión
	err = db.Ping()
	if err != nil {
		fmt.Printf("Error conectando: %v", err)
		return
	}
	// Listo, aquí ya podemos usar a db!
	fmt.Printf("Conectado correctamente")

	c := Contacto{
		nombre:     "benito",
		usuario:    "toca melas",
		contraseña: "12345",
	}

	_err := insertar(c)
	if _err != nil {
		fmt.Println("Error insertando: %v", _err)
	} else {
		fmt.Println("Insertado con exito")
	}

	// contactos, err := obtenerContactos()
	// if err != nil {
	// 	fmt.Printf("Error obteniendo contactos: %v", err)
	// 	return
	// }
	// for _, contacto := range contactos {
	// 	fmt.Printf("%v\n", contacto)
	// }
}

func insertar(c Contacto) (e error) {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("Insert to into prueba (nombre, usuario, contraseña) VALUES(?,?,?)")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()

	_, err = sentenciaPreparada.Exec(c.nombre, c.usuario, c.contraseña)
	if err != nil {
		return err
	}
	return nil
}

// func obtenerContactos() ([]Contacto, error) {
// 	contactos := []Contacto{}
// 	db, err := obtenerBaseDeDatos()
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer db.Close()
// 	filas, err := db.Query("SELECT id, nombre, usuario, contraseña, contraseña_ FROM prueba")

// 	if err != nil {
// 		return nil, err
// 	}
// 	// Si llegamos aquí, significa que no ocurrió ningún error
// 	defer filas.Close()

// 	// Aquí vamos a "mapear" lo que traiga la consulta en el while de más abajo
// 	var c Contacto

// 	// Recorrer todas las filas, en un "while"
// 	for filas.Next() {
// 		err = filas.Scan(&c.ID, &c.nombre, &c.usuario, &c.contraseña, &c.contraseña_)
// 		// Al escanear puede haber un error
// 		if err != nil {
// 			return nil, err
// 		}
// 		// Y si no, entonces agregamos lo leído al arreglo
// 		contactos = append(contactos, c)
// 	}
// 	// Vacío o no, regresamos el arreglo de contactos
// 	return contactos, nil
// }
