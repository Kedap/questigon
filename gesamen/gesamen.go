package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	estudiante := new(Estudiante)
	w := a.NewWindow("Hello World")
	w.SetContent(widget.NewLabel("hola mundo!"))

	examen, err := NuevoExamen("examen.json", estudiante)
	if err != nil {
		panic(err)
	}
	cuestionario := Cuestionario{
		nombre:           examen.Nombre,
		app:              a,
		preguntas:        &examen.Preguntas,
		ventanaSiguiente: w,
	}
	preguntas := NuevoCuestionario(cuestionario)
	inicio := NuevaVentanaInicio(a, preguntas, estudiante)
	inicio.ShowAndRun()
}
