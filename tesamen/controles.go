package main

import (
	"fmt"
	"os"
	"time"

	"github.com/nsf/termbox-go"
)

type Controlador struct {
	pantallaSubscriptora Pantalla
}

func (c *Controlador) EliminarPant() {
	c.pantallaSubscriptora = nil
}
func (c *Controlador) IntercambiarPant(p Pantalla) {
	c.EliminarPant()
	c.pantallaSubscriptora = p
}

func (c *Controlador) Lanzar() {
	err := termbox.Init() // Inicializa la biblioteca Termbox para manejar eventos.
	if err != nil {
		fmt.Println("Oh no, ocurrió un error al inicializar los controladores:", err)
		time.Sleep(time.Second)
		os.Exit(1)
	}
}

func (c *Controlador) Escuchar() {
	for {
		evento := termbox.PollEvent()
		if evento.Type == termbox.EventKey {
			switch {
			case evento.Ch == 68 || evento.Key == termbox.KeyArrowLeft:
				c.pantallaSubscriptora.TclIzquierda(c)
			case evento.Ch == 67 || evento.Key == termbox.KeyArrowRight:
				pantalla := c.pantallaSubscriptora.TclDerecha(c)
				if pantalla != nil {
					termbox.Close()
					defer mostrarPantalla(pantalla)
					return
				}
			case evento.Ch == 65 || evento.Key == termbox.KeyArrowUp:
				c.pantallaSubscriptora.TclArriba()
			case evento.Ch == 66 || evento.Key == termbox.KeyArrowDown:
				c.pantallaSubscriptora.TclAbajo()
			case evento.Key == termbox.KeyEnter:
				c.pantallaSubscriptora.TclEnter()
			}
		}
	}
}
