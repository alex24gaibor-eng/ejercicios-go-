package main
 
import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"unicode"
)
 
func main() {
	reader := bufio.NewReader(os.Stdin)
 

	fmt.Print("Ingrese una frase o texto: ")
	texto, _ := reader.ReadString('\n')
	texto = strings.TrimSpace(texto)
 
	palabras := strings.FieldsFunc(strings.ToLower(texto), func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})
 
	if len(palabras) == 0 {
		fmt.Println("No se ingresaron palabras.")
		return
	}

	conteo := make(map[string]int)
	for _, palabra := range palabras {
		conteo[palabra]++
	}
 
	claves := make([]string, 0, len(conteo))
	for palabra := range conteo {
		claves = append(claves, palabra)
	}
	sort.Strings(claves)
 
	fmt.Println("\nConteo de palabras:")
	for _, palabra := range claves {
		fmt.Printf("  %-15s -> %d\n", palabra, conteo[palabra])
	}
 
	masRepetida := ""
	maxVeces := 0
	for _, palabra := range claves { 
		if conteo[palabra] > maxVeces {
			maxVeces = conteo[palabra]
			masRepetida = palabra
		}
	}
 
	fmt.Printf("\nPalabra más repetida: \"%s\" (%d veces)\n", masRepetida, maxVeces)
}