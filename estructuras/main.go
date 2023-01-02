package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	   EJERCICIO 1
	   Crear un programa que cumpla los siguiente puntos:
	  1 Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
	  2 Tener un slice global de Product llamado Products instanciado con valores.
	  3 2 métodos asociados a la estructura Product: Save(), GetAll().
	  	El método Save() deberá tomar el slice de Products y añadir el producto desde el cual se llama al método.
	  	El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
	  4 Una función getById() al cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado.
	  5 Ejecutar al menos una vez cada método y función definido desde main().
	*/

	p1 := Product{ID: 1, Name: "Remera", Price: 500, Description: "Color rojo, talle M", Category: "Verano"}
	p1.Save(p1)
	p1.GetAll()
	fmt.Println("-------------")

	p2 := Product{ID: 2, Name: "Pantalon", Price: 1000, Description: "Color negro, talle M", Category: "Invierno"}
	p2.Save(p2)
	p2.GetAll()
	fmt.Println("-------------")

	fmt.Println(getById(1))
	fmt.Println(getById(2))

	/*
		EJERCICIO 2
		Una empresa necesita realizar una buena gestión de sus empleados, para esto realizaremos un pequeño programa nos ayudará
		a gestionar correctamente dichos empleados. Los objetivos son:
		1 Crear una estructura Person con los campos ID, Name, DateOfBirth.
		2 Crear una estructura Employee con los campos: ID, Position y una composicion con la estructura Person.
		3 Realizar el método a la estructura Employe que se llame PrintEmployee(), lo que hará es realizar la impresión de los campos de un empleado.
		4 Instanciar en la función main() tanto una Person como un Employee cargando sus respectivos campos y por último ejecutar el método PrintEmployee().
		Si logras realizar este pequeño programa pudiste ayudar a la empresa a solucionar la gestión de los empleados.
	*/
	p := Person{ID: 1, Name: "Cosme Fulanito", DateOfBirth: time.Date(1975, 12, 5, 0, 0, 0, 0, time.UTC)}
	e := Employee{ID: 2, Position: "Software Developer", Person: p}
	e.PrintEmployee()
}

// EJERCICIO 1
type Product struct {
	ID          int
	Name        string
	Price       int
	Description string
	Category    string
}

func (p Product) Save(Product) {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	for _, product := range Products {
		fmt.Println(product)
	}
}

func getById(ID int) Product {
	for _, prod := range Products {
		if prod.ID == ID {
			return prod
		}
	}
	return Product{}
}

var Products []Product

//EJERCICIO 2

type Person struct {
	ID          int
	Name        string
	DateOfBirth time.Time
}

type Employee struct {
	ID       int
	Position string
	Person
}

func (e *Employee) PrintEmployee() {
	fmt.Printf("ID: %d\n", e.ID)
	fmt.Printf("Nombre: %s\n", e.Name)
	fmt.Printf("Fecha de cumpleaños: %s\n", e.DateOfBirth.Format("02-01-2006"))
	fmt.Printf("Posicion: %s\n", e.Position)
}
