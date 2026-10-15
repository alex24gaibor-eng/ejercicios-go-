package main
 
import "fmt"
 

type Producto struct {
	Nombre string
	Precio float64
	Stock  int
}

func MostrarInfo(p Producto) {
	fmt.Printf("Producto: %s | Precio: $%.2f | Cantidad: %d\n", p.Nombre, p.Precio, p.Stock)
}

func ValorTotal(p Producto) float64 {
	return p.Precio * float64(p.Stock)
}

func TieneStock(p Producto) bool {
	return p.Stock > 0
}
 
func main() {
	productos := []Producto{
		{Nombre: "Laptop", Precio: 850.50, Stock: 5},
		{Nombre: "Mouse", Precio: 15.99, Stock: 40},
		{Nombre: "Teclado", Precio: 45.00, Stock: 0},
	}

	fmt.Println("=== Información de un producto ===")
	MostrarInfo(productos[0])

	fmt.Println("\n=== Inventario completo ===")
	for _, p := range productos {
		estado := "sin stock"
		if TieneStock(p) {
			estado = "disponible"
		}
		fmt.Printf("Nombre: %-10s | Precio: $%8.2f | Cantidad: %3d | Valor total: $%9.2f | Estado: %s\n",
			p.Nombre, p.Precio, p.Stock, ValorTotal(p), estado)
	}
}