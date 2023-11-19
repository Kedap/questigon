package main

import "testing"

func TestPreguntaResponder(t *testing.T) {
	pregunta := Pregunta{
		Pregunta:          "¿Cual es el dia de la independencia?",
		Respuestas:        []string{"11/11", "25/12", "16/9"},
		RespuestaCorrecta: 3,
	}

	if pregunta.resuelta {
		t.Error("FALLO: La pregunta no puede estar resuelta por defecto")
	}

	const RESPUESTA_CORRECTA = "16/9"
	pregunta.Responder(RESPUESTA_CORRECTA)

	if !pregunta.resuelta {
		t.Error("FALLO: El metodo responder no funciona correctamente")
	}
}
