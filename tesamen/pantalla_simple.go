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

// Interfaz Pantalla que define métodos que deben ser implementados por las pantallas.
type Pantalla interface {
	Cabezero()                  // Muestra un encabezado en la pantalla.
	cuerpo()                    // Define el cuerpo de la pantalla.
	ObtenerSiguiente() Pantalla // Obtiene la siguiente pantalla.
	ObtenerAnterior() Pantalla  // Obtiene la pantalla anterior.
	renderizarPregunta()        // Renderiza una pregunta en la pantalla.
}

// Implementación de los métodos de la interfaz Pantalla para PantallaSimple.
func (p *PantallaSimple) cuerpo()                    {}
func (p *PantallaSimple) renderizarPregunta()        {}
func (p *PantallaSimple) ObtenerSiguiente() Pantalla { return p.PSiguiente }
func (p *PantallaSimple) ObtenerAnterior() Pantalla  { return p.PAnterior }
func (p *PantallaSimple) Cabezero() {
	fmt.Printf("\x1bc")
	fmt.Printf("\t\t\t--==[ %s ]==--\n\n\t\t%s\n\n", p.TituloExamen, p.Descripcion)
}

// Función mostrarPantalla que muestra una pantalla en la interfaz de la terminal.
func mostrarPantalla(p Pantalla) {
	// Se utiliza un switch para determinar el tipo de pantalla y ejecutar la acción correspondiente.
	switch p.(type) {
	case *PantallaIncio:
		// Si es una pantalla de inicio, muestra el encabezado y el cuerpo, luego pasa a la siguiente pantalla.
		p.Cabezero()
		p.cuerpo()
		mostrarPantalla(p.ObtenerSiguiente())
	case *TutorialPregunta:
		// Si es una pantalla de tutorial de pregunta, renderiza la pregunta.
		p.renderizarPregunta()
	case *PPregunta:
		// Si es una pantalla de pregunta, renderiza la pregunta.
		p.renderizarPregunta()
	case *PantallaConfirmacion:
		// Si es una pantalla de confirmación, muestra el encabezado y el cuerpo.
		p.Cabezero()
		p.cuerpo()
	case *TutorialCalificacion:
		// Si es una pantalla de tutorial de calificación, muestra el encabezado y pasa a la siguiente pantalla.
		p.Cabezero()
		mostrarPantalla(p.ObtenerSiguiente())
	case *PantallaCalif:
		// Si es una pantalla de calificación, muestra el encabezado y el cuerpo.
		p.Cabezero()
		p.cuerpo()
	default:
		// Si no se reconoce el tipo de pantalla, se inicia Termbox para manejar eventos.
		err := termbox.Init()
		if err != nil {
			fmt.Println("Oh no, ocurrió el error", err)
			os.Exit(1)
		}
		defer termbox.Close()

		// Se crea un canal de eventos para manejar eventos de Termbox en una goroutine.
		canalEventos := make(chan termbox.Event)
		go func() {
			for {
				canalEventos <- termbox.PollEvent()
			}
		}()

		// Se muestra el encabezado y el cuerpo de la pantalla y se espera eventos.
		p.Cabezero()
		p.cuerpo()

		// El programa se queda en un bucle infinito para manejar eventos.
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
