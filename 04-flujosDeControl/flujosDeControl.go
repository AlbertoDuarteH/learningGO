package main

import "fmt"

func main() {

	//*operadores de comparacion
	//Se usan para comparar dos valores y devolver un valor booleano segun el
	//resultado de la comparacion
	comparacion()

	//*Operadores lógicos
	// Los operadores lógicos en Go son utilizados para evaluar expresiones lógicas
	//  y producir un resultado booleano (verdadero o falso). En Go, existen tres operadores
	// lógicos: AND lógico (&&), OR lógico (||) y NOT lógico (!).
	logicos()

}

func comparacion() {

	var (
		valor1 int = 5
		valor2 int = 1
	)

	fmt.Printf("valor1 = %d, valor2 = %d\n", valor1, valor2)
	//Operador de comparacion de igualdad ==
	fmt.Println("Operador de igualdad ==", valor1 == valor2)

	//Operador de comparacion de desigualdad !=
	fmt.Println("Operador de desigualdad ==", valor1 != valor2)

	//operador de comparacion mayor que >
	fmt.Println("Operador mayor que >", valor1 > valor2)

	//operdor de comparacion menor que <
	fmt.Println("Operador menor que <", valor1 < valor2)

	//operador de comparacion mayor o igual que >=
	fmt.Println("Operador de mayor o igual que >= ", valor1 >= valor2)

	//operador de comparacion menor o igual que <=
	fmt.Println("Operador de menor o igual que <= ", valor1 <= valor2)
}

func logicos() {
	var (
		x int = 5
		y int = 10
		z int = 15
	)

	fmt.Printf("\n>Valores iniciales\nx=%d, y=%d, z=%d\n", x, y, z)

	fmt.Println("x == y = ", x == y)
	fmt.Println("x != y = ", x != y)
	fmt.Println("true && false = ", true && false)
	fmt.Println("true || false = ", true || false)

	fmt.Println("Expresiones")
	resultado := ((x+y)*z)/(x*y) > z && x != y

	fmt.Printf("\n>Valores iniciales\nx=%d, y=%d, z=%d\n", x, y, z)
	fmt.Println("((x+y)*z)/(x*y) > z && x != y")
	fmt.Println("resultado = ", resultado)

}
