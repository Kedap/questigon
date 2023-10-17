package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/fatih/color"
	"github.com/nsf/termbox-go"
)

type PPregunta struct {
	*PantallaCompuesta
	Pregunta              string
	Respuestas            []string
	RespuestaCorrecta     int
	estudiante            *Estudiante
	respuestaSeleccionada int
	respuestaElegida      int
	instrucciones         string
	respondioBien         bool
}
type Pregunta interface {
	responder()
}

func (p *PPregunta) responder() {
	if p.respuestaElegida == p.RespuestaCorrecta-1 && !p.respondioBien {
		p.estudiante.ResponderBien()
		p.respondioBien = !p.respondioBien
	} else if p.respuestaElegida != p.RespuestaCorrecta-1 && p.respondioBien {
		p.estudiante.ResponderMal()
		p.respondioBien = !p.respondioBien
	}
}

func (p *PPregunta) Cabezero() {
	fmt.Printf("\x1bc")
	fmt.Printf("\t\t\t--==[ %s ]==--\n\n%s\n\n\t\t\t", p.TituloExamen, p.Descripcion)
	estiloPregunta := color.New(color.FgWhite).Add(color.Italic).Add(color.Bold)
	estiloPregunta.Println(p.titulo)
}
func (p *PPregunta) cuerpo() {
	estiloRespuesta := color.New(color.FgWhite).Add(color.BgMagenta).Add(color.Bold)
	for i, e := range p.Respuestas {
		var respuesta string
		if i == p.respuestaElegida {
			respuesta = "[*] " + e
		} else {
			respuesta = "[ ] " + e
		}
		if i == p.respuestaSeleccionada {
			estiloRespuesta.Println(respuesta)
		} else {
			fmt.Println(respuesta)
		}
	}
}
func (p *PPregunta) renderizarPregunta() {
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
				default:
					fmt.Println(ev.Ch)

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
