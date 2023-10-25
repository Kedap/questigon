package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/nsf/termbox-go"
)

// Definición de la estructura TutorialPregunta que extiende la estructura PPregunta.
type TutorialPregunta struct {
	*PPregunta // Referencia a una pregunta en forma de tutorial.
}

// Implementación del método responder para TutorialPregunta.
func (p *TutorialPregunta) responder() {
	// Comprueba si la respuesta elegida por el estudiante en el tutorial es correcta.
	if p.respuestaElegida == p.RespuestaCorrecta-1 && !p.respondioBien {
		p.estudiante.ResponderPruebaBien() // Incrementa la cantidad de preguntas de prueba respondidas correctamente por el estudiante.
		p.respondioBien = !p.respondioBien // Cambia el estado de "respondioBien" para evitar múltiples incrementos.
	} else if p.respuestaElegida != p.RespuestaCorrecta-1 && p.respondioBien {
		p.estudiante.ResponderPruebaMal()  // Decrementa la cantidad de preguntas de prueba respondidas correctamente por el estudiante.
		p.respondioBien = !p.respondioBien // Cambia el estado de "respondioBien" para evitar múltiples decrementos.
	}
}

func (p *TutorialPregunta) renderizarPregunta() {
	err := termbox.Init()
	if err != nil {
		fmt.Println("Oh no, ocurrió el error al inicializar los controladores:", err)
		os.Exit(1)
	}
	defer termbox.Close()
	canalEventos := make(chan termbox.Event)
	go func() {
		for {
			canalEventos <- termbox.PollEvent()
		}
	}()
	p.Cabezero()
	p.cuerpo()
	fmt.Println("\n\n", p.instrucciones)
	p.responder()
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
						mostrarPantalla(anterior)
					}
				case ev.Ch == 67 || ev.Key == termbox.KeyArrowRight:
					siguiente := p.ObtenerSiguiente()
					if siguiente != nil {
						termbox.Interrupt()
						termbox.Close()
						mostrarPantalla(siguiente)
					}
				case ev.Ch == 65 || ev.Key == termbox.KeyArrowUp:
					if p.respuestaSeleccionada-1 != -1 {
						p.respuestaSeleccionada--
					}
				case ev.Ch == 66 || ev.Key == termbox.KeyArrowDown:
					if p.respuestaSeleccionada+1 != len(p.Respuestas) {
						p.respuestaSeleccionada++
					}
				case ev.Key == termbox.KeyEnter:
					p.respuestaElegida = p.respuestaSeleccionada
					p.responder()
				}

				if runtime.GOOS == "windows" {
					termbox.Interrupt()
					termbox.Close()
				}
				p.renderizarPregunta()
			}
		}
	}
}

// Definición de la estructura TutorialCalificacion que extiende la estructura PantallaCalif.
type TutorialCalificacion struct {
	*PantallaCalif // Referencia a una pantalla de calificación.
}

// Implementación del método Cabezero para TutorialCalificacion.
func (p *TutorialCalificacion) Cabezero() {
	const N_PREGUNTAS_TUTORIAL = 2                            // Número de preguntas en el tutorial.
	fmt.Printf("\x1bc")                                       // Limpia la pantalla.
	estiloTitulo := color.New(color.FgYellow).Add(color.Bold) // Configura el estilo del título.
	fmt.Printf("\t\t\t--==[ %s ]==--\n\n%s\n\n\t\t\t", p.TituloExamen, p.Descripcion)
	estiloTitulo.Print(strings.ToUpper(p.titulo)) // Muestra el título en mayúsculas.

	// Calcula el promedio de las respuestas correctas en el tutorial y lo muestra junto con otros detalles del estudiante.
	promedio := (p.estudiante.preguntasPruebaCorrectas * 100) / N_PREGUNTAS_TUTORIAL
	fmt.Printf("\n\n\t%s\t%d\n\t\t%d/%d\t\t\t(%d%%)\n\n\t\t",
		p.estudiante.Nombre, p.estudiante.Grupo,
		p.estudiante.preguntasPruebaCorrectas,
		N_PREGUNTAS_TUTORIAL, promedio)

	estiloCalif := color.New() // Configura el estilo de la calificación.

	// Muestra "APROBADO!!!" en verde y en negrita si el promedio es igual o mayor al 70%.
	// Muestra "NO APROBADO" en rojo y en negrita si el promedio es menor al 70%.
	if promedio >= 70 {
		estiloCalif.Add(color.FgGreen).Add(color.Bold)
		estiloCalif.Println("APROBADO!!!")
	} else {
		estiloCalif.Add(color.FgRed).Add(color.Bold)
		estiloCalif.Println("NO APROBADO")
	}
	fmt.Println(`Bien! has resuelto el examen de prueba que NO TIENE VALOR,
presiona ENTER para comenzar con el VERADERO EXAMEN`)
	fmt.Scanln() // Espera a que el usuario presione Enter para continuar con el examen real.
}
