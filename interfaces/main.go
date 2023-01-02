package main

import "fmt"

func main() {
	/*
		EJERCICIO 1
		Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para
		imprimir el detalle de los datos de cada uno de ellos/as, de la siguiente manera:

			Nombre: [Nombre del alumno]
			Apellido: [Apellido del alumno]
			DNI: [DNI del alumno]
			Fecha: [Fecha ingreso alumno]

		Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
		Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle
	*/

	alumno := Alumno{"Andrea", "Budassi", 30, "22/12/1992"}
	alumno.detalle()

	/*
		EJERCICIO 2
		Algunas tiendas ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
		La empresa tiene 3 tipos de productos: Pequeño, Mediano y Grande. (Se espera que sean muchos más)

		Y los costos adicionales son:
		Pequeño: solo tiene el costo del producto
		Mediano: el precio del producto + un 3% de mantenerlo en la tienda
		Grande: el precio del producto + un 6% de mantenerlo en la tienda, y adicional a eso $2500 de costo de envío.

		El porcentaje de mantenerlo en la tienda es en base al precio del producto.
		El costo de mantener el producto en stock en la tienda es un porcentaje del precio del producto.
		Se requiere una función factory que reciba el tipo de producto y el precio y retorne una interfaz Producto que tenga el método Precio.
		Se debe poder ejecutar el método Precio y que el método me devuelva el precio total en base al costo
		del producto y los adicionales en caso que los tenga
	*/

	fmt.Println(factory("Pequeño", 100).Costo())
}

// EJERCICIO 1
type Alumno struct {
	Nombre   string
	Apellido string
	DNI      int
	Fecha    string
}

func (alumno Alumno) detalle() {
	fmt.Printf("Nombre: %v\nApellido: %v\nDNI: %d\nFecha: %v\n", alumno.Nombre, alumno.Apellido, alumno.DNI, alumno.Fecha)
}

// EJERCICIO 2
// INTERFACE PRODUCTO
type Producto interface {
	Costo() float64
}

// ESTRUCTURA DE LOS PRODUCTOS
type Grande struct {
	Precio float64
}
type Pequeño struct {
	Precio float64
}
type Mediano struct {
	Precio float64
}

// METODOS DE LOS PRODUCTOS/COMPOSICION:RECEPTOR + NOMBRE DE FUNCION + TIPO DE DATO
func (p Mediano) Costo() float64 {
	return p.Precio * 1.03
}

func (p Pequeño) Costo() float64 {
	return p.Precio + 20
}

func (p Grande) Costo() float64 {
	return 2500 + p.Precio*1.06
}

// FUNCION QUE RECIBE EL TIPO DE PRODUCTO Y EL PRECIO Y RETORNA LA INTERFAZ DE PRODUCTO SELECCIONADO
func factory(producto string, precio float64) Producto {

	switch producto {
	case "Grande":
		return Grande{precio}
	case "Pequeño":
		return Pequeño{precio}
	case "Mediano":
		return Mediano{precio}
	default:
		return nil
	}

}
