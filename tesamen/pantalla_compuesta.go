package main

import (
	"fmt"
	"strings"
)

/*
PantallaCompuesta, la estructura PantallaSimple, la interfaz Pantalla, y las
implementaciones de métodos para PantallaSimple. La estructura
PantallaCompuesta utiliza una referencia a una PantallaSimple para mostrar
información adicional en la pantalla. Las implementaciones de métodos para
PantallaSimple son funciones vacías o de retorno que permiten a las pantallas
simples cumplir con la interfaz Pantalla.

Definición de la estructura PantallaCompuesta, que incorpora una PantallaSimple.
*/
type PantallaCompuesta struct {
	*PantallaSimple         // Una referencia a una pantalla simple.
	titulo           string // El título de la pantalla compuesta.
	preguntasTotales int    // La cantidad de preguntas totales en el examen.
}

// Función para mostrar un encabezado en la pantalla compuesta.
func (p *PantallaCompuesta) Cabezero() {
	fmt.Printf("\x1bc")
	fmt.Printf("\t\t\t--==[ %s ]==--\n\n%s\n\n\t\t\t%s\n\n", p.TituloExamen, p.Descripcion, strings.ToUpper(p.titulo))
}
