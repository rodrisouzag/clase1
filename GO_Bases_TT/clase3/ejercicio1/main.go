package main

import "fmt"

type Usuario struct {
	Nombre     string
	Apellido   string
	Edad       int
	Correo     string
	Contraseña string
}

func (e *Usuario) CambiarNombre(nombre string, apellido string) {
	e.Nombre = nombre
	e.Apellido = apellido
}

func (e *Usuario) CambiarEdad(edad int) {
	e.Edad = edad
}

func (e *Usuario) CambiarCorreo(correo string) {
	e.Correo = correo
}

func (e *Usuario) CambiarContraseña(contraseña string) {
	e.Contraseña = contraseña
}

func main() {
	user := Usuario{"Rodrigo", "Souza", 24, "rodrigo.souza@mercadolibre.com", "1234"}
	fmt.Println(user)

	user.CambiarNombre("Rodrigo", "Souza Garcia")
	fmt.Println(user)
}
