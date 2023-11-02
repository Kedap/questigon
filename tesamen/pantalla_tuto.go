package main

import (
	"fmt"
)

/*
Este fragmento de código se centra en mostrar las instrucciones y el mensaje
final de un tutorial en la terminal.

Definición de la estructura PantallaTutorial que extiende la estructura PantallaCompuesta.
*/
type PantallaTutorial struct {
	*PantallaCompuesta        // Referencia a una pantalla compuesta.
	instrucciones      string // Las instrucciones del tutorial.
	msgFinal           string // Un mensaje final del tutorial.
}

// Implementación del método cuerpo para PantallaTutorial.
func (p *PantallaTutorial) cuerpo() {
	fmt.Println(p.instrucciones)
	fmt.Println(p.msgFinal)
}
func (p *PantallaTutorial) TclDerecha(c *Controlador) Pantalla {
	siguiente := p.ObtenerSiguiente()
	if siguiente != nil {
		if siguiente.NecesitaControlador() {
			mostrarPantalla(siguiente) // Navega a la pantalla siguiente.
			c.IntercambiarPant(siguiente)
		} else {
			return siguiente
		}
	}
	return nil
}
