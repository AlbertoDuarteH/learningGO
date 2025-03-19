package main

import (
	"fmt"

	"ejercicio.com/area"
	"ejercicio.com/perimetro"
)

func main() {

	var (
		catetoA float32
		catetoO float32
	)
	fmt.Println("Programa que calcula el perimetro y el area de un triangulo R")
	fmt.Println("El programa calcula la hipotenusa")
	fmt.Println("Proporcione el valor del lado pequeño")
	fmt.Scanln(&catetoA)
	fmt.Println("Proporcione el valor del lado grade")
	fmt.Scanln(&catetoO)

	perimetro.PerimetroTR(catetoA, catetoO)
	area.AreaT(catetoA, catetoO)
}
