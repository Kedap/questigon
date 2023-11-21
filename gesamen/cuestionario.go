package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Cuestionario struct {
	nombre    string
	app       fyne.App
	preguntas *[]Pregunta
}

func NuevoCuestionario(c Cuestionario) fyne.Window {
	ventanaCuestionario := c.app.NewWindow(c.nombre)
	contenedorCuestionario := container.NewVBox()
	preguntas := c.preguntas
	var estudiante *Estudiante
	for i, p := range *preguntas {
		preguntaContenedor := p.aContenedor(preguntas, i)
		contenedorCuestionario.Add(preguntaContenedor)
		if i == len(*preguntas)-1 {
			estudiante = p.estudiante
		}
	}
	btnResponder := widget.NewButtonWithIcon(
		"Responder",
		theme.ConfirmIcon(),
		func() {
			ventanaCuestionario.Hide()
			resultados := NuevoResultados(c.app, estudiante, len(*preguntas))
			resultados.Show()
			ventanaCuestionario.Close()
		},
	)
	contenedorCuestionario.Add(btnResponder)
	cuestionarioScrolleable := container.NewVScroll(contenedorCuestionario)
	ventanaCuestionario.SetContent(cuestionarioScrolleable)
	return ventanaCuestionario
}
