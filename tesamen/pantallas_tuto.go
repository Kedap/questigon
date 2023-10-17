package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/fatih/color"
	"github.com/nsf/termbox-go"
)

type TutorialPregunta struct {
	*PPregunta
}

func (p *TutorialPregunta) responder() {
	if p.respuestaElegida == p.RespuestaCorrecta-1 && !p.respondioBien {
		p.estudiante.ResponderPruebaBien()
		p.respondioBien = !p.respondioBien
	} else if p.respuestaElegida != p.RespuestaCorrecta-1 && p.respondioBien {
		p.estudiante.ResponderPruebaMal()
		p.respondioBien = !p.respondioBien
	}
}
func (p *TutorialPregunta) renderizarPregunta() {
	err := termbox.Init()
	if err != nil {
		fmt.Println("Oh no, ocurrio el error al inicializar los controladores:", err)
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

type TutorialCalificacion struct {
	*PantallaCalif
}

func (p *TutorialCalificacion) Cabezero() {
	const N_PREGUNTAS_TUTORIAL = 2
	fmt.Printf("\x1bc")
	estiloTitulo := color.New(color.FgYellow).Add(color.Bold)
	fmt.Printf("\t\t\t--==[ %s ]==--\n\n%s\n\n\t\t\t", p.TituloExamen, p.Descripcion)
	estiloTitulo.Print(strings.ToUpper(p.titulo))
	fmt.Printf("\n\n\t%s\t%d\n\t\t%d/%d\n\n\t\t", p.estudiante.Nombre, p.estudiante.Grupo, p.estudiante.preguntasPruebaCorrectas, N_PREGUNTAS_TUTORIAL)
	promedio := (p.estudiante.preguntasPruebaCorrectas * 100) / N_PREGUNTAS_TUTORIAL
	estiloCalif := color.New()
	if promedio >= 70 {
		estiloCalif.Add(color.FgGreen).Add(color.Bold)
		estiloCalif.Println("APROBADO!!!")
	} else {
		estiloCalif.Add(color.FgRed).Add(color.Bold)
		estiloCalif.Println("NO APROBADO")
	}
	fmt.Println("Presiona ENTER para comenzar con el verdadero examen!!!")
	fmt.Scanln()
}
