package main

import "fmt"

func main() {
	var (
		nombre    = "Andrea"
		direccion = "El Challao, Mendoza, Argentina"
	)

	fmt.Println("Mi nombre es ", nombre, " y vivo en ", direccion)

	/* Una empresa de meteorología quiere una aplicación donde pueda tener
	   la temperatura, humedad y presión atmosférica de distintos lugares.
	   Declará 3 variables especificando el tipo de dato, como valor deben tener
	   la temperatura, humedad y presión de donde te encuentres.
	   Imprimí los valores de las variables en consola.
	   ¿Qué tipo de dato le asignarías a las variables? */

	var temperatura float32
	var humedad int
	var presion float32

	temperatura = 27
	humedad = 52
	presion = 1013

	fmt.Println("La temperatura en Mendoza es de ", temperatura, " grados centigrados", " - La humedad es de ", humedad, "%", " - La presion es de ", presion, " hPa")
	fmt.Printf("Temperatura: %f, Humedad: %v, Presion: %v\n", temperatura, humedad, presion)

	/* Un profesor de programación está corrigiendo los exámenes de sus estudiantes de la materia Programación I
			para poder brindarles las correspondientes devoluciones. Uno de los puntos del examen consiste en declarar
			distintas variables.
		Necesita ayuda para:
		Detectar cuáles de estas variables que declaró el alumno son correctas.
		Corregir las incorrectas.

	  var 1nombre string X
	  var apellido string V
	  var int edad X
	  1apellido := 6 X
	  var licencia_de_conducir = true X
	  var estatura de la persona int X
	  cantidadDeHijos := 2 V
	*/

	/* Un estudiante de programación intentó realizar declaraciones de variables con sus
			respectivos tipos en Go, pero tuvo varias dudas mientras lo hacía. A partir de esto,
			nos brindó su código y pidió la ayuda de un desarrollador experimentado que pueda:
		Verificar su código y realizar las correcciones necesarias.

	  var apellido string = "Gomez" V
	  var edad int = "35" X
	  boolean := "false"; X
	  var sueldo string = 45857.90 X
	  var nombre string = "Julián" V

	*/

	var edad int = 35
	boolean := false
	var sueldo float32 = 45857.90

	fmt.Print(edad, " ", boolean, " ", sueldo)

}
