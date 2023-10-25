package main

/*
Aquí se define una estructura llamada "Estudiante" que representa a un
estudiante. Esta estructura tiene campos para el nombre del estudiante, su
grupo, la cantidad de preguntas correctas en general y la cantidad de preguntas
de prueba correctas. Esto es útil para llevar un registro del rendimiento del
estudiante.
*/
type Estudiante struct {
	Nombre                   string
	Grupo                    int
	preguntasCorrectas       int
	preguntasPruebaCorrectas int
}

/*
Estas son funciones de método asociadas a la estructura "Estudiante". Cada
una de estas funciones se utiliza para actualizar las estadísticas del
estudiante. Por ejemplo, "ResponderBien" aumenta la cantidad de preguntas
correctas en general cuando un estudiante responde una pregunta
correctamente.
*/
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
