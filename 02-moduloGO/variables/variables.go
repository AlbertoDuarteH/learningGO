package variables

import "fmt"

var variableFuera1, variableFuera2 = "estoy afuera", "86"

func UsoVariables() {

	//1ra forma de Declaracion de variable
	var nombreVariable string           //Declaracion:  var + nombreVariable + tipo
	nombreVariable = "soy una variable" //asignacion de la variable
	fmt.Println(nombreVariable)         //uso de la variable (las variables en Go tienen que ser usadas)

	//2da forma de declaracion de variable
	var otraVariable string = "otra variable" //declaracion y asignación
	fmt.Println(otraVariable)

	//3ra forma de declaracion de varias variables
	var nombreVar, apellidoPVar, apellidoMVar string
	var edadVar int

	nombreVar = "Alberto"
	apellidoPVar = "Duarte"
	apellidoMVar = "Hurtado"
	edadVar = 42

	fmt.Println(nombreVar, " ", apellidoPVar, " ", apellidoMVar, ": ", edadVar)

	//4ta forma de declaracion de varias variables
	var (
		pokemon    string
		entrenador string
		posicion   int
	)

	pokemon = "Buelbasaur"
	entrenador = "Gary"
	posicion = 5

	fmt.Println(pokemon, entrenador, posicion)

	//5ta forma de declaracion varias varaibles y asignacion de valor
	var (
		elemento string = "Carbono"
		simboloQ string = "C"
		tipoE    string = "no metal"
		numeroA  int    = 6
	)

	fmt.Println(elemento, simboloQ, tipoE, numeroA)

	//6ta forma de declaracion varias varaibles y asignacion de valor sin especificar el tipo
	var (
		pelicula = "Matrix"
		director = "Hermanas Wachowski"
		anio     = 1999
	)
	fmt.Println(pelicula, director, anio)

	//7ma forma, declaracion y asignacion de variables por su posicion
	var cancion, artista, album, pista = "let it be", "the Beatles", "let it be", 5

	fmt.Println(cancion, artista, album, pista)

	//8va declaracion de variables fuera de la funcion (ver inicio del archivo)
	fmt.Println(variableFuera1, variableFuera2)

	/*
		En los sigueintes casos se usa un atajo, la direfenrecia entre usar el atajo y la palabra
		reservada "var" es que las variables que usen el atajo *SOLAMENTE* van dentro de funciones
		mientras que las declaradas por var, pueden ir afuera de funciones
	*/

	//9ra forma, declaracion con el atajo :=
	atajoVariable := "variable con atajo" //de esta manera no se usa var ni el tipado, GO asume el tipo
	fmt.Println(atajoVariable)

	//10ma forma, varias variables con atajo
	videoJuego, desarrollador, año := "Diablo IV", "BLIZZARD", 2023
	fmt.Println(videoJuego, desarrollador, año)
}
