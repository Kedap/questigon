package main

import (
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
	}
}

func (p *Pregunta) a_contenedor() *fyne.Container {
	contenedorPregunta := container.NewVBox(widget.NewLabel(p.Pregunta))
	opciones := widget.NewRadioGroup(p.Respuestas, func(r string) { p.Responder(r) })
	contenedorPregunta.Add(opciones)
	return contenedorPregunta
}
