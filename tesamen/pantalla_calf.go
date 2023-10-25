package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

/*
este código muestra la calificación del estudiante en un formato agradable en
la terminal, y cambia el estilo del texto en función de si el estudiante aprobó
o no el examen (según el promedio de respuestas correctas).

Definición de la estructura PantallaCalif que extiende la estructura PantallaCompuesta.
*/
type PantallaCalif struct {
	*PantallaCompuesta             // Referencia a una pantalla compuesta.
	estudiante         *Estudiante // Referencia al objeto Estudiante asociado.
}

// Implementación del método Cabezero para PantallaCalif.
func (p *PantallaCalif) Cabezero() {
	fmt.Printf("\x1bc")                                       // Limpia la pantalla.
	estiloTitulo := color.New(color.FgYellow).Add(color.Bold) // Configura el estilo del título.
	fmt.Printf("\t\t\t--==[ %s ]==--\n\n%s\n\n\t\t\t", p.TituloExamen, p.Descripcion)
	estiloTitulo.Print(strings.ToUpper(p.titulo)) // Muestra el título en mayúsculas.

	// Calcula el promedio de las respuestas correctas y lo muestra junto con otros detalles del estudiante.
	promedio := (p.estudiante.preguntasCorrectas * 100) / p.preguntasTotales
	fmt.Printf("\n\n\t%s\t%d\n\t\t%d/%d\t\t\t(%d%%)\n\n\t\t",
		p.estudiante.Nombre, p.estudiante.Grupo,
		p.estudiante.preguntasCorrectas,
		p.preguntasTotales, promedio)
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
	for {
		time.Sleep(time.Minute) // Espera indefinidamente.
	}
}
