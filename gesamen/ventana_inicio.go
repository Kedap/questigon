package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

func NuevaVentanaIncio(a fyne.App, ventanaSiguiente fyne.Window, e *Estudiante) fyne.Window {
	ventanaInicio := a.NewWindow("Bienvenido: ingresa tus datos")
	nombre := widget.NewEntry()
	grupo := widget.NewEntry()
	instrucciones := widget.NewLabel("Ingresa tus datos")
	funcConfirmacion := func(c bool) {
		if c {
			ventanaInicio.Close()
			ventanaSiguiente.Show()
		}
	}
	confirmacion := dialog.NewConfirm("¿Estas seguro de continuar?", "¿Ingresaste bien tus datos?",
		funcConfirmacion, ventanaInicio)
	formulario := &widget.Form{
		Items: []*widget.FormItem{
			{Widget: instrucciones},
			{Text: "Nombre:", Widget: nombre},
			{Text: "Grupo", Widget: grupo},
		},
		OnSubmit: func() {
			nGrupo, err := strconv.Atoi(grupo.Text)
			if err != nil {
				dialog.NewInformation("Error",
					"El grupo debe de ser un numero entero chico",
					ventanaInicio)
				return
			}
			e.Nombre = nombre.Text
			e.Grupo = nGrupo
			confirmacion.Show()
		},
		OnCancel: func() {
			nombre.SetText("")
			grupo.SetText("")
		},
		SubmitText: "Continuar",
		CancelText: "Borrar",
	}
	ventanaInicio.SetContent(formulario)
	return ventanaInicio
}
