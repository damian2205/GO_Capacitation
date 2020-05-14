package calculadora_test

import (
	"testing"

	"github.com/damian2205/GO_Capacitation/calculadora"
)

func TestSuma(t *testing.T) {
	data := []struct {
		a, b int
	}{
		{2, 3},
	}
	for _, d := range data {
		suma := calculadora.Suma(d.a, d.b)
		if suma != 5 {
			t.Error("La prueba faio")
			t.Fail()
		} else {
			t.Log("YE MEN")
		}

	}

}
