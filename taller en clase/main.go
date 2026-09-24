package main

import (
	"fmt"
)


func main() {
	var opcion string

	for {
		fmt.Println("\n--- MENÚ PRINCIPAL ---")
		fmt.Println("1. Calificaciones de estudiantes")
		fmt.Println("2. Suma de 1 hasta 'n'")
		fmt.Println("3. Convertir Celsius a Fahrenheit")
		fmt.Println("4. Convertir Fahrenheit a Celsius")
		fmt.Println("0 o 'salir' para terminar")
		fmt.Print("Elige una opción: ")
		fmt.Scan(&opcion)

		if opcion == "0" || opcion == "salir" {
			fmt.Println("¡Hasta luego!")
			break
		}

		switch opcion {
		case "1":
			opcion1()
		case "2":
			opcion2()
		case "3":
			opcion3()
		case "4":
			opcion4()
		default:
			fmt.Println("Opción no válida. Intenta de nuevo.")
		}
	}
}


func opcion1() {
	var n int
	fmt.Print("Ingresa la cantidad de estudiantes: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("La cantidad debe ser mayor a 0.")
		return
	}

	var suma float64
	for i := 1; i <= n; i++ {
		var nota float64
		fmt.Printf("Ingresa la nota del estudiante %d (0 a 100): ", i)
		fmt.Scan(&nota)
		suma += nota
	}
3
	promedio := averageGrade(suma, n)
	fmt.Printf("\nEl promedio del curso es: %.2f\n", promedio)

	
	if promedio >= 70 {
		fmt.Println("Estado: APROBADO")
	} else {
		fmt.Println("Estado: REPROBADO")
	}

	
	switch {
	case promedio >= 90 && promedio <= 100:
		fmt.Println("Rendimiento: Excellent performance")
	case promedio >= 80 && promedio < 90:
		fmt.Println("Rendimiento: Good performance")
	case promedio >= 70 && promedio < 80:
		fmt.Println("Rendimiento: Satisfactory performance")
	default:
		fmt.Println("Rendimiento: Needs improvement")
	}
}


func averageGrade(suma float64, cantidad int) float64 {
	return suma / float64(cantidad)
}


func opcion2() {
	var n int
	fmt.Print("Ingresa un número (n): ")
	fmt.Scan(&n)

	suma := 0
	for i := 1; i <= n; i++ {
		suma += i
	}
	fmt.Printf("La suma del 1 al %d es: %d\n", n, suma)
}


func opcion3() {
	var celsius float64
	fmt.Print("Ingresa la temperatura en Celsius: ")
	fmt.Scan(&celsius)

	fahrenheit := (celsius * 9 / 5) + 32
	fmt.Printf("%.2f°C equivalen a %.2f°F\n", celsius, fahrenheit)
}


func opcion4() {
	var fahrenheit float64
	fmt.Print("Ingresa la temperatura en Fahrenheit: ")
	fmt.Scan(&fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Printf("%.2f°F equivalen a %.2f°C\n", fahrenheit, celsius)
}