package main

import (
	"errors"
	"fmt"
)

func main() {

	/*
	   EJERCICIO 1
	   En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
	   Creá un error personalizado con un struct que implemente “Error()” con el mensaje
	   “Error: el salario ingresado no alcanza el mínimo imponible" y lanzalo en caso de que “salary”
	   sea menor a 150.000. De lo contrario, tendrás que imprimir por consola el mensaje “Debe pagar impuesto”.
	*/

	/*salary := 100000
	salary := 200000
	err := ChechkSalary(salary)
	if err != nil{
		fmt.Println(err)
		// EL RETURN ESTA AGREGADO PARA QUE NO IMPRIMA EL DEBE PAGAR IMPUESTO
		return
	}
	fmt.Println("Debe pagar impuesto")
	*/

	/*
	   EJERCICIO 2
	   En tu función “main”, definí una variable llamada “salary” y asignale un valor de tipo “int”.
	   Creá un error personalizado con un struct que implemente “Error()” con el mensaje
	   “Error: el salario es menor a 10.000" y lanzalo en caso de que “salary” sea menor o igual a  10.000.
	   La validación debe ser hecha con la función “Is()” dentro del “main”.
	*/

	/*salary := 10000
	salary := 15000
	err := ChechkSalary(salary)
	if errors.Is(err, &ErrorSalary{}) {
		fmt.Println(err)
		return
	}
	fmt.Println("No hay error")
	*/

	/*
	   EJERCICIO 3
	   Hacé lo mismo que en el ejercicio anterior pero reformulando el código para que, en reemplazo de “Error()”,  se implemente “errors.New()”.
	*/

	/*
		salary := 10000
		//salary := 15000
		err := ChechkSalary(salary)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("No hay error")
	*/

	/*
	   EJERCICIO 4
	   Repetí el proceso anterior, pero ahora implementando “fmt.Errorf()”,
	   para que el mensaje de error reciba por parámetro el valor de “salary”
	   indicando que no alcanza el mínimo imponible (el mensaje mostrado por consola deberá decir: “Error: el mínimo
	   imponible es de 150.000 y el salario ingresado es de: [salary]”, siendo [salary] el valor de tipo int pasado por parámetro).
	*/

	/*
		//salary := 10000
		salary := 160000
		err := ChechkSalary(salary)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("No hay error")
	*/

	/*
	   EJERCICIO 5
	   Desarrollá las funciones necesarias para permitir a la empresa calcular:
	   Salario mensual de un trabajador según la cantidad de horas trabajadas.
	   	- La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
	   	- Dicha función deberá retornar más de un valor (salario calculado y error).
	   	- En caso de que el salario mensual sea igual o superior a $150.000, se le deberá descontar el 10 % en concepto de impuesto.
	   	- En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo, la función debe devolver un error.
	   			El mismo tendrá que indicar “Error: el trabajador no puede haber trabajado menos de 80 hs mensuales”.
	*/

	result, err := MonthlySalary(500, 6000)
	if err != nil {
		panic(err)
	}
	fmt.Printf("El salario es: %f\n", result)
}

// EJERCICIO 1
/*
type ErrorSalary struct {
}

func ChechkSalary(salary int) error {
	if salary < 150000 {
		return &ErrorSalary{}
	}
	return nil
}

func (e *ErrorSalary) Error() string {
	return "Error: el salario ingresado no alcanza el mínimo imponible"
}*/

// EJERCICIO 2
/*type ErrorSalary struct{}

func ChechkSalary(salary int) error {
	if salary <= 10000 {
		return &ErrorSalary{}
	}
	return nil
}

func (e *ErrorSalary) Error() string {
	return "Error: el salario es menor a 10000"
}
*/

//EJERCICIO 3
/*
func ChechkSalary(salary int) error {
	if salary <= 10000 {
		return errors.New("Error: el salario es menor a 10.000")
	}
	return nil
}
*/

// EJERCICIO 4
/*
func ChechkSalary(salary int) error {
	if salary <= 150000 {
		return fmt.Errorf("Error: el mínimo imponible es de 150.000 y el salario ingresado es de: %d", salary)
	}
	return nil
}
*/

//EJERCICIO 5

func MonthlySalary(hours int, salary float64) (result float64, err error) {
	if hours < 80 {
		err = errors.New("El trabajador no puede haber trabajado menos de 80 hs mensuales")
		return
	}
	result = float64(hours) * salary
	if result >= 150000 {
		result -= result * 0.1
	}
	return
}
