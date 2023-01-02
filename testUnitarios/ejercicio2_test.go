package testUnitarios

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromedio(t *testing.T) {

	//RESULTADO OK
	/*resultadoEsperado := 7.0
	resultado := Promedio(10, 9, 4, 5)
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")*/

	//RESULTADO FAIL
	resultadoEsperado := 8.0
	resultado := Promedio(10, 9, 4, 5)
	assert.Equal(t, resultadoEsperado, resultado, "deben ser iguales")

}
