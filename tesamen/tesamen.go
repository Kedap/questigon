/*
Esto declara que el código se encuentra en el paquete principal. En Go, el
paquete principal se utiliza para construir un ejecutable. Sería similar a un
módulo principal en Python.
*/
package main

// En esta sección, se importan varias bibliotecas que se utilizarán
import "os"

// main es la función principal del programa.
func main() {
	var ruta string
	argumentos := os.Args
	// Verifica si se proporciona una ruta de archivo como argumento, de lo contrario, usa un valor predeterminado.
	if len(argumentos) == 1 {
		ruta = "examen.json"
	} else {
		ruta = argumentos[1]
	}
	// Crea una instancia de examen a partir del archivo JSON especificado.
	examenMates := NuevoExamen(ruta)
	// Muestra la primera pantalla del examen para comenzar.
	mostrarPantalla(examenMates.primeraPantalla)
}
