package main

type Estudiante struct {
	Nombre                   string
	Grupo                    int
	preguntasCorrectas       int
	preguntasPruebaCorrectas int
}

func (e *Estudiante) ResponderBien() {
	e.preguntasCorrectas++
}
func (e *Estudiante) ResponderMal() {
	e.preguntasCorrectas--
}
func (e *Estudiante) ResponderPruebaBien() {
	e.preguntasPruebaCorrectas++
}
func (e *Estudiante) ResponderPruebaMal() {
	e.preguntasPruebaCorrectas--
}
