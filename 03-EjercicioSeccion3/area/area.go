package area

import "fmt"

func AreaT(base float32, altura float32) {
	var resultado float32 = (base * altura) / 2

	fmt.Printf("La base es: %.2f\n", base)
	fmt.Printf("la altura es %.2f\n", altura)
	fmt.Printf("El área es: %.2f\n", resultado)
}
