package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NuevoResultados(a fyne.App, e *Estudiante, pregTotales int) fyne.Window {
	promedio := (e.RespuestasCorrectas * 100) / pregTotales
	ventanaResultados := a.NewWindow("Restultados")
	contenedor := container.NewVBox()
	lblNombre := widget.NewLabel(fmt.Sprint("Nombre:", e.Nombre))
	contenedor.Add(lblNombre)
	lblGrupo := widget.NewLabel(fmt.Sprint("Grupo:", e.Grupo))
	contenedor.Add(lblGrupo)
	lblCalif := widget.NewLabel(fmt.Sprintf("Calificacion: %d/%d",
		e.RespuestasCorrectas,
		pregTotales))
	contenedor.Add(lblCalif)
	if promedio >= 70 {
		contenedor.Add(widget.NewLabel("Aprobaste!!!!"))
	} else {
		contenedor.Add(widget.NewLabel("No aprobaste UnU"))
	}
	ventanaResultados.SetContent(contenedor)
	return ventanaResultados
}
