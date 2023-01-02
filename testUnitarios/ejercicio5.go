package testUnitarios

/*
El refugio de animales envió una queja ya que el cálculo total de alimento a comprar no fue el correcto y
compraron menos alimento del que necesitaban. Para mantener satisfecho a nuestro cliente deberemos realizar
los test necesarios para verificar que todo funcione correctamente:
Verificar el cálculo de la cantidad de alimento para los perros.
Verificar el cálculo de la cantidad de alimento para los gatos.
Verificar el cálculo de la cantidad de alimento para los hamster.
Verificar el cálculo de la cantidad de alimento para las tarántulas.
*/

func CalcularAlimento(animal string, cantidad int) int {
	switch animal {
	case "perro":
		return perro(cantidad)
	case "gato":
		return gato(cantidad)
	case "hamster":
		return hamster(cantidad)
	case "tarantula":
		return tarantula(cantidad)
	default:
		return 0.0
	}
}

func perro(cantidad int) int {
	return cantidad * 10000
}

func gato(cantidad int) int {
	return cantidad * 5000
}
func hamster(cantidad int) int {
	return cantidad * 250
}
func tarantula(cantidad int) int {
	return cantidad * 150
}
