package greetings

import "fmt"

//Hello regresa saludos para el nombre proporcionado
//La funcion Hello comienza con mayusculas, las funciones con mayusculas pueden ser llamadas por funciones
//en otros paquetes
func Hello(name string) string {
	//regresa un saludo que incluye el name proporcionado
	//  := es un ajatajo para declrar a inicializar la variable
	message := fmt.Sprintf("Hola, %v. Bienvenido!", name)
	return message
}

//funcion de prueba con nombre comenzando en minuscula
func bye(name string) string {
	// otra forma de declarar variables
	var message string
	message = fmt.Sprintf("Adios %v.", name)
	return message
}
