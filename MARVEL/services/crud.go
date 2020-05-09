package repo

import (
	"log"

	ODB "github.com/damian2205/GO_Capacitation/MARVEL/db"
	"github.com/damian2205/GO_Capacitation/MARVEL/dto"
	"github.com/damian2205/GO_Capacitation/MARVEL/models"
)

func ObtenerUser() ([]models.ModelContacto, error) {
	contactos := []models.ModelContacto{}
	db, err := ODB.ObtenerBaseDeDatos()
	if err != nil {
		return nil, err
	}
	defer db.Close()
	filas, err := db.Query("SELECT IDusuarios, Nombre, Usuario, Contraseña FROM usuarios")
	if err != nil {
		return nil, err
	}
	var c models.ModelContacto

	for filas.Next() {
		err = filas.Scan(&c.IDusuarios, &c.Nombre, &c.Usuario, &c.Contraseña)
		if err != nil {
			return nil, err
		}
		contactos = append(contactos, c)
	}
	defer filas.Close()

	return contactos, nil
}

func InsertarUser(d dto.Contacto) (e error) {
	UserModel := models.ModelContacto{Nombre: d.Nombre, Usuario: d.Usuario, Contraseña: d.Contraseña}
	db, err := ODB.ObtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("INSERT INTO usuarios (Nombre, Usuario, Contraseña) VALUES(?, ?, ?)")
	if err != nil {
		return err
	}

	defer sentenciaPreparada.Close()
	log.Printf("%b,%s,%s,%s", UserModel.IDusuarios, UserModel.Nombre, UserModel.Usuario, UserModel.Contraseña)

	_, err = sentenciaPreparada.Exec(UserModel.Nombre, UserModel.Usuario, UserModel.Contraseña)
	if err != nil {
		return err
	}
	return nil

}

func EliminarUser(d dto.Contacto) error {
	IDuser := models.ModelContacto{IDusuarios: d.IDusuarios, Nombre: d.Nombre, Usuario: d.Usuario, Contraseña: d.Contraseña}
	db, err := ODB.ObtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("DELETE FROM usuarios WHERE IDusuarios = ?")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()
	// log.Printf("User eliminado{ id:%b, name:%s, user:%s, pass:%s }", IDuser.IDusuarios, IDuser.Nombre, IDuser.Usuario, IDuser.Contraseña)

	_, err = sentenciaPreparada.Exec(IDuser.IDusuarios)
	if err != nil {
		return err
	}
	return nil
}

func ActualizarUser(d dto.Contacto) error {
	User := models.ModelContacto{IDusuarios: d.IDusuarios, Nombre: d.Nombre, Usuario: d.Usuario, Contraseña: d.Contraseña}
	db, err := ODB.ObtenerBaseDeDatos()
	if err != nil {
		return err
	}
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("UPDATE usuarios SET Nombre = ?, Usuario = ?, Contraseña = ? WHERE IDusuarios = ?")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()
	log.Printf("Ususario acualizado: %b, %s, %s, %s", User.IDusuarios, User.Nombre, User.Usuario, User.Contraseña)
	// _, err = sentenciaPreparada.Exec(User.IDusuarios)
	_, err = sentenciaPreparada.Exec(User.IDusuarios, User.Nombre, User.Usuario, User.Contraseña)
	return nil
}
