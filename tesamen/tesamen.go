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

/*
Los binarios de Go pueden generar falsos positivos en los antivirus debido a la
forma en que Go compila y empaqueta sus aplicaciones. Puede consultar las
razones en los enlaces que explican por qué los binarios de Go pueden ser
marcados como falsos positivos por los antivirus

https://go.dev/doc/faq#virus
https://groups.google.com/g/golang-nuts/c/Au1FbtTZzbk
https://github.com/golang/go/issues/44323
https://github.com/golang/go/issues/55042
*/
