package main

import "testing"

type pruebasEstudiante struct {
	estadoInicial, estadoEsperado, iteraciones int
}

func TestResponderBien(t *testing.T) {
	pruebas := []pruebasEstudiante{
		{0, 1, 1},
		{2, 5, 3},
		{5, 5, 0},
		{-1, 0, 1},
	}

	for _, prueba := range pruebas {
		estudiante := Estudiante{
			Nombre:             "Test",
			preguntasCorrectas: prueba.estadoInicial,
		}
		for i := 0; i < prueba.iteraciones; i++ {
			estudiante.ResponderBien()
		}
		if estudiante.preguntasCorrectas != prueba.estadoEsperado {
			t.Errorf("FALLO: obtenido: %d esperado : %d",
				estudiante.preguntasCorrectas, prueba.estadoEsperado)
		}
	}
}

func TestResponderMal(t *testing.T) {
	pruebas := []pruebasEstudiante{
		{0, -1, 1},
		{2, -1, 3},
		{5, 5, 0},
		{-1, -2, 1},
	}

	for _, prueba := range pruebas {
		estudiante := Estudiante{
			Nombre:             "Test",
			preguntasCorrectas: prueba.estadoInicial,
		}
		for i := 0; i < prueba.iteraciones; i++ {
			estudiante.ResponderMal()
		}
		if estudiante.preguntasCorrectas != prueba.estadoEsperado {
			t.Errorf("FALLO: obtenido: %d esperado : %d",
				estudiante.preguntasCorrectas, prueba.estadoEsperado)
		}
	}
}

func TestResponderPruebaBien(t *testing.T) {
	pruebas := []pruebasEstudiante{
		{0, 1, 1},
		{2, 5, 3},
		{5, 5, 0},
		{-1, 0, 1},
	}

	for _, prueba := range pruebas {
		estudiante := Estudiante{
			Nombre:                   "Test",
			preguntasPruebaCorrectas: prueba.estadoInicial,
		}
		for i := 0; i < prueba.iteraciones; i++ {
			estudiante.ResponderPruebaBien()
		}
		if estudiante.preguntasPruebaCorrectas != prueba.estadoEsperado {
			t.Errorf("FALLO: obtenido: %d esperado : %d",
				estudiante.preguntasCorrectas, prueba.estadoEsperado)
		}
	}
}

func TestResponderPruebaMal(t *testing.T) {
	pruebas := []pruebasEstudiante{
		{0, -1, 1},
		{2, -1, 3},
		{5, 5, 0},
		{-1, -2, 1},
	}

	for _, prueba := range pruebas {
		estudiante := Estudiante{
			Nombre:                   "Test",
			preguntasPruebaCorrectas: prueba.estadoInicial,
		}
		for i := 0; i < prueba.iteraciones; i++ {
			estudiante.ResponderPruebaMal()
		}
		if estudiante.preguntasPruebaCorrectas != prueba.estadoEsperado {
			t.Errorf("FALLO: obtenido: %d esperado : %d",
				estudiante.preguntasCorrectas, prueba.estadoEsperado)
		}
	}
}
