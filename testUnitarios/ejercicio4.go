package testUnitarios

/*
Los profesores de la universidad de Colombia, entraron al programa de análisis de datos  de Google,
el cual premia a los mejores estadísticos de la región. Por ello los profesores nos solicitaron
comprobar el correcto funcionamiento de nuestros cálculos estadísticos. Se solicita la siguiente tarea:
Realizar test para calcular el mínimo de calificaciones.
Realizar test para calcular el máximo de calificaciones.
Realizar test para calcular el promedio de calificaciones.
*/

func Estadisticas(operacion string, notas ...float64) float64 {
	switch operacion {
	case "minimum":
		return opMinVal(notas...)
	case "average":
		return opAvgVal(notas...)
	case "maximum":
		return opMaxVal(notas...)
	default:
		return 0.0
	}
}

func opMinVal(values ...float64) (value float64) {
	value = values[0]
	for _, number := range values {
		if number < value {
			value = number
		}
	}
	return
}

func opAvgVal(values ...float64) (value float64) {
	sum := 0.0
	for _, number := range values {
		sum += number
	}
	return sum / float64(len(values))
}

func opMaxVal(values ...float64) (value float64) {
	value = values[0]
	for _, number := range values {
		if number > value {
			value = number
		}
	}
	return
}