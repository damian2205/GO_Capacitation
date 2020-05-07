package repo

import (
	ODB "github.com/damian2205/GO_Capacitation/MARVEL/db"
)

func ObtenerContactos() ([]Contacto, error) {
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
	var c Contacto

	for filas.Next() {
		err = filas.Scan(&c.idusurios, &c.nombre, &c.usuario, &c.contraseña)
		if err != nil {
			return nil, err
		}
		contactos = append(contactos, c)
	}
	defer filas.Close()

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
