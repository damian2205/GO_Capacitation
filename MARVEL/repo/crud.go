package repo

import (
	"log"

	ODB "github.com/damian2205/GO_Capacitation/MARVEL/db"
)

// type Contacto struct {
// 	IDusurios  int    `json:"idusarios"`
// 	Nombre     string `json:"nombre"`
// 	Usuario    string `json:"usuario"`
// 	Contraseña string `json:"contraseña"`
// }

type ModelContacto struct {
	IDusurios  int
	Nombre     string
	Usuario    string
	Contraseña string
}

type ResponseDTO struct {
	Message string `json:"message"`
	Data    string `json:"data"`
}

func ObtenerUser() ([]Contacto, error) {
	contactos := []Contacto{}
	db, err := ODB.ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	filas, err := db.Query("SELECT * FROM usuarios")
	if err != nil {
		return nil, err
	}
	var c Contacto

	for filas.Next() {
		err = filas.Scan(&c.IDusurios, &c.Nombre, &c.Usuario, &c.Contraseña)
		if err != nil {
			return nil, err
		}
		contactos = append(contactos, c)
	}
	defer filas.Close()

	return contactos, nil
}

func InsertarUser(d Contacto) (e error) {
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

	// _, err = sentenciaPreparada.Exec(d.Nombre, d.Usuario, d.Contraseña)
	// if err != nil {
	// 	return err
	// }

	userModel := ModelContacto{Nombre: d.Nombre, Usuario: d.Usuario, Contraseña: d.Contraseña}
	log.Println("OKAY SE CREO", userModel)
	// return ResponseDTO{Message: "se creo", Data: ""}

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

	_, err = sentenciaPreparada.Exec(c.IDusurios)
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
	_, err = sentenciaPreparada.Exec(c.Nombre, c.Usuario, c.Contraseña, c.IDusurios)
	return err
}
