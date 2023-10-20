package main

import "os"

func main() {
	var ruta string
	argumentos := os.Args
	if len(argumentos) == 1 {
		ruta = "examen.json"
	} else {
		ruta = argumentos[1]
	}
	examenMates := NuevoExamen(ruta)
	mostrarPantalla(examenMates.primeraPantalla)
}
