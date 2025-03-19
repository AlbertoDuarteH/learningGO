module example.com/hello

go 1.24.1

replace (
	example.com/greetings => ../greetings
	example.com/variables => ../variables
)

require (
	example.com/aritmetica v1.1.0
	example.com/constantes v1.1.0
	example.com/greetings v1.1.0
	example.com/paqueteFmt v1.1.0
	example.com/tiposDatosB v1.1.0
	example.com/variables v1.1.0
)

replace example.com/constantes => ../constantes

replace example.com/tiposDatosB => ../tiposDatosB

replace example.com/paqueteFmt => ../paqueteFmt

replace example.com/aritmetica => ../aritmetica
