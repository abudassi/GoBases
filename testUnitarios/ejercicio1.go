package testUnitarios

/*
La  empresa de chocolates que anteriormente necesitaba calcular el impuesto de sus empleados,
al momento de depositar el sueldo de los mismos ahora nos solicitó validar que los cálculos
de estos impuestos están correctos. Para esto nos encargaron el trabajo de realizar los test correspondientes para:
Calcular el impuesto en caso de que el empleado gane por debajo de $50.000.
Calcular el impuesto en caso de que el empleado gane por encima de $50.000.
Calcular el impuesto en caso de que el empleado gane por encima de $150.000.
*/

func Impuesto(salario float64) float64 {
	if salario > 150000 {
		return salario * 0.27
	} else if salario > 50000 {
		return salario * 0.17
	} else {
		return 0.0
	}

}
