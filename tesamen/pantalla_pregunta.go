package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/fatih/color"
	"github.com/nsf/termbox-go"
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
}

// Implementación del método renderizarPregunta para PPregunta.
func (p *PPregunta) renderizarPregunta() {
	err := termbox.Init() // Inicializa la biblioteca Termbox para manejar eventos.
	if err != nil {
		fmt.Println("Oh no, ocurrió un error al inicializar los controladores:", err)
		os.Exit(1)
	}
	defer termbox.Close()

	canalEventos := make(chan termbox.Event) // Crea un canal de eventos para manejar eventos de Termbox en una goroutine.
	go func() {
		for {
			canalEventos <- termbox.PollEvent()
		}
	}()

	p.Cabezero()                         // Muestra el encabezado de la pregunta.
	p.cuerpo()                           // Muestra el cuerpo de la pregunta.
	fmt.Println("\n\n", p.instrucciones) // Muestra las instrucciones de la pregunta.

	p.responder() // Llama al método responder para evaluar la respuesta del estudiante.

	for {
		select {
		case ev := <-canalEventos:
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Ch == 68 || ev.Key == termbox.KeyArrowLeft:
					anterior := p.ObtenerAnterior()
					if anterior != nil {
						termbox.Interrupt()
						termbox.Close()
						mostrarPantalla(anterior) // Navega a la pantalla anterior.
					}
				case ev.Ch == 67 || ev.Key == termbox.KeyArrowRight:
					siguiente := p.ObtenerSiguiente()
					if siguiente != nil {
						termbox.Interrupt()
						termbox.Close()
						mostrarPantalla(siguiente) // Navega a la pantalla siguiente.
					}

				case ev.Ch == 65 || ev.Key == termbox.KeyArrowUp:
					if p.respuestaSeleccionada-1 != -1 {
						p.respuestaSeleccionada-- // Mueve la selección de respuesta hacia arriba.
					}

				case ev.Ch == 66 || ev.Key == termbox.KeyArrowDown:
					if p.respuestaSeleccionada+1 != len(p.Respuestas) {
						p.respuestaSeleccionada++ // Mueve la selección de respuesta hacia abajo.
					}

				case ev.Key == termbox.KeyEnter:
					p.respuestaElegida = p.respuestaSeleccionada // Registra la respuesta elegida por el estudiante.
					p.responder()                                // Llama al método responder para evaluar la respuesta.
				}

				if runtime.GOOS == "windows" {
					termbox.Interrupt()
					termbox.Close()
				}
				p.renderizarPregunta() // Renderiza nuevamente la pregunta.
			}
		}
	}
}
