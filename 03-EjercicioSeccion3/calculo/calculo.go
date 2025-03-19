package calculo

import "math"

func Hipotenusa(catetoA float32, catetoO float32) float32 {
	var resultado float32

	sumaCat := math.Pow(float64(catetoA), 2) + math.Pow(float64(catetoO), 2)
	raizSum := math.Sqrt(sumaCat)

	resultado = float32(raizSum)
	return resultado
}
