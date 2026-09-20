package main

import "fmt"

// Función para sumar dos números
func suma(a, b int) int {
	return a + b
}

// Función para saludar de forma general
func saludar() {
	fmt.Println("¡Hola! Qué bueno verte por aquí.")
}

// Función de bienvenida personalizada
func bienvenida(nombre string) {
	fmt.Printf("Bienvenido/a, %s.\n", nombre)
}

// Ejemplo de función variádica (recibe múltiples argumentos)
func mostrarNum(numeros ...int) {
	fmt.Println("Números variádicos recibidos:")
	for _, n := range numeros {
		fmt.Println("-", n)
	}
}

func main() {
	var usr string

	fmt.Println("Ingresa tu nombre:")
	fmt.Scan(&usr)

	saludar()
	bienvenida(usr)

	fmt.Println("El resultado de la suma fija (4 + 5) es:", suma(4, 5))

	var a, b int
	fmt.Println("Ingresa el primer número:")
	fmt.Scan(&a)
	fmt.Println("Ingresa el segundo número:")
	fmt.Scan(&b)

	fmt.Println("El resultado de tu suma es:", suma(a, b))

	// Probando la función variádica
	fmt.Println("")
	mostrarNum(1, 2, 3, 4, 5)
}