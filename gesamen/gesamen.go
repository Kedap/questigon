package main

import (
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	estudiante := new(Estudiante)

	examen, err := NuevoExamen("examen.json", estudiante)
	if err != nil {
		panic(err)
	}
	cuestionario := Cuestionario{
		nombre:    examen.Nombre,
		app:       a,
		preguntas: &examen.Preguntas,
	}
	preguntas := NuevoCuestionario(cuestionario)
	inicio := NuevaVentanaInicio(a, preguntas, estudiante)
	inicio.ShowAndRun()
}
