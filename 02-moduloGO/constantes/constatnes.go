package constantes

import "fmt"

/*
las constantes se declaran con la palabra reservada "const" seguido del nombre de las constantes
pueden ser declaradas a nivel de paquete o a nivel de funcion, pero por convencion y por su uso
deberian de ser usadas a nivel de paquete y siempre con la primera letra Mayuscula
*/

const Pi float32 = 3.1415 //^constante en GO

//declaracion de varios valores constantes
const (
	Entero      = 100
	Binario     = 0b1010
	Octal       = 0o12
	Hexadecimal = 0xFF
)

//declaracion de constantes con el valor IOTA
const (
	Domingo = iota + 1 //^esto le asigna al domingo el valor iota=0 + 1 , entonces domingo tiene asignado el 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

//en el ejemplo anterior al usar IOTA, estamos diciendo que el valor de las constantes será consecutivo
//comenzando con el valor iniciar de la primer constante y aumentando siempre en 1

func UsoConstantes() {

	fmt.Println("el valor de Pi es ", Pi)
	fmt.Println(Entero, Binario, Octal, Hexadecimal)

	fmt.Println("Lunes es le valor de ", Lunes)
	fmt.Println("Miercoles es le valor de ", Miercoles)
	fmt.Println("Viernes es le valor de ", Viernes)
	fmt.Println("Por el valor de IOTA")
}
