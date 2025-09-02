package main

import (
	"fmt"
	"math"
	"strconv"

	"rsc.io/quote"
)

// Declaracion de constantes
const Pi float32 = 3.14
const (
	x = 100
	y = 0b100
	z = 0o12
	w = 0xFF
)

const (
	Domingo = iota + 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	fmt.Println("Hola, Mundo")
	fmt.Println(quote.Go())
	fmt.Println("El valor de Pi es:", Pi)
	fmt.Println(x, y, z, w)
	fmt.Println(Domingo, Lunes, Martes, Miercoles, Jueves, Viernes, Sabado)

	//Declaracion de variables 1ERA FORMA
	//var firstName string = "Luis"
	//var age int = 23

	//Declaracion de las variables segunda forma

	var firstName, age = "Luis", 23

	//Declaracion de las variables tercera forma
	lastName, country, year := "Juarez", "Argentina", 2024

	//conversion de tipos

	var integer16 int16 = 1000
	var integer32 int32 = 50

	fmt.Println("Suma:", integer16+int16(integer32))

	s := "100"
	i, _ := strconv.Atoi(s)
	fmt.Println("Entero:", i)

	n := 43
	s = strconv.Itoa(n)

	fmt.Println("String:", s)
	fmt.Println("Nombre:", firstName)
	fmt.Println("Edad:", age)
	fmt.Println("Apellido:", lastName)
	fmt.Println("Pais:", country)
	fmt.Println("Año:", year)

	//package fmt

	fmt.Print("Ingresar un nombre: ")
	fmt.Scanln(&firstName)
	fmt.Print("Ingrese su edad: ")
	fmt.Scanln(&age)

	fmt.Printf("Hola, me llamo %s y tengo %d años.\n", firstName, age)

	//resolver problemas de area de un triangulo rectangulo

	var lado1, lado2 float64

	fmt.Print("Ingrese el lado 1:")
	fmt.Scanln(&lado1)
	fmt.Print("Ingrese el lado 2:")
	fmt.Scanln(&lado2)

	//process

	area := (lado1 * lado2) / 2
	hipotenusa := math.Sqrt(math.Pow(lado1, 2) + math.Pow(lado2, 2))
	perimetro := lado1 + lado2 + hipotenusa

	fmt.Printf("El area del triangulo es: %.2f\n", area)
	fmt.Printf("La hipotenusa del triángulo es: %.2f\n", hipotenusa)
	fmt.Printf("El perímetro del triángulo es: %.2f\n", perimetro)
}
