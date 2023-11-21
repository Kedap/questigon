package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type Cuestionario struct {
	nombre           string
	app              fyne.App
	preguntas        *[]Pregunta
	ventanaSiguiente fyne.Window
}

func NuevoCuestionario(c Cuestionario) fyne.Window {
	ventanaCuestionario := c.app.NewWindow(c.nombre)
	contenedorCuestionario := container.NewVBox()
	preguntas := c.preguntas
	for i, p := range *preguntas {
		preguntaContenedor := p.aContenedor(preguntas, i)
		contenedorCuestionario.Add(preguntaContenedor)
	}
	btnResponder := widget.NewButtonWithIcon(
		"Responder",
		theme.ConfirmIcon(),
		func() {
			ventanaCuestionario.Hide()
			c.ventanaSiguiente.Show()
		},
	)
	contenedorCuestionario.Add(btnResponder)
	cuestionarioScrolleable := container.NewVScroll(contenedorCuestionario)
	ventanaCuestionario.SetContent(cuestionarioScrolleable)
	return ventanaCuestionario
}
