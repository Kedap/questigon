package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
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
func (p *TutorialPregunta) TclDerecha(c *Controlador) Pantalla {
	siguiente := p.ObtenerSiguiente()
	if siguiente != nil {
		if siguiente.NecesitaControlador() {
			mostrarPantalla(siguiente) // Navega a la pantalla siguiente.
			c.IntercambiarPant(siguiente)
		} else {
			return siguiente
		}
	}
	return nil
}
func (p *TutorialPregunta) TclIzquierda(c *Controlador) {
	anterior := p.ObtenerAnterior()
	if anterior != nil {
		mostrarPantalla(anterior) // Navega a la pantalla anterior.
		c.IntercambiarPant(anterior)
	}
}
func (p *TutorialPregunta) TclAbajo() {
	if p.respuestaSeleccionada+1 != len(p.Respuestas) {
		p.respuestaSeleccionada++ // Mueve la selección de respuesta hacia abajo.
		mostrarPantalla(p)
	}
}
func (p *TutorialPregunta) TclArriba() {
	if p.respuestaSeleccionada-1 != -1 {
		p.respuestaSeleccionada-- // Mueve la selección de respuesta hacia arriba.
		mostrarPantalla(p)
	}
}
func (p *TutorialPregunta) TclEnter() {
	p.respuestaElegida = p.respuestaSeleccionada // Registra la respuesta elegida por el estudiante.
	p.responder()                                // Llama al método responder para evaluar la respuesta.
	mostrarPantalla(p)
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
	siguiente := p.ObtenerSiguiente()
	controles := Controlador{
		pantallaSubscriptora: siguiente,
		cancelar:             make(chan struct{}),
	}
	controles.Lanzar()
	mostrarPantalla(siguiente)
	controles.Escuchar()
}
