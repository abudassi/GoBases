package testUnitarios

import (
	"testing"

	"github.com/stretchr/testify/assert" // Se importa testify
)

func Test_perro(t *testing.T) {

	//Arrange. Datos a testear
	animal := "perro"
	cantidad := 5
	resultado1 := 50000

	//Act. Inicialización del test
	resultadoEsperado1 := CalcularAlimento(animal, cantidad)

	//Assert. Comparación
	assert.Equal(t, resultadoEsperado1, resultado1, "deben ser iguales")

}

func Test_loro(t *testing.T) {

	//Resultado FAIL no hay loro
	animal := "loro"
	cantidad := 5
	resultado1 := 50000

	resultadoEsperado1 := CalcularAlimento(animal, cantidad)

	assert.Equal(t, resultadoEsperado1, resultado1, "deben ser iguales")

}

func TestPerro(t *testing.T) {

	//Resultado FAIL
	cantidad := 6
	resultadoEsperado1 := 50000

	resultado1 := perro(cantidad)

	assert.Equal(t, resultadoEsperado1, resultado1, "deben ser iguales")

}

func TestGato(t *testing.T) {

	//Resultado OK
	cantidad := 5
	resultadoEsperado1 := 25000

	resultado1 := gato(cantidad)

	assert.Equal(t, resultadoEsperado1, resultado1, "deben ser iguales")

}

func TestHamster(t *testing.T) {

	//Resultado FAIL
	cantidad := 8
	resultadoEsperado1 := 1250

	resultado1 := hamster(cantidad)

	assert.Equal(t, resultadoEsperado1, resultado1, "deben ser iguales")

}

func TestTarantula(t *testing.T) {

	//Resultado OK
	cantidad := 10
	resultadoEsperado1 := 1500

	resultado1 := tarantula(cantidad)

	assert.Equal(t, resultadoEsperado1, resultado1, "deben ser iguales")

}
