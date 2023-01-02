package main

import "fmt"

func main() {
	/*
	   EJERCICIO 1
	   La Real Academia Española quiere saber cuántas letras tiene una palabra y luego tener cada una de las letras por separado para deletrearla. Para eso tendrán que:
	   1 Crear una aplicación que reciba  una variable con la palabra e imprimir la cantidad de letras que contiene la misma.
	   2 Luego, imprimí cada una de las letras.
	*/

	var palabra string
	var cantidadLetras int
	fmt.Println("Ingrese la palabra")
	fmt.Scanf("%v\n", &palabra)
	//fmt.Println("La palabra es", palabra)
	for i, letra := range palabra {
		fmt.Println(string(letra))
		cantidadLetras = i + 1
	}
	fmt.Println("La palabra ", palabra, " tiene ", cantidadLetras, " letras")

	/*
	   EJERCICIO 2
	   Un banco quiere otorgar préstamos a sus clientes, pero no todos pueden acceder a los mismos.
	   Para ello tiene ciertas reglas para saber a qué cliente se le puede otorgar.
	   Solo le otorga préstamos a clientes cuya edad sea mayor a 22 años, se encuentren empleados y tengan más de un año de antigüedad en su trabajo.
	   Dentro de los préstamos que otorga no les cobrará interés a los que posean un sueldo superior a $100.000.
	   Es necesario realizar una aplicación que reciba  estas variables y que imprima un mensaje de acuerdo a cada caso.
	   Tip: tu código tiene que poder imprimir al menos 3 mensajes diferentes.
	*/

	var edad int
	var opcion int
	var tieneTrabajo bool
	var antiguedad int
	var salario int

	fmt.Println("Ingrese su edad")
	fmt.Scanf("%d\n", &edad)

	fmt.Println("Tiene trabajo? Coloque 1 Si tiene, si no presione 2")
	fmt.Scanf("%d\n", &opcion)

	if opcion != 1 && opcion != 2 {
		for opcion != 1 && opcion != 2 {
			fmt.Println("La opcion es incorrecta. Tiene trabajo? Coloque 1 Si tiene, si no presione 2")
			fmt.Scanf("%d\n", &opcion)
		}
	}

	if opcion == 1 {
		tieneTrabajo = true
	} else {
		tieneTrabajo = false
	}

	if opcion == 1 {
		fmt.Println("Ingrese la antiguedad en años de trabajo")
		fmt.Scanf("%d\n", &antiguedad)
	} else {
		antiguedad = 0
	}

	if opcion == 1 {
		fmt.Println("De cuanto es su salario")
		fmt.Scanf("%d\n", &salario)
	} else {
		antiguedad = 0
	}

	switch {
	case edad >= 22 && tieneTrabajo == true && antiguedad > 1 && salario < 100000:
		fmt.Println("Usted puede acceder a un prestamo. Se le cobrara interés.")
	case edad >= 22 && tieneTrabajo == true && antiguedad > 1 && salario >= 100000:
		fmt.Println("Usted puede acceder a un prestamo, sin el cobro de interés.")
	default:
		fmt.Println("Usted no puede acceder a un préstamo.")
	}

	/*
	   EJERCICIO 3
	   Realizar una aplicación que reciba  una variable con el número del mes.
	   Según el número, imprimir el mes que corresponda en texto.
	   ¿Se te ocurre que se puede resolver de distintas maneras? ¿Cuál elegirías y por qué?
	   Si, se puede con if, switch o for. Yo elegi switch porque la cantidad de alternativas es limitada en 12. Si fuera en dos eligiria un if pero
	   al ser 12 seria un if anidado muy largo. En el caso de ser ilimitadas las opciones eligiria un for.
	   Creo tambien que se podria hacer un array en el cual se lea la posicion -1 que se ingrese
	   Ej: 7, Julio
	   Nota: Validar que el número del mes, sea correcto.
	*/

	var numeroMes int

	fmt.Println("Ingrese el numero de mes")
	fmt.Scanf("%d\n", &numeroMes)

	if numeroMes <= 1 || numeroMes >= 12 {
		for numeroMes <= 1 || numeroMes >= 12 {
			fmt.Println("La opcion es incorrecta. Ingrese un numero correcto entre 1 y 12")
			fmt.Scanf("%d\n", &numeroMes)
		}
	}

	switch numeroMes {
	case 1:
		fmt.Println("Enero")
	case 2:
		fmt.Println("Febrero")
	case 3:
		fmt.Println("Marzo")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Mayo")
	case 6:
		fmt.Println("Junio")
	case 7:
		fmt.Println("Julio")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Septiembre")
	case 10:
		fmt.Println("Octubre")
	case 11:
		fmt.Println("Noviembre")
	case 12:
		fmt.Println("Diciembre")
	default:
		fmt.Println("El mes ingresado NO existe")
	}

	/*
	   EJERCICIO 4
	   Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa, ayudá a imprimir la edad de Benjamin.
	     var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	   Por otro lado también es necesario:
	   Saber cuántos de sus empleados son mayores de 21 años.
	   Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
	   Eliminar a Pedro del mapa.

	*/

	//crear mapa
	var empleados = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}
	fmt.Println(empleados)

	//imprimir la edad de Benjamin
	fmt.Println("La edad de Benjamin es de", empleados["Benjamin"], "años.")

	//cuantos empleados son mayor a 21 años
	cantidadEmpleadosMayores := 0
	for _, empleado := range empleados {
		if empleado >= 21 {
			cantidadEmpleadosMayores++
		}
	}
	fmt.Println("La cantidad de empleados mayores a 21 años es de", cantidadEmpleadosMayores)

	//agregar a Federico
	empleados["Federico"] = 25
	fmt.Println(empleados)

	//eliminar a Pedro
	delete(empleados, "Pedro")
	fmt.Println(empleados)
}
