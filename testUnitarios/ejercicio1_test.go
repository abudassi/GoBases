package testUnitarios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImpuesto(t *testing.T) {
	// Se inicializan los datos a usar en el test y los esperados
	salario1 := 15000.0
	salario2 := 70000.0
	salario3 := 170000.0
	impuestoEsperado1 := 0.0
	impuestoEsperado2:= 11900.0
	impuestoEsperado3 := 45900.0

	// Se ejecuta el test
	resultado1 := Impuesto(salario1)
	resultado2 := Impuesto(salario2)
	resultado3 := Impuesto(salario3)

	// Se comparan los resultados
	assert.Equal(t, impuestoEsperado1, resultado1, "deben ser iguales")
	assert.Equal(t, impuestoEsperado2, resultado2, "deben ser iguales")
	assert.Equal(t, impuestoEsperado3, resultado3, "deben ser iguales")

}
