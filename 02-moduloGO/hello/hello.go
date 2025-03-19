package main

import (
	"fmt"

	"example.com/greetings"

	"example.com/variables"

	"example.com/constantes"

	"example.com/tiposDatosB"

	"example.com/paqueteFmt"

	"example.com/aritmetica"
)

func main() {
	message := greetings.Hello("Alberto")
	fmt.Println(message)
	fmt.Println("-----------------Variables------------------")
	variables.UsoVariables()
	fmt.Println("-----------------Constantes------------------")
	constantes.UsoConstantes()
	fmt.Println("-----------------Tipos de datos basicos------------------")
	tiposDatosB.TiposDatos()
	fmt.Println("-----------------Conversion de Tipos------------------")
	tiposDatosB.ConversionDTipos()
	fmt.Println("-----------------Uso de fmt------------------")
	paqueteFmt.UsoDeFmt()
	fmt.Println("-----------------Operaciones------------------")
	aritmetica.Operaciones()
	aritmetica.BitWise()
	aritmetica.UsoMath()
}
