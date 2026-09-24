package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ConvertirMoneda(dolares float64, moneda string) {
	var resultado float64
	switch strings.ToUpper(moneda) {
	case "EUROS":
		resultado = dolares * 0.92
		fmt.Printf("%.2f Dólares equivalen a %.2f Euros\n", dolares, resultado)
	case "LB":
		resultado = dolares * 0.79
		fmt.Printf("%.2f Dólares equivalen a %.2f Libras Esterlinas\n", dolares, resultado)
	case "WON":
		resultado = dolares * 1350.0
		fmt.Printf("%.2f Dólares equivalen a %.2f Wones\n", dolares, resultado)
	case "BTC":
		resultado = dolares / 65000.0
		fmt.Printf("%.2f Dólares equivalen a %.6f BTC\n", dolares, resultado)
	default:
		fmt.Println("Moneda no permitida. Use: Euros, LB, Won o BTC")
	}
}

func ContarVocales(frase string) {
	frase = strings.ToLower(frase)
	a, e, i, o, u := 0, 0, 0, 0, 0

	for _, char := range frase {
		switch char {
		case 'a', 'á':
			a++
		case 'e', 'é':
			e++
		case 'i', 'í':
			i++
		case 'o', 'ó':
			o++
		case 'u', 'ú', 'ü':
			u++
		}
	}

	fmt.Println("Conteo de vocales:")
	fmt.Printf("A: %d\n", a)
	fmt.Printf("E: %d\n", e)
	fmt.Printf("I: %d\n", i)
	fmt.Printf("O: %d\n", o)
	fmt.Printf("U: %d\n", u)
}

func main() {
	var dolares float64
	var moneda string

	fmt.Println("--- CONVERSOR DE MONEDAS ---")
	fmt.Print("Ingrese la cantidad en dólares: ")
	fmt.Scan(&dolares)
	fmt.Print("Ingrese la moneda destino (Euros, LB, Won, BTC): ")
	fmt.Scan(&moneda)
	ConvertirMoneda(dolares, moneda)

	fmt.Println("\n--- CONTADOR DE VOCALES ---")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Escriba una frase: ")
	frase, _ := reader.ReadString('\n')
	frase = strings.TrimSpace(frase)
	ContarVocales(frase)
}