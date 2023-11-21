package main

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Pregunta struct {
	Pregunta          string
	Respuestas        []string
	RespuestaCorrecta int
	resuelta          bool
	estudiante        *Estudiante
}

func (p *Pregunta) Responder(respuesta string) {
	if respuesta == p.Respuestas[p.RespuestaCorrecta-1] && !p.resuelta {
		p.estudiante.RespuestasCorrectas++
		p.resuelta = true
	} else if respuesta != p.Respuestas[p.RespuestaCorrecta-1] && p.resuelta {
		p.estudiante.RespuestasCorrectas--
		p.resuelta = false
	} else if respuesta != p.Respuestas[p.RespuestaCorrecta-1] && !p.resuelta {
		if p.estudiante.RespuestasCorrectas != 0 {
			p.estudiante.RespuestasCorrectas--
			p.resuelta = false
		}
	}
}

// aContenedor method : Retorna un contenedor con el nombre de la pregunta
// opciones y una función que la resuelve `p.Responder(respuestaElegida)`
func (p *Pregunta) aContenedor(preguntas *[]Pregunta, id int) *fyne.Container {
	contenedorPregunta := container.NewVBox(widget.NewLabel(p.Pregunta))
	opciones := widget.NewRadioGroup(p.Respuestas, func(r string) {
		preguntas := *preguntas
		preguntaActual := preguntas[id]
		preguntaActual.Responder(r)
		log.Println(p.estudiante.RespuestasCorrectas, preguntaActual.resuelta)
	})
	contenedorPregunta.Add(opciones)
	return contenedorPregunta
}
