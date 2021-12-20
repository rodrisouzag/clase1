/*Varias tiendas de ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
Las empresas tienen 3 tipos de productos:
Pequeño, Mediano y Grande. (Se espera que sean muchos más)
Existen costos adicionales por mantener el producto en el almacén de la tienda, y costos de envío.

Sus costos adicionales son:
Pequeño: El costo del producto (sin costo adicional)
Mediano: El costo del producto + un 3% por mantenerlo en existencia en el almacén de la tienda.
Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.

Requerimientos:
Crear una estructura “tienda” que guarde una lista de productos.
Crear una estructura “producto” que guarde el tipo de producto, nombre y precio
Crear una interface “Producto” que tenga el método “CalcularCosto”
Crear una interface “Ecommerce” que tenga los métodos “Total” y “Agregar”.
Se requiere una función “nuevoProducto” que reciba el tipo de producto, su nombre y precio y devuelva un Producto.
Se requiere una función “nuevaTienda” que devuelva un Ecommerce.
Interface Producto:
El método “CalcularCosto” debe calcular el costo adicional según el tipo de producto.
Interface Ecommerce:
 - El método “Total” debe retornar el precio total en base al costo total de los productos y los adicionales si los hubiera.
 - El método “Agregar” debe recibir un producto y añadirlo a la lista de la tienda

*/

package main

const (
	pequeño = "pequeño"
	mediano = "mediano"
	grande  = "grande"
)

type Producto struct {
	Nombre string
	Precio float64
	Tipo   string
}

func CalcularCosto(p Producto) float64 {
	total := 0.0
	switch p.Tipo {
	case "pequeño":
		total = p.Precio
	case "mediano":
		total = p.Precio + p.Precio*0.03
	case "grande":
		total = p.Precio + p.Precio*0.06 + 2500.0
	}
	return total
}

func NuevoProducto(nombre string, precio float64, tipo string) Producto {
	return Producto{nombre, precio, tipo}
}

func NuevaTienda() Tienda {
	var t Tienda
	return t
}

type Tienda struct {
	Productos []Producto
}

func (t Tienda) Total() float64 {
	total := 0.0
	for _, p := range t.Productos {
		total += CalcularCosto(p)
	}
	return total
}

func (t *Tienda) Agregar(p Producto) {
	t.Productos = append(t.Productos, p)
}

type IProducto interface {
	CalcularCosto() float64
}

type Ecommerce interface {
	Total() float64
	Agregar(p Producto)
}

func main() {

}
