package db_test

import (
	"fmt"
	"testing"

	"github.com/damian2205/GO_Capacitation/MARVEL/db"
)

type DBase struct {
	base              db.DBClient
	Usuario           string
	Pass              string
	Host              string
	NombreBaseDeDatos string
}

func (dba *DBase) TestNewDBBuiulder(t *testing.T) {
	dba.Usuario = "root"
	dba.Pass = "12345"
	dba.Host = "tcp(127.0.0.1:3306)"
	dba.NombreBaseDeDatos = "prueba"
	cosas := []struct {
		name    string
		errMsg  string
		sqlType string
		sqlPath string
	}{
		{
			name:    "NewDBBuilder/OK",
			sqlType: "mysql",
			sqlPath: fmt.Sprintf("%s:%s@%s/%s", dba.Usuario, dba.Pass, dba.Host, dba.NombreBaseDeDatos),
		},
		{
			name:   "NewDBBuilder/OpenError",
			errMsg: "sql: driver desconocido \"%s\"olvido importar la libreria",
		},
	}

	for _, valor := range cosas {
		db, err := dba.base.NewDBBuilder()
		if err != nil {
			t.Error("ERORRR, %v", err)
		}
	}

}
