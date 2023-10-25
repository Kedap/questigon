package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
La estructura PantallaConfirmacion parece representar una pantalla de
confirmación en un programa de examen. El método cuerpo se encarga de mostrar
instrucciones al usuario y permite al usuario confirmar su calificación. Si el
usuario escribe "ESTOY DE ACUERDO," se muestra la pantalla siguiente, de lo
contrario, se retrocede a la pantalla anterior.

Definición de la estructura PantallaConfirmacion que extiende la estructura PantallaCompuesta.
*/
type PantallaConfirmacion struct {
	*PantallaCompuesta // Referencia a una pantalla compuesta.
}

// Implementación del método cuerpo para PantallaConfirmacion.
func (p *PantallaConfirmacion) cuerpo() {
	// Muestra instrucciones para la confirmación.
	fmt.Println("Escribe \"ESTOY DE ACUERDO\" en MAYÚSCULAS para dar tu calificación.")
	fmt.Println("Importante: Hacer esta acción ya no te dejara retroceder a las preguntas")
	fmt.Println("Al igual de modificar las respuestas del examen.")
	fmt.Println("En el caso de escribir otra cosa, podrá modificar el las respuestas del examen")
	fmt.Print("\n> ")

	lector := bufio.NewScanner(os.Stdin) // Crea un lector de entrada.
	lector.Scan()                        // Lee la entrada del usuario.

	err := lector.Err()
	if err != nil {
		println("Ocurrió un error al leer la entrada :c\nPulse ENTER para salir del programa")
		fmt.Scanln()
		os.Exit(1)
	}
	// Comprueba si el usuario escribió "ESTOY DE ACUERDO" para continuar o retroceder.
	if lector.Text() == "ESTOY DE ACUERDO" {
		mostrarPantalla(p.ObtenerSiguiente()) // Muestra la siguiente pantalla.
	} else {
		mostrarPantalla(p.ObtenerAnterior()) // Muestra la pantalla anterior.
	}
}
