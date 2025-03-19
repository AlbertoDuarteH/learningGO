package perimetro

import (
	"fmt"

	"ejercicio.com/calculo"
)

/*
calculo del perimetro del triangulo rectangulo
*/
func PerimetroTR(catetoA float32, catetoO float32) {
	var resultado float32

	fmt.Printf("Cateto Adyacente del triángulo:%.2f\n", catetoA)
	fmt.Printf("Cateto Opuesto triángulo: %.2f\n", catetoO)
	hipotenusa := calculo.Hipotenusa(catetoA, catetoO)
	fmt.Printf("hipotenusa del triángulo: %.2f \n", hipotenusa)
	resultado = catetoA + catetoO + hipotenusa
	fmt.Printf("el perimetro es: %.2f\n", resultado)
}
