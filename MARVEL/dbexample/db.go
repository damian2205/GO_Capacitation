package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Contacto struct {
	nombre, usuario, contraseña string
	idusuarios                  int
}

func obtenerBaseDeDatos() (db *sql.DB, e error) {
	usuario := "root"
	pass := "12345"
	host := "tcp(127.0.0.1:3306)"
	nombreBaseDeDatos := "prueba"
	// Debe tener la forma usuario:contraseña@protocolo(host:puerto)/nombreBaseDeDatos
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", usuario, pass, host, nombreBaseDeDatos))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	menu := `¿Qué deseas hacer?
[1] -- Insertar
[2] -- Mostrar
[3] -- Actualizar
[4] -- Eliminar
[5] -- Salir
----->	`
	var eleccion int
	var c Contacto
	for eleccion != 5 {
		fmt.Print(menu)
		fmt.Scanln(&eleccion)
		scanner := bufio.NewScanner(os.Stdin)
		switch eleccion {
		case 1:
			fmt.Println("Ingresa el nombre:")
			if scanner.Scan() {
				c.nombre = scanner.Text()
			}
			fmt.Println("Ingresa el usuario:")
			if scanner.Scan() {
				c.usuario = scanner.Text()
			}
			fmt.Println("Ingresa la contraseña:")
			if scanner.Scan() {
				c.contraseña = scanner.Text()
			}
			err := insertar(c)
			if err != nil {
				fmt.Printf("Error insertando: %v", err)
			} else {
				fmt.Println("Insertado correctamente")
			}
		case 2:
			contactos, err := obtenerContactos()
			if err != nil {
				fmt.Printf("Error obteniendo contactos: %v", err)
			} else {
				for _, contacto := range contactos {
					fmt.Println("====================")
					fmt.Printf("Id: %d\n", contacto.idusuarios)
					fmt.Printf("Nombre: %s\n", contacto.nombre)
					fmt.Printf("Usuario: %s\n", contacto.usuario)
					fmt.Printf("Contraseña: %s\n", contacto.contraseña)
				}
			}
		case 3:
			fmt.Println("Ingresa el id:")
			fmt.Scanln(&c.idusuarios)
			fmt.Println("Ingresa el nuevo nombre:")
			if scanner.Scan() {
				c.nombre = scanner.Text()
			}
			fmt.Println("Ingresa el nueva usuario:")
			if scanner.Scan() {
				c.usuario = scanner.Text()
			}
			fmt.Println("Ingresa la nueva contraseña:")
			if scanner.Scan() {
				c.contraseña = scanner.Text()
			}
			err := actualizar(c)
			if err != nil {
				fmt.Printf("Error actualizando: %v", err)
			} else {
				fmt.Println("Actualizado correctamente")
			}
		case 4:
			fmt.Println("Ingresa el ID del contacto que deseas eliminar:")
			fmt.Scanln(&c.idusuarios)
			err := eliminar(c)
			if err != nil {
				fmt.Printf("Error eliminando: %v", err)
			} else {
				fmt.Println("Eliminado correctamente")
			}
		}
	}
}

func eliminar(c Contacto) error {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("DELETE FROM usuarios WHERE idusuarios = ?")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()

	_, err = sentenciaPreparada.Exec(c.idusuarios)
	if err != nil {
		return err
	}
	return nil
}

func insertar(c Contacto) (e error) {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	// Preparamos para prevenir inyecciones SQL
	sentenciaPreparada, err := db.Prepare("INSERT INTO usuarios (nombre, usuario, contraseña) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()
	// Ejecutar sentencia, un valor por cada '?'
	_, err = sentenciaPreparada.Exec(c.nombre, c.usuario, c.contraseña)
	if err != nil {
		return err
	}
	return nil
}

func obtenerContactos() ([]Contacto, error) {
	contactos := []Contacto{}
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	filas, err := db.Query("SELECT idusuarios, nombre, usuario, contraseña FROM usuarios")

	if err != nil {
		return nil, err
	}
	// Si llegamos aquí, significa que no ocurrió ningún error
	defer filas.Close()

	// Aquí vamos a "mapear" lo que traiga la consulta en el while de más abajo
	var c Contacto

	// Recorrer todas las filas, en un "while"
	for filas.Next() {
		err = filas.Scan(&c.idusuarios, &c.nombre, &c.usuario, &c.contraseña)
		// Al escanear puede haber un error
		if err != nil {
			return nil, err
		}
		// Y si no, entonces agregamos lo leído al arreglo
		contactos = append(contactos, c)
	}
	// Vacío o no, regresamos el arreglo de contactos
	return contactos, nil
}

func actualizar(c Contacto) error {
	db, err := obtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("UPDATE usuarios SET nombre = ?, usuario = ?, contraseña = ? WHERE idusuarios = ?")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()
	// Pasar argumentos en el mismo orden que la consulta
	_, err = sentenciaPreparada.Exec(c.nombre, c.usuario, c.contraseña, c.idusuarios)
	return err // Ya sea nil o sea un error, lo manejaremos desde donde hacemos la llamada
}
