package testUnitarios

import (
	"testing"

	"github.com/stretchr/testify/assert" 
)

func TestEstadisticas(t *testing.T) {

	//Se cargan los resultados esperados
	resultadoEsperado1 := 5.0
	resultadoEsperado2 := 10.0
	resultadoEsperado3 := 1.0

	// Se ejecuta el test, se pasan los valores a comparar
	resultado1 := Estadisticas("average", 0.0, 5.0, 10.0)
	resultado2 := Estadisticas("maximum", 10.0, 2.0, 2.0)
	resultado3 := Estadisticas("minimum", 1.0, 5.0, 5.0)

	// Se validan los resultados
	assert.Equal(t, resultadoEsperado1, resultado1, "deben ser iguales")
	assert.Equal(t, resultadoEsperado2, resultado2, "deben ser iguales")
	assert.Equal(t, resultadoEsperado3, resultado3, "deben ser iguales")

}
