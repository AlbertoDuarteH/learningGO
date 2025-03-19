package aritmetica

import (
	"fmt"
	"math"
)

func Operaciones() {

	//^ Operadores Aritméticos
	// En GO se utilizan los sigueintes operadores Aritmeticos:
	var (
		a float32 = 10.0
		b float32 = 3.0
		c int     = 4
		d int     = 5
	)
	fmt.Printf("El valor de a = %g y el de b = %g \n", a, b)
	fmt.Printf("El valor de c = %d y el de d = %d \n", c, d)
	fmt.Printf("Suma: a + b = %g\n", a+b)
	fmt.Printf("Resta: a - b = %g\n", a-b)
	fmt.Printf("Multiplicacion: a * b = %g\n", a*b)
	fmt.Printf("Division: a / b = %g\n", a/b)
	fmt.Printf("Modulo: a %% b= %d\n", int(a)%int(b)) //solo se utiliza la parte entera
	c++
	fmt.Printf("incrementando c++ %d", c)
	d--
	fmt.Printf("decremento d-- %d", d)

	//operadores en asignacion
	c += 5 // es lo mismo que c=c+5
	c -= 5
	c *= 2
	c /= 2
	c %= 1
	fmt.Println("el valor de c es", c)
	fmt.Printf("El valor de c = %d y el de d = %d \n", c, d)
}

func BitWise() {

	/*
	 Un operador bitwise (u operador a nivel de bits) es un operador que realiza operaciones
	 directamente sobre los bits de los números enteros. Estos operadores trabajan con la
	 representación binaria de los datos, lo que los hace muy útiles en situaciones donde
	 se necesita manipular bits individualmente, como en programación de bajo nivel, optimización
	 de memoria o manejo de hardware.

	 En Go (y en muchos otros lenguajes de programación), los operadores bitwise más
	 comunes son:
	*/
	a := 12 // 1100 en binario
	b := 10 // 1010 en binario

	fmt.Println("AND", a&b)        // AND: 1000 (8 en decimal)
	fmt.Println("OR", a|b)         // OR: 1110 (14 en decimal)
	fmt.Println("XOR", a^b)        // XOR: 0110 (6 en decimal)
	a = 5                          // 0101 en binario
	fmt.Println("COMPLEMENTO", ^a) // 1010 (en un sistema de 4 bits, sería -6 en complemento a dos)
	a = 12
	fmt.Println("DESPLAZAMIENTO IZQ", a<<2) // Desplazamiento a la izquierda: 48
	fmt.Println("DESPLAZAMIENTO DER", a>>1) // Desplazamiento a la derecha: 6
}

func UsoMath() {
	var (
		a int = 3
		b int = 2
		c int = 10
		d int = 5
	)
	fmt.Println("potencia ", math.Pow(float64(a), float64(b))) //potencia, eleva un numoer a la potencia
	fmt.Println("raiz cuadrada", math.Sqrt(float64(c)))
	fmt.Println("raiz cúbica", math.Cbrt(float64(d)))
}
