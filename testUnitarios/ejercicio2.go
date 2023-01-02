package testUnitarios

/*
El colegio informó que las operaciones para calcular el promedio no se están realizando correctamente,
por lo que ahora nos corresponde realizar los test correspondientes:
Calcular el promedio de las notas de los alumnos.
*/

func Promedio(notas ...float64) float64 {
	var sumatoria float64
	for _, nota := range notas {
		if nota >= 0 {
			sumatoria += nota
		}
	}
	return sumatoria / float64(len(notas))
}