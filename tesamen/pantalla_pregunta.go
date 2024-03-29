package main

import (
	"fmt"

	"github.com/fatih/color"
)

/*
Se define la estructura PPregunta que representa una pantalla de pregunta en el
examen. La estructura contiene detalles sobre la pregunta, opciones de
respuesta, la respuesta correcta y el seguimiento del estudiante. También se
implementa un método responder que evalúa si la respuesta del estudiante es
correcta o incorrecta y actualiza las estadísticas del estudiante. El método
Cabezero se utiliza para mostrar el encabezado de la pantalla de la pregunta en
la terminal.

Definición de la estructura PPregunta que extiende la estructura PantallaCompuesta.
*/
type PPregunta struct {
	*PantallaCompuesta                // Referencia a una pantalla compuesta.
	Pregunta              string      // La pregunta que se mostrará.
	Respuestas            []string    // Un slice de respuestas posibles.
	RespuestaCorrecta     int         // El índice de la respuesta correcta.
	estudiante            *Estudiante // Referencia al objeto Estudiante asociado.
	respuestaSeleccionada int         // El índice de la respuesta seleccionada por el estudiante.
	respuestaElegida      int         // El índice de la respuesta que ha sido elegida (correcta o incorrecta).
	instrucciones         string      // Instrucciones relacionadas con la pregunta.
	respondioBien         bool        // Variable que indica si el estudiante respondió bien o mal.
}

// Definición de la interfaz Pregunta con un método responder.
type Pregunta interface {
	responder()
}

// Implementación del método responder para PPregunta.
func (p *PPregunta) responder() {
	// Comprueba si la respuesta elegida por el estudiante es correcta.
	if p.respuestaElegida == p.RespuestaCorrecta-1 && !p.respondioBien {
		p.estudiante.ResponderBien()       // Incrementa la cantidad de preguntas correctas del estudiante.
		p.respondioBien = !p.respondioBien // Cambia el estado de "respondioBien" para evitar múltiples incrementos.
	} else if p.respuestaElegida != p.RespuestaCorrecta-1 && p.respondioBien {
		p.estudiante.ResponderMal()        // Decrementa la cantidad de preguntas correctas del estudiante.
		p.respondioBien = !p.respondioBien // Cambia el estado de "respondioBien" para evitar múltiples decrementos.
	}
}

// Implementación del método Cabezero para PPregunta.
func (p *PPregunta) Cabezero() {
	fmt.Printf("\x1bc") // Limpia la pantalla.
	fmt.Printf("\t\t\t--==[ %s ]==--\n\n%s\n\n\t\t\t", p.TituloExamen, p.Descripcion)
	estiloPregunta := color.New(color.FgWhite).Add(color.Italic).Add(color.Bold) // Configura el estilo del título de la pregunta.
	estiloPregunta.Println(p.titulo)                                             // Muestra el título de la pregunta con el estilo configurado.
}

// Implementación del método cuerpo para PPregunta.
func (p *PPregunta) cuerpo() {
	estiloRespuesta := color.New(color.FgWhite).Add(color.BgMagenta).Add(color.Bold) // Configura el estilo de las respuestas.

	// Itera a través de las opciones de respuesta.
	for i, e := range p.Respuestas {
		var respuesta string
		if i == p.respuestaElegida {
			respuesta = "[*] " + e // Marca la respuesta elegida con "[*]".
		} else {
			respuesta = "[ ] " + e // Muestra las otras respuestas con "[ ]".
		}
		if i == p.respuestaSeleccionada {
			estiloRespuesta.Println(respuesta) // Muestra la respuesta seleccionada con el estilo configurado.
		} else {
			fmt.Println(respuesta) // Muestra las demás respuestas.
		}
	}
	p.responder()
	fmt.Println("\n\n", p.instrucciones) // Muestra las instrucciones de la pregunta.
}
func (p *PPregunta) TclDerecha(c *Controlador) Pantalla {
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
func (p *PPregunta) TclIzquierda(c *Controlador) {
	anterior := p.ObtenerAnterior()
	if anterior != nil {
		mostrarPantalla(anterior) // Navega a la pantalla anterior.
		c.IntercambiarPant(anterior)
	}
}
func (p *PPregunta) TclAbajo() {
	if p.respuestaSeleccionada+1 != len(p.Respuestas) {
		p.respuestaSeleccionada++ // Mueve la selección de respuesta hacia abajo.
		mostrarPantalla(p)
	}
}
func (p *PPregunta) TclArriba() {
	if p.respuestaSeleccionada-1 != -1 {
		p.respuestaSeleccionada-- // Mueve la selección de respuesta hacia arriba.
		mostrarPantalla(p)
	}
}
func (p *PPregunta) TclEnter() {
	p.respuestaElegida = p.respuestaSeleccionada // Registra la respuesta elegida por el estudiante.
	p.responder()                                // Llama al método responder para evaluar la respuesta.
	mostrarPantalla(p)
}
