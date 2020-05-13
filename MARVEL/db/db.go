package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type DBClient interface {
	NewDBBuilder() (*sql.DB, error)
}

type DB struct {
	Usuario           string
	Pass              string
	Host              string
	NombreBaseDeDatos string
	// GetConexion       func(drive string, estruct string) (db *sql.DB, e error)
}

func NewDB() DBClient {
	return &DB{
		Usuario:           "root",
		Pass:              "12345",
		Host:              "tcp(127.0.0.1:3306)",
		NombreBaseDeDatos: "prueba",
	}
}

func (User *DB) NewDBBuilder() (*sql.DB, error) {
	// U := DB{
	User.Usuario = "root"
	User.Pass = "12345"
	User.Host = "tcp(127.0.0.1:3306)"
	User.NombreBaseDeDatos = "prueba"
	// }
	db, sqlerr := sql.Open("mysql", fmt.Sprintf("%s:%s@%s/%s", User.Usuario, User.Pass, User.Host, User.NombreBaseDeDatos))
	if sqlerr != nil {
		return nil, sqlerr
	}
	return db, nil

}

// func NewDB() DBClientBuilder {
// 	return &dbClientBuilder{
// 		dbClient: &DB{
// 			Usuario:           DB_USER,
// 			Pass:              DB_PASS,
// 			Host:              DB_HOST,
// 			NombreBaseDeDatos: DB_NDB,
// 			Conexion: func(dirver, estruct) (db *sql.DB, e error) {
// 				db, err := sql.Open(nose, culquiera)
// 				if err != nil {
// 					return nil, err
// 				}
// 				return db, nil
// 			},
// 		},
// 	}
// }
