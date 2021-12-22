package main

import "fmt"

type Usuario struct {
	Nombre    string
	Apellido  string
	Correo    string
	Productos []*Producto
}

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

func NuevoProducto(nombre string, precio float64) *Producto {
	var p = new(Producto)
	p.Nombre = nombre
	p.Precio = precio
	return p
}

func AgregarProducto(p *Producto, u *Usuario, cantidad int) {
	p.Cantidad = cantidad
	u.Productos = append(u.Productos, p)
}

func BorrarProducto(u *Usuario) {
	u.Productos = nil
}

func main() {
	var user Usuario
	user.Nombre = "Rodrigo"
	user.Apellido = "Souza"
	user.Correo = "Correo"

	p1 := NuevoProducto("prod1", 125.0)
	AgregarProducto(p1, &user, 20)
	fmt.Println(user)

	p2 := NuevoProducto("prod2", 456.70)
	AgregarProducto(p2, &user, 130)
	fmt.Println(user)

	BorrarProducto(&user)
	fmt.Println(user)
}
