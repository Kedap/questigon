package main

import "testing"

func TestResponder(t *testing.T) {
	pruebas := []struct {
		estudianteCorrectas         int
		estudianteCorrectasEsperado int
		respuestaElegida            int
		respuestaCorrecta           int
	}{
		{0, 1, 0, 1},
		{-11, -10, 3, 4},
		{9, 10, 1, 2},
		{-50, -49, 8, 9},
		{99, 99, 0, 2},
		{-50, -50, 1, 1},
	}

	for _, prueba := range pruebas {
		estudiante := Estudiante{
			Nombre:             "Test",
			preguntasCorrectas: prueba.estudianteCorrectas,
		}
		pregunta := PPregunta{
			Pregunta:          "Test",
			RespuestaCorrecta: prueba.respuestaCorrecta,
			estudiante:        &estudiante,
			respuestaElegida:  prueba.respuestaElegida,
		}
		pregunta.responder()
		if estudiante.preguntasCorrectas != prueba.estudianteCorrectasEsperado {
			t.Errorf("FALLO: obtenido: %d esperado: %d",
				estudiante.preguntasCorrectas, prueba.estudianteCorrectasEsperado)
		}
	}
}
