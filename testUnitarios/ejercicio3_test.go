package testUnitarios

import (
	//Importar testing
	"testing"
	// Importar testify: recordar colocar comando en consola -> go get github.com/stretchr/testify
	"github.com/stretchr/testify/assert"
)

func TestSalario(t *testing.T) {
	// Se inicializan los datos a testear en cada categoria mas los datos esperados
	categoriaA := "A"
	horasA := 300.0
	resultadoEsperadoA := 1350000.0

	categoriaB := "B"
	horasB := 300.0
	resultadoEsperadoB := 540000.0

	categoriaC := "C"
	horasC := 300.0
	resultadoEsperadoC := 300000.0

	// Se ejecutan los test
	resultadoA := Salario(categoriaA, horasA)
	resultadoB := Salario(categoriaB, horasB)
	resultadoC := Salario(categoriaC, horasC)

	// Se comparan los resultados
	assert.Equal(t, resultadoEsperadoA, resultadoA, "deben ser iguales")
	assert.Equal(t, resultadoEsperadoB, resultadoB, "deben ser iguales")
	assert.Equal(t, resultadoEsperadoC, resultadoC, "deben ser iguales")

}
