package main

import "testing"

func TestPreguntaResponder(t *testing.T) {
	pregunta := Pregunta{
		Pregunta:          "¿Cual es el dia de la independencia?",
		Respuestas:        []string{"11/11", "25/12", "16/9"},
		RespuestaCorrecta: 3,
	}

	const RESPUESTA_CORRECTA = "16/9"
	pregunta.Responder(RESPUESTA_CORRECTA)

	if !pregunta.resuelta {
		t.Error("FATAL: El metodo responder no funciona correctamente")
	}
}
