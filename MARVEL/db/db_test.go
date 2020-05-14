package db_test

import (
	"fmt"
	"testing"

	"github.com/damian2205/GO_Capacitation/MARVEL/db"
	"github.com/stretchr/testify/assert"
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
		db, error := dba.base.NewDBBuilder()
		if error != nil {
			message := fmt.Sprintf(valor.errMsg)
			assert.EqualValues(t, message)
			fmt.Printf("ERORRR, %v", error)
		}
	}

}
