package main

import "fmt"

type PantallaSimple struct {
	TituloExamen string
	Descripcion  string
	PSiguiente   Pantalla
	PAnterior    Pantalla
}

// Interfaz Pantalla que define métodos que deben ser implementados por las pantallas.
type Pantalla interface {
	Cabezero()                  // Muestra un encabezado en la pantalla.
	cuerpo()                    // Define el cuerpo de la pantalla.
	ObtenerSiguiente() Pantalla // Obtiene la siguiente pantalla.
	ObtenerAnterior() Pantalla  // Obtiene la pantalla anterior.
	TclDerecha(c *Controlador) Pantalla
	TclIzquierda(c *Controlador)
	TclAbajo()
	TclArriba()
	TclEnter()
	NecesitaControlador() bool
}

// Implementación de los métodos de la interfaz Pantalla para PantallaSimple.
func (p *PantallaSimple) cuerpo()                            {}
func (p *PantallaSimple) ObtenerSiguiente() Pantalla         { return p.PSiguiente }
func (p *PantallaSimple) ObtenerAnterior() Pantalla          { return p.PAnterior }
func (p *PantallaSimple) TclDerecha(c *Controlador) Pantalla { return nil }
func (p *PantallaSimple) TclIzquierda(c *Controlador)        {}
func (p *PantallaSimple) TclAbajo()                          {}
func (p *PantallaSimple) TclArriba()                         {}
func (p *PantallaSimple) TclEnter()                          {}
func (p *PantallaSimple) NecesitaControlador() bool          { return true }
func (p *PantallaSimple) Cabezero() {
	fmt.Printf("\x1bc")
	fmt.Printf("\t\t\t--==[ %s ]==--\n\n\t\t%s\n\n", p.TituloExamen, p.Descripcion)
}

func mostrarPantalla(p Pantalla) {
	p.Cabezero()
	p.cuerpo()
}
