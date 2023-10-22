package main

import (
	"fmt"
	"os"

	"github.com/nsf/termbox-go"
)

type PantallaSimple struct {
	TituloExamen string
	Descripcion  string
	PSiguiente   Pantalla
	PAnterior    Pantalla
	//Conrtoles
}
type Pantalla interface {
	Cabezero()
	cuerpo()
	ObtenerSiguiente() Pantalla
	ObtenerAnterior() Pantalla
	renderizarPregunta()
}

func (p *PantallaSimple) cuerpo()                    {}
func (p *PantallaSimple) renderizarPregunta()        {}
func (p *PantallaSimple) ObtenerSiguiente() Pantalla { return p.PSiguiente }
func (p *PantallaSimple) ObtenerAnterior() Pantalla  { return p.PAnterior }
func (p *PantallaSimple) Cabezero() {
	fmt.Printf("\x1bc")
	fmt.Printf("\t\t\t--==[ %s ]==--\n\n\t\t%s\n\n", p.TituloExamen, p.Descripcion)
}

func mostrarPantalla(p Pantalla) {
	switch p.(type) {
	case *PantallaIncio:
		p.Cabezero()
		p.cuerpo()
		mostrarPantalla(p.ObtenerSiguiente())
	case *TutorialPregunta:
		p.renderizarPregunta()
	case *PPregunta:
		p.renderizarPregunta()
	case *PantallaConfirmacion:
		p.Cabezero()
		p.cuerpo()
	case *TutorialCalificacion:
		p.Cabezero()
		mostrarPantalla(p.ObtenerSiguiente())
	case *PantallaCalif:
		p.Cabezero()
		p.cuerpo()
	default:
		err := termbox.Init()
		if err != nil {
			fmt.Println("Oh no, ocurrió el error", err)
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
		for {
			select {
			case ev := <-canalEventos:
				if ev.Type == termbox.EventKey {
					switch {
					case ev.Ch == 68 || ev.Key == termbox.KeyArrowLeft:
						anterior := p.ObtenerAnterior()
						if anterior != nil {
							termbox.Close()
							mostrarPantalla(anterior)
						}
					case ev.Ch == 67 || ev.Key == termbox.KeyArrowRight:
						siguiente := p.ObtenerSiguiente()
						if siguiente != nil {
							termbox.Close()
							mostrarPantalla(siguiente)
						}
					}
				}
			}
		}
	}
}
