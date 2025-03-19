package tiposDatosB

import (
	"fmt"
	"strconv"
)

// tipos de datos
func TiposDatos() {
	//^ Números enteros con signo int
	// En Go los enteros vienen con int, int8, int32, int64
	var numeroInt int = 0 //^ por default el int es realtivo a la arquitectura del procesar del equipo (32 o 64)
	// En realidad parece que los difernetes int correspondes a los short, int, long  en java
	// por loq ue van en realcion al tamaño, por ejemplo int8 tiene un rango máximo de -128 a 127

	//^ Números enteros SIN signo uint
	// Similar a los int los uint pero sin signo, solo son valores positivos,
	// si tienen un valor negativo se produce un error
	var numeroUInt uint = 234
	// var numeroUInt2 uint = -1  //! Error, no puede asignarse negativo a un uint

	//^ Valores booleanos
	var flag bool = true

	//^ Datos Byte
	//es un tipo de dato que almacenan valores numericos del 0 al 255, realcionados con los
	// caracteres ASCII y otros datos que se almacenan en formato byte
	var myByte byte = 'a' //->  imprime 97
	//en este mismo contexto si tenemos la siguiente cadena de caracteres
	var cadenaS = "adios"
	fmt.Println("cadenaS[0] ->", cadenaS[0]) //-> imprime 97 que es el primer caracter en al cadena

	//^ Datos de tipo rune
	// Los datos de tipo rune almacenan numeros enteros positivos tal como int32, pero relacionados
	//al uso de caracteres UNICODE en GO
	var myRune rune = '☺'  //carita feliz -imprime  9786
	var myRune2 rune = '🚬' // cigarro (simbolo de fumar)
	fmt.Println("mi Rune", myRune)
	fmt.Println("mi Rune2", myRune2)

	fmt.Println(numeroInt)
	fmt.Println(numeroUInt)
	fmt.Println(flag)
	fmt.Println(myByte)
}

// conversiones de tipos
func ConversionDTipos() {
	//las conversiones de tipos en GO se deben realizar de manera explicita
	var integer8 int8 = 112
	var integer16 int16 = 3456

	//para realizar la suma, se tiene que hacer la conversion de tipo (en java seria Casteo)
	fmt.Println("conversion de int16(integer8) + integer16", int16(integer8)+integer16)

	//? Que pasa con operaciones o conversion de tipos cuando se quiere hacer una conversion de
	//? un tipo mayor a un tipo menor??????

	var integer32 int32 = 99999
	var intege8 int8 = 124

	fmt.Println("Conversion de intege8 + int8(integer32)", intege8, int8(integer32)) //lo acepta sin errores, pero el 99999 escede el valor de int8
	fmt.Print(" =", intege8+int8(integer32))
	//^la salida de la linea anterior es "27"
	// esto sucede porque hay un truncamiento debido aque la variable integer32 es mas grande solo se
	// toman los bits relacionados al tamaño del int 8 (o sea los primeros 8 bits, dejando aun lado el resto de los 32
	// por lo que la converiosn  99999 a int8 = -97  y 124+(-97) = 27

	//^ Conversion de tipos de string a int y viceversa
	var cadena string = "100"
	//En la conversion de string a int se tienen varios puntos a considerar
	//* El uso del paquete "strconv" que es un paquete estandar de go
	//* la funcion "strconv.Atoi(cadena)" devuelve 2 valores, el "int" y un "err", es por eso que
	//* en la asignacion se usa la notacion 'var entero, _ =' ya que como no nos interesa obtener el
	//* resultado del error se pone un guión bajo, si no se pone marca error,
	var entero, _ = strconv.Atoi(cadena)
	fmt.Println(entero)

	//para la conversion de int a string se usa la sigueinte funcion
	var cadenaEntero string = strconv.Itoa(100)
	fmt.Println(cadenaEntero)

}
