package main

import "fmt"

func main() {
	/*
	   EJERCICIO 1
	   Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de
	   depositar el sueldo, para cumplir el objetivo es necesario crear una función que devuelva
	   el impuesto de un salario.
	   Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del
	   sueldo y si gana más de $150.000 se le descontará además un 10 % (27% en total).
	*/

	fmt.Println("El impuesto del salario es de $", impuestoSalario(40000))
	//fmt.Println("El impuesto del salario es de $", impuestoSalario(90000))
	//fmt.Println("El impuesto del salario es de $", impuestoSalario(200000))

	/*
		EJERCICIO 2
		Un colegio necesita calcular el promedio (por estudiante) de sus calificaciones.
		Se solicita generar una función en la cual se le pueda pasar N cantidad de enteros y
		devuelva el promedio. No se pueden introducir notas negativas.
	*/

	//MEJORAR. AVERIGUAR COMO PODER GUARDAR VALORES EN UN SLICE DESDE UNA FUNCION

	/*var cantidad int
	fmt.Println("Cuantas notas va a ingresar")
	fmt.Scanf("%d\n", &cantidad)

	notas := make([]int, cantidad)
	for i := 0; i < len(notas); i++ {
		fmt.Println("Ingrese la nota")
		fmt.Scan(&notas[i])
		if notas[i] < 0 {
			for notas[i] < 0 {
				fmt.Println("No es valido ingresar notas negativas. Ingrese una nota valida")
				fmt.Scan(&notas[i])
			}
		}
	}

	fmt.Println("El promedio del alumno es de ", promedio(notas...))*/

	/*
	   EJERCICIO 3
	   Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.
	   1 Categoría C, su salario es de $1.000 por hora.
	   2 Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
	   3 Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.
	   Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes, la categoría y que devuelva su salario.
	*/
	/*
		var categoria string
		var horasTrabajadas float64

		fmt.Println("En que categoria se encuentra? Ingrese A, B o C")
		fmt.Scanf("%v\n", &categoria)

		if categoria != "A" && categoria != "B" && categoria != "C" {
			for categoria != "A" && categoria != "B" && categoria != "C" {
				fmt.Println("La opcion es incorrecta. INGRESE UNA CATEGORIA CORRECTA (A,B,C)")
				fmt.Scanf("%v\n", &categoria)
			}
		}

		fmt.Println("Ingrese la cantidad de horas trabajadas")
		fmt.Scanf("%v\n", &horasTrabajadas)

		fmt.Println("Su salario mensual es de", calcularSalario(categoria, horasTrabajadas))
	*/

	/*
		EJERCICIO 4
		Los profesores de una universidad de Colombia necesitan calcular algunas estadísticas de calificaciones de los/as estudiantes de un curso.
		Requieren calcular los valores mínimo, máximo y promedio de sus calificaciones.
		Para eso, se solicita generar una función que indique qué tipo de cálculo se quiere realizar (mínimo, máximo o promedio)
		y que devuelva otra función y un mensaje (en caso que el cálculo no esté definido) que se le pueda pasar una cantidad N de enteros
		y devuelva el cálculo que se indicó en la función anterior.
	*/

	/*
		var operacion string
		var cantidad int
		fmt.Println("Cuantas notas va a ingresar")
		fmt.Scanf("%d\n", &cantidad)

		notas := make([]float64, cantidad)
		for i := 0; i < len(notas); i++ {
			fmt.Println("Ingrese la nota")
			fmt.Scan(&notas[i])
			if notas[i] < 0 {
				for notas[i] < 0 {
					fmt.Println("No es valido ingresar notas negativas. Ingrese una nota valida")
					fmt.Scan(&notas[i])
				}
			}
		}

		fmt.Println("Que operación desea realizar (minimum, maximum, average)")
		fmt.Scanf("%v\n", &operacion)

		if operacion != "minimum" && operacion != "maximum" && operacion != "average" {
			for operacion != "minimum" && operacion != "maximum" && operacion != "average" {
				fmt.Println("La opcion es incorrecta. INGRESE OPERACION CORRECTA (minimum, maximum, average)")
				fmt.Scanf("%v\n", &operacion)
			}
		}

		fmt.Println("El ", operacion, " es ",estadisticas(operacion, notas...))
	*/

	/*
		EJERCICIO 5
		Un refugio de animales necesita calcular cuánto alimento debe comprar para las mascotas.
		Por el momento solo tienen tarántulas, hamsters, perros y gatos, pero se espera que puedan darle refugio a muchos animales más.
		Tienen los siguientes datos:
			Perro 10 kg de alimento.
			Gato 5 kg de alimento.
			Hamster 250 g de alimento.
			Tarántula 150 g de alimento.
		Se solicita:
		Implementar una función Animal que reciba como parámetro un valor de tipo texto con el animal especificado y
		que retorne una función y un mensaje (en caso que no exista el animal).
		Una función para cada animal que calcule la cantidad de alimento en base a la cantidad del tipo de animal especificado.
		Ejemplo:

			const (
		   dog    = "dog"
		   cat    = "cat"
		)

		...

		animalDog, msg := animal(dog)
		animalCat, msg := animal(cat)

		...

		var amount float64
		amount += animalDog(10)
		amount += animalCat(10)

	*/

	var cantidad int
	var animal string

	fmt.Println("Ingrese un animal")
	fmt.Scanf("%v\n", &animal)

	fmt.Printf("Cuantos %v quiere ingresar?", animal)
	fmt.Scanf("%d\n", &cantidad)

	fmt.Println("Se necesitan", calcularAlimento(animal, cantidad), " gramos para", animal)

}

