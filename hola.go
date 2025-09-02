package main

import (
	"fmt"
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
}
