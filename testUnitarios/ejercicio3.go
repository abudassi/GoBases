package testUnitarios

/*
La empresa marinera no está de acuerdo con los resultados obtenidos en los cálculos de los salarios,
por ello nos piden realizar una serie de tests sobre nuestro programa.
Necesitaremos realizar las siguientes pruebas en nuestro código:
Calcular el salario de la categoría “A”.
Calcular el salario de la categoría “B”.
Calcular el salario de la categoría “C”.
*/

func Salario(categoria string, horas float64) float64 {
	switch categoria {
	case string('A'):
		return (horas * 3000.0) + (horas * 3000 * 0.5)
	case string('B'):
		return (horas * 1500.0) + (horas * 1500 * 0.2)
	case string('C'):
		return (horas * 1000.0)
	default:
		return 0
	}
}