// FUNCION EJERCICIO 1
func impuestoSalario(salario float32) float32 {
	var impuesto float32
	if salario >= 50000 && salario < 150000 {
		impuesto = salario * 0.17
	} else if salario >= 150000 {
		impuesto = salario * 0.27
	} else {
		impuesto = 0
	}
	return impuesto
}

// FUNCION EJERCICIO 2
func promedio(notas ...int) int {
	var prom int
	var sumatoria int
	prom = 0
	sumatoria = 0

	for index := range notas {
		sumatoria += notas[index]
	}

	cantidadNotas := len(notas)
	prom = sumatoria / cantidadNotas

	return prom
}

/*
func Promedio(notas ...float64) float64 {
	var suma_notas float64
	for _, nota := range notas {
		if nota >= 0 {
			suma_notas += nota
		}
	}
	return suma_notas / float64(len(notas))
}
*/

// FUNCION EJERCICIO 3
func calcularSalario(categoria string, horasTrabajadas float64) float64 {
	switch categoria {
	case string('C'):
		return horasTrabajadas * 1000.0
	case string('B'):
		return (horasTrabajadas * 1500.0) + (horasTrabajadas*1500.0)*0.20
	case string('A'):
		return (horasTrabajadas * 3000.0) + (horasTrabajadas*3000.0)*0.50
	default:
		return 0
	}
}

// EJERCICIO 4
/*const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)*/

func estadisticas(operacion string, notas ...float64) float64 {
	/*
			minFunc, err := operation(operacion)
			if err != nil {
				fmt.Println(err)
			} else {
				println(minFunc)
			}

			averageFunc, err := operation(operacion)
			if err != nil {
				fmt.Println(err)
			} else {
				println(averageFunc)
			}

			maxFunc, err := operation(operacion)
			if err != nil {
				fmt.Println(err)
			} else {
				println(maxFunc)
			}

			switch operacion {
			case "minimun":
				return minFunc(notas...)
			case "maximum":
				return maxFunc(notas...)
			case "avegage":
				return averageFunc(notas...)
			default:
				return 0
			}
		}
	*/

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

// EJERCICIO 5
func calcularAlimento(animal string, cantidad int) float64 {
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

func perro(cantidad int) float64 {
	fmt.Println(float64(cantidad) * 10000)
	return float64(cantidad) * 10000
}

func gato(cantidad int) float64 {
	fmt.Println(float64(cantidad) * 5000)
	return float64(cantidad) * 5000
}
func hamster(cantidad int) float64 {
	fmt.Println(float64(cantidad) * 250)
	return float64(cantidad) * 250
}
func tarantula(cantidad int) float64 {
	fmt.Println(float64(cantidad) * 150)
	return float64(cantidad) * 150
}
