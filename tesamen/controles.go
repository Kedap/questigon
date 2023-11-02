package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

type Controlador struct {
	pantallaSubscriptora Pantalla
	cancelar             chan struct{}
	canalEventos         chan termbox.Event
}

func (c *Controlador) EliminarPant() {
	c.pantallaSubscriptora = nil
}
func (c *Controlador) IntercambiarPant(p Pantalla) {
	c.EliminarPant()
	c.pantallaSubscriptora = p
}

// TODO: Cambiar el nombre del metodo
func (c *Controlador) ejecutar() {
	err := termbox.Init() // Inicializa la biblioteca Termbox para manejar eventos.
	if err != nil {
		fmt.Println("Oh no, ocurrió un error al inicializar los controladores:", err)
		time.Sleep(time.Second)
		os.Exit(1)
	}
	c.canalEventos = make(chan termbox.Event) // Crea un canal de eventos para manejar eventos de Termbox en una goroutine.
	go func() {
		for {
			select {
			case <-c.cancelar:
				termbox.Close()
				close(c.canalEventos)
				return
			case c.canalEventos <- termbox.PollEvent():
			}
		}
	}()

	// TODO: Ver que pedo con estos metodos

	// p.Cabezero()                         // Muestra el encabezado de la pregunta.
	// p.cuerpo()                           // Muestra el cuerpo de la pregunta.
	// fmt.Println("\n\n", p.instrucciones) // Muestra las instrucciones de la pregunta.
	// p.responder() // Llama al método responder para evaluar la respuesta del estudiante.

}
func (c *Controlador) Escuchar() {
	for ev := range c.canalEventos {
		if ev.Type == termbox.EventKey {
			switch {
			case ev.Ch == 68 || ev.Key == termbox.KeyArrowLeft:
				c.pantallaSubscriptora.TclIzquierda(c)
			case ev.Ch == 67 || ev.Key == termbox.KeyArrowRight:
				c.pantallaSubscriptora.TclDerecha(c)
			case ev.Ch == 65 || ev.Key == termbox.KeyArrowUp:
				c.pantallaSubscriptora.TclArriba()
			case ev.Ch == 66 || ev.Key == termbox.KeyArrowDown:
				c.pantallaSubscriptora.TclAbajo()
			case ev.Key == termbox.KeyEnter:
				c.pantallaSubscriptora.TclEnter()
			}
		}
	}
}
func (c *Controlador) Parar() {
	c.cancelar <- struct{}{}
	defer time.Sleep(500 * time.Microsecond)
}
