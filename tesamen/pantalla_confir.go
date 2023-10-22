package main

import (
	"bufio"
	"fmt"
	"os"
)

type PantallaConfirmacion struct {
	*PantallaCompuesta
}

func (p *PantallaConfirmacion) cuerpo() {
	fmt.Println("Escribe \"ESTOY DE ACUERDO\" para dar tu calificación.")
	fmt.Println("Importante: Hacer esta acción ya no te dejara retroceder a las preguntas")
	fmt.Println("Al igual de modificar las respuestas del examen.")
	fmt.Print("\n> ")
	lector := bufio.NewScanner(os.Stdin)
	lector.Scan()
	err := lector.Err()
	if err != nil {
		println("Ocurrió un error al leer la entrada :c\nPulse ENTER para salir del programa")
		fmt.Scanln()
		os.Exit(1)
	}
	if lector.Text() == "ESTOY DE ACUERDO" {
		mostrarPantalla(p.ObtenerSiguiente())
	} else {
		mostrarPantalla(p.ObtenerAnterior())
	}
}
