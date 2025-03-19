module ejercicio.com/ejercicios3

go 1.24.1

replace ejercicio.com/perimetro => ../perimetro

replace ejercicio.com/area => ../area

replace ejercicio.com/calculo => ../calculo

require (
	ejercicio.com/area v0.0.0-00010101000000-000000000000
	ejercicio.com/perimetro v0.0.0-00010101000000-000000000000
)

require ejercicio.com/calculo v1.1.0 // indirect
