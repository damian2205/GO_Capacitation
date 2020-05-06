package main

import (
	"bufio"
	"fmt"
	"os"

	ODB "github.com/damian2205/GO_Capacitation/MARVEL/db"
	_ "github.com/go-sql-driver/mysql"
)

type Contacto struct {
	idusurios                   int
	nombre, usuario, contraseña string
}

func main() {
	db, err := ODB.ObtenerBaseDeDatos()
	if err != nil {
		fmt.Printf("Error obteniendo base de datos: %v", err)
		return
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		fmt.Printf("Error conectando: %v", err)
		return
	}
	fmt.Printf("Conectado correctamente\n")

	menu := `Escoge una opcion del CRUD
	[1] -- Mostrar datos de la DB
	[2] -- Insertar datos a la DB
	[3] -- Eliminar un dato de la DB 
	[4] -- Actualizar un dato de la DB
	[5] -- Salir del menú`

	var eleccion int
	var c Contacto
	for eleccion != 5 {
		fmt.Println(menu)
		fmt.Scan(&eleccion)
		scanner := bufio.NewScanner(os.Stdin)

		switch eleccion {
		case 1:
			// Llamado funcion para obtener los datos de la DB
			contactos, err := obtenerContactos()
			if err != nil {
				fmt.Printf("Error obteniendo contactos: %v", err)
				return
			}
			for _, contacto := range contactos {
				fmt.Printf("\n%v\n\n", contacto)
			}

		case 2:
			//Insertar usuario
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
			_err := InsertarUser(c)
			if _err != nil {
				fmt.Printf("Error insertando: %v", _err)
			} else {
				fmt.Println("Insertado correctamente")
			}

		case 3:
			//Eliminar Usuario
			fmt.Println("Ingresa el ID del usuario que deseas eliminar:")
			fmt.Scanln(&c.idusurios)
			err_ := EliminarUser(c)
			if err_ != nil {
				fmt.Printf("Error eliminando: %v", err_)
			} else {
				fmt.Println("Eliminado correctamente")
			}

		case 4:
			fmt.Println("Ingresa el id:")
			fmt.Scanln(&c.idusurios)
			fmt.Println("Ingresa el nuevo nombre:")
			if scanner.Scan() {
				c.nombre = scanner.Text()
			}
			fmt.Println("Ingresa el nuevo usuario:")
			if scanner.Scan() {
				c.usuario = scanner.Text()
			}
			fmt.Println("Ingresa la nueva contraseña:")
			if scanner.Scan() {
				c.contraseña = scanner.Text()
			}
			err := ActualizarUser(c)
			if err != nil {
				fmt.Printf("Error actualizando: %v", err)
			} else {
				fmt.Println("Actualizado correctamente")
			}
		}
	}

}

func obtenerContactos() ([]Contacto, error) {
	contactos := []Contacto{}
	db, err := ODB.ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	filas, err := db.Query("SELECT idusuarios, nombre, usuario, contraseña FROM usuarios")

	if err != nil {
		return nil, err
	}
	defer filas.Close()

	var c Contacto

	for filas.Next() {
		err = filas.Scan(&c.idusurios, &c.nombre, &c.usuario, &c.contraseña)
		if err != nil {
			return nil, err
		}
		contactos = append(contactos, c)
	}
	return contactos, nil
}

func InsertarUser(c Contacto) (e error) {
	db, err := ODB.ObtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("INSERT INTO usuarios (nombre, usuario, contraseña) VALUES(?, ?, ?)")
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

func EliminarUser(c Contacto) error {
	db, err := ODB.ObtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("DELETE FROM usuarios WHERE idusuarios = ?")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()

	_, err = sentenciaPreparada.Exec(c.idusurios)
	if err != nil {
		return err
	}
	return nil
}

func ActualizarUser(c Contacto) error {
	db, err := ODB.ObtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("UPDATE usuarios SET nombre = ?, usuario = ?, contraseña = ? WHERE idusuarios = ?")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()
	_, err = sentenciaPreparada.Exec(c.nombre, c.usuario, c.contraseña, c.idusurios)
	return err
}
