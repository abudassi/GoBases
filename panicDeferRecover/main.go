package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	/*
		EJERCICIO 1
		Un estudio contable necesita acceder a los datos de sus empleados para poder realizar distintas liquidaciones.
		Para ello, cuentan con todo el detalle necesario en un archivo .txt.
		Tendrás que desarrollar la funcionalidad para poder leer el archivo .txt que nos indica el cliente, sin embargo,
		no han pasado el archivo a leer por nuestro programa.
		Desarrollá el código necesario para leer los datos del archivo llamado “customers.txt” (recordá lo visto sobre el pkg “os”).
		Dado que no contamos con el archivo necesario, se obtendrá un error y, en tal caso, el programa deberá arrojar
		un panic al intentar leer un archivo que no existe, mostrando el mensaje “el archivo indicado no fue encontrado o está dañado”.
		Sin perjuicio de ello, deberá siempre imprimirse por consola “ejecución finalizada”.
	*/

	/*
		fmt.Println("-------------------")
		fmt.Println("INICIO EJERCICIO 1")

		_, err := os.Open("customers.txt")
		leerArchivo(err)

		fmt.Println("FIN EJERCICIO 1")
		fmt.Println("-------------------")
	*/

	/*
	   EJERCICIO 2
	   A continuación, vamos a crear un archivo “customers.txt” con información de los clientes del estudio.
	   Ahora que el archivo sí existe, el panic no debe ser lanzado.
	   Creamos el archivo “customers.txt” y le agregamos la información de los clientes.
	   Extendemos el código del punto uno para que podamos leer este archivo e imprimir los datos que contenga.
	   En el caso de no poder leerlo, se debe lanzar un “panic”.
	   Recordemos que siempre que termina la ejecución, independientemente del resultado, debemos tener un “defer”
	   que nos indique que la ejecución finalizó. También recordemos cerrar los archivos al finalizar su uso.
	*/
	/*
		fmt.Println("-------------------")
		fmt.Println("INICIO EJERCICIO 2")

		archivo, err := os.ReadFile("customers.txt")
		leerArchivo(err)
		fmt.Println(string(archivo))

		fmt.Println("FIN EJERCICIO 2")
		fmt.Println("-------------------")
	*/
	/*
		EJERCICIO 3
			El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder registrar datos de nuevos clientes. Los datos requeridos son:
				Legajo
				Nombre
				DNI
				Número de teléfono
				Domicilio

			Tarea 1: Antes de registrar a un cliente, debés verificar si el mismo ya existe.
			Para ello, necesitás leer los datos de un array. En caso de que esté repetido,
			debes manipular adecuadamente el error como hemos visto hasta aquí. Ese error deberá:
				1.- generar un panic;
				2.- lanzar por consola el mensaje: “Error: el cliente ya existe”, y continuar con la ejecución del programa normalmente.
			Tarea 2: Luego de intentar verificar si el cliente a registrar ya existe, desarrollá una función para validar
			que todos los datos a registrar de un cliente contienen un valor distinto de cero. Esta función debe retornar,
			al menos, dos valores. Uno de ellos tendrá que ser del tipo error para el caso de que se ingrese por
			parámetro algún valor cero (recordá los valores cero de cada tipo de dato, ej: 0, “”, nil).
			Tarea 3: Antes de finalizar la ejecución, incluso si surgen panics, se deberán imprimir por consola los
			siguientes mensajes: “Fin de la ejecución” y “Se detectaron varios errores en tiempo de ejecución”.
			Utilizá defer para cumplir con este requerimiento.

			Requerimientos generales:
			Utilizá recover para recuperar el valor de los panics que puedan surgir
			Recordá realizar las validaciones necesarias para cada retorno que pueda contener un valor error.
			Generá algún error, personalizandolo a tu gusto utilizando alguna de las funciones de Go (realiza también la validación pertinente para el caso de error retornado).
	*/

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("Se detectaron varios errores en tiempo de ejecución.")
		}
		fmt.Println("Ejecución finalizada.")
	}()

	filename := "customers.csv"

	// Abrir lista de usuarios
	OpenCustomers(filename)

	// Imprimir lista de usuarios
	PrintCustomers(customers)

	// Crear nuevos usuarios
	newCustomer := RegisterCustomer(120, "John Smith", "12345678", "12345678", "123 Main St")

	// Verificar si el usuario existe
	if CheckCustomerExists(newCustomer) {
		panic("Error: el cliente ya existe.")
	}

	// Validar usuario
	valid, err := ValidateCustomer(newCustomer)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("Customer data is valid. %v\n", valid)
		fmt.Println()
	}

	// Append Customer
	customers = append(customers, newCustomer)

	// Print Customers
	PrintCustomers(customers)

	// Save Customers
	SaveCustomers(customers)
	fmt.Println("Customers saved successfully.")

}

// EJERCICIO 1
func leerArchivo(err error) {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if err != nil {
		panic("el archivo indicado no fue encontrado o está dañado")
	}
	fmt.Println("Ejecucion finalizada")
}

// EJERCICIO 2
//USE MISMA FUNCION QUE EJERCICIO 1

// EJERCICIO 3
type Customer struct {
	ID      int
	Name    string
	DNI     string
	Phone   string
	Address string
}

var customers []*Customer

func RegisterCustomer(id int, name string, dni string, phone string, address string) *Customer {
	customer := &Customer{
		ID:      id,
		Name:    name,
		DNI:     dni,
		Phone:   phone,
		Address: address,
	}
	return customer
}

func CheckCustomerExists(customer *Customer) bool {
	for _, c := range customers {
		if c.ID == customer.ID {
			return true
		}
	}
	return false
}

func ValidateCustomer(customer *Customer) (bool, error) {
	if customer.ID == 0 {
		return false, errors.New("Error: ID must not be zero.")
	}
	if customer.Name == "" {
		return false, errors.New("Error: Name must not be empty.")
	}
	if customer.DNI == "" {
		return false, errors.New("Error: DNI must not be empty.")
	}
	if customer.Phone == "" {
		return false, errors.New("Error: Phone must not be empty.")
	}
	if customer.Address == "" {
		return false, errors.New("Error: Address must not be empty.")
	}
	return true, nil
}

func PrintCustomers(customer []*Customer) {
	fmt.Println("Customers:")
	for _, c := range customers {
		fmt.Printf("ID: %d, Name: %s, DNI: %s, Phone: %s, Address: %s\n", c.ID, c.Name, c.DNI, c.Phone, c.Address)
	}
	fmt.Println()
}

func OpenCustomers(filename string) (err error) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading records:", err)
		return
	}

	for _, record := range records {
		id, _ := strconv.Atoi(record[0])
		customer := &Customer{
			ID:      id,
			Name:    record[1],
			DNI:     record[2],
			Phone:   record[3],
			Address: record[4],
		}
		customers = append(customers, customer)
	}
	return
}

func SaveCustomers(customers []*Customer) (err error) {
	file, err := os.Create("customers.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}

	defer file.Close()

	for _, customer := range customers {
		_, err := fmt.Fprintf(file, "%d,%s,%s,%s,%s\n", customer.ID, customer.Name, customer.DNI, customer.Phone, customer.Address)
		if err != nil {
			return err
		}
	}
	return nil
}
