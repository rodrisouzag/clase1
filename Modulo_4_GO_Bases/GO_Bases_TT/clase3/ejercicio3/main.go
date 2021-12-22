package main

import "fmt"

type Producto struct {
	Nombre   string
	Precio   float64
	Cantidad int
}

type Servicio struct {
	Nombre            string
	Precio            float64
	MinuntosTrabjados int
}

type Mantenimiento struct {
	Nombre string
	Precio float64
}

func SumarProductos(productos ...Producto) <-chan float64 {
	r := make(chan float64)

	go func() {
		total := 0.0
		for _, p := range productos {
			total += p.Precio * float64(p.Cantidad)
		}
		r <- total
	}()

	return r
}

func SumarServicios(servicios ...Servicio) <-chan float64 {
	r := make(chan float64)

	go func() {
		total := 0.0
		for _, s := range servicios {
			if s.MinuntosTrabjados < 30 {
				total += s.Precio * 30.0
			} else {
				total += s.Precio * float64(s.MinuntosTrabjados)
			}
		}
		r <- total
	}()

	return r
}

func SumarMantenimiento(mantenimientos ...Mantenimiento) <-chan float64 {
	r := make(chan float64)

	go func() {
		total := 0.0
		for _, m := range mantenimientos {
			total += m.Precio
		}
		r <- total
	}()
	return r
}

func main() {
	p1 := Producto{"prod1", 225.60, 15}
	p2 := Producto{"prod2", 456.25, 53}
	p3 := Producto{"prod3", 123.50, 10}

	s1 := Servicio{"serv1", 130.5, 20}
	s2 := Servicio{"serv2", 250.0, 50}

	m1 := Mantenimiento{"man1", 130}

	p, s, m := SumarProductos(p1, p2, p3), SumarServicios(s1, s2), SumarMantenimiento(m1)

	fmt.Printf("El resultado es %.2f", <-p+<-s+<-m)
}
