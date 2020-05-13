package repo

import (
	// "log"

	"log"

	"github.com/damian2205/GO_Capacitation/MARVEL/db"
	"github.com/damian2205/GO_Capacitation/MARVEL/models"
	"github.com/damian2205/GO_Capacitation/MARVEL/services"

	"github.com/damian2205/GO_Capacitation/MARVEL/dto"
)

type envioServices struct {
	interfaceDB db.DBClient
}

func NewEnvioServices(db db.DBClient) services.EnvioServices {
	return &envioServices{db}
}

func (dba *envioServices) ObtenerUser() (datos []dto.Contacto, e error) {
	contactos := []dto.Contacto{}
	db, err := dba.interfaceDB.NewDBBuilder()
	if err != nil {
		log.Printf("ERROR AL TRAER LA DB, %v", err)
	}
	defer db.Close()
	filas, err := db.Query("SELECT IDusuarios, Nombre, Usuario, Contraseña FROM usuarios")
	if err != nil {
		return nil, err
	}
	var c dto.Contacto

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

func (dba *envioServices) InsertarUser(d dto.Contacto) (e error) {
	UserModel := models.ModelContacto{Nombre: d.Nombre, Usuario: d.Usuario, Contraseña: d.Contraseña}
	db, err := dba.interfaceDB.NewDBBuilder()
	if err != nil {
		log.Printf("ERROR AL TRAER LA DB, %v", err)
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

func (dba *envioServices) EliminarUser(d dto.Contacto) (e error) {
	IDuser := models.ModelContacto{IDusuarios: d.IDusuarios, Nombre: d.Nombre, Usuario: d.Usuario, Contraseña: d.Contraseña}
	db, err := dba.interfaceDB.NewDBBuilder()
	if err != nil {
		log.Printf("ERROR AL TRAER LA DB, %v", err)
	}
	defer db.Close()

	sentenciaPreparada, err := db.Prepare("DELETE FROM usuarios WHERE IDusuarios = ?")
	if err != nil {
		return err
	}
	defer sentenciaPreparada.Close()
	log.Printf("User eliminado{ id:%b, name:%s, user:%s, pass:%s }", IDuser.IDusuarios, IDuser.Nombre, IDuser.Usuario, IDuser.Contraseña)

	_, err = sentenciaPreparada.Exec(IDuser.IDusuarios)
	if err != nil {
		return err
	}
	return nil
}

func (dba *envioServices) ActualizarUser(d dto.Contacto) (e error) {
	User := models.ModelContacto{IDusuarios: d.IDusuarios, Nombre: d.Nombre, Usuario: d.Usuario, Contraseña: d.Contraseña}
	db, err := dba.interfaceDB.NewDBBuilder()
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
	_, err = sentenciaPreparada.Exec(User.Nombre, User.Usuario, User.Contraseña, User.IDusuarios)
	return nil
}
