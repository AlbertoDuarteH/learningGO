package paqueteFmt

import (
	"fmt"
	"math"
)

func UsoDeFmt() {

	/*
		* El paquete FMT es uno de los paquetes estandard mas usados en GO ya que porporciona
		entrdas y salidas basicas para el programa
	*/

	/*
		* Antes de continuar hay que saber algunas cosas sobre el FORMATO

		el paquete fmt implementa formatos de entradas y salidas con funciones análogas como el lenguaje C lo
		hace con printf y scanf pero de manera mas simple

		GO utiliza modificadores llamados *"VERBOS"* con estos verbos es como se le da formato
		https://pkg.go.dev/fmt#pkg-index -> documentation

		^ Verbos generales
		%v	Se utiliza para representar un valor por defecto cuando se imprimen Structs (estructuras)
			cuando se usa con un signo de mas (%+v) incluye el nombre del campo
		%#v	Es una sintaxis de representacion de valores en GO, cuando es un valor de punto flotante
			infinito imprime "±Inf" y, en caso de no ser un número, imprime "NaN"
		%T	Es una sintaxys de representacion en GO imprime el tipo del dato que se esta pasando
		%%	se utiliza para escribir el caracter %

		*El formato predeterminado para %v es:
		bool: %t
		int, int8 etc.: %d
		uint, uint8 etc.: %d, %#x si se imprime con %#v
		float32, complex64, etc.: %g
		string: %s
		chan: %p    de channel o canales
		pointer: %p

		* Para los objetos compuestos, los elementos se imprimen utilizando estas reglas,
		* de forma recursiva, dispuestas de la siguiente manera:
		struct: {campo0 campo1 ...}
		matriz, slice: [elem0 elem1 ...]
		mapas: map[clave1:valor1 clave2:valor2 ...]
		puntero a lo anterior: &{}, &[], &map[]
	*/
	var (
		cadena string  = "hola"
		numero int     = 86
		zero   float64 = 0.0
	)

	type Persona struct { // este es un tipo de dato Estructura se estudiará mas adelante
		Nombre   string
		Apellido string
		Edad     int8
	}

	p1 := Persona{Nombre: "Isaac", Apellido: "Molina", Edad: 26}
	fmt.Println("******Formato con verbos generales *******************")
	fmt.Printf("%v agente %v\n", cadena, numero)
	fmt.Printf("%v estructura persona con %%v\n", p1)
	fmt.Printf("%+v estructura persona con %%+v\n", p1)
	fmt.Printf("%#v este es un valor infinito de 1.0/0.0 con %%#v\n", (float64(1.0) / zero))
	fmt.Printf("%#v este es un valor infinito de -1.0/0.0 con %%#v\n", (float64(-1.0) / zero))
	fmt.Printf("%#v este no es un numero\n", math.NaN())
	fmt.Printf("el tipo de dato de la variable \"cadena\" es : %T\n", cadena)

	/*
		^ Verbos booleanos
		%t	imprime la palabra true o false dependeindo del valor de la variable
	*/
	fmt.Println("******Formato con verbos booleanos *******************")
	var myboolean bool = true
	fmt.Printf("el valor de myboolean es %t \n", myboolean)

	/*
		^ Verbos de numeros Enteros
		%b base 2
		%c el carácter representado por el punto de código Unicode correspondiente
		%d base 10
		%o base 8
		%O base 8 con prefijo 0o
		%q es un carácter literal entre comillas simples que se escapa de forma segura con la sintaxis Go.
		%x base 16, con letras minúsculas para af
		%X base 16, con letras mayúsculas para AF
		%U Formato Unicode: U+1234; igual que "U+%04X"
	*/

	/*
		^ Verbos de numeros de punto flotante y complejos
		%b notación científica sin decimales con exponente potencia de dos,
			de la misma manera que strconv.FormatFloat con el formato 'b',
			p. ej. -123456p-78
		%e notación científica, p. ej. -1,234456e+78
		%E notación científica, p. ej. -1,234456E+78
		%f punto decimal pero sin exponente, p. ej. 123,456
		%F sinónimo de %f
		%g %e para exponentes grandes, %f en caso contrario. La precisión se analiza más adelante.
		%G %E para exponentes grandes, %F en caso contrario
		%x notación hexadecimal (con exponente decimal de potencia dos), p. ej. -0x1.23abcp+20
		%X notación hexadecimal en mayúsculas, p. ej.: -0X1.23ABCP+20

		El exponente es siempre un entero decimal.
		Para formatos distintos de %b, el exponente tiene al menos dos dígitos.
	*/

	/*
		^ Verbos de cadenas y de porciones de bytes
		%s los bytes no interpretados de la cadena o porción
		%q una cadena entre comillas dobles escapada de forma segura con sintaxis Go
		%x base 16, minúsculas, dos caracteres por byte
		%X base 16, mayúsculas, dos caracteres por byte

	*/

	/*
		^ Verbos para slices
		%p dirección del elemento 0 en notación base 16, con 0x inicial
	*/

	/*
		^ Verbos para punteros
		Notación base 16 %p, con 0x inicial.
		Los verbos %b, %d, %o, %x y %X también funcionan con punteros,
		formateando el valor exactamente como si fuera un entero.
	*/

	//---------------------------------------------SCAN
	//El uso de fmt.Scan, es para poder introducir valores a la consola

	var (
		tuNombre string
		tuEdad   int8
	)

	fmt.Println("*******************Uso de fmt.Scan*********************")
	fmt.Println("proporcioname tu nombre")
	fmt.Scanln(&tuNombre)
	fmt.Println("proporcioname tu edad")
	fmt.Scanln(&tuEdad)

	fmt.Printf("ASi que tu eres %s y tienes %d años", tuNombre, tuEdad)

}
