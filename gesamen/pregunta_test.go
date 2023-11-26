package main

import "testing"

func TestPreguntaResponder(t *testing.T) {
	pregunta := Pregunta{
		Pregunta:          "¿Cual es el dia de la independencia?",
		Respuestas:        []string{"11/11", "25/12", "16/9"},
		RespuestaCorrecta: 3,
		estudiante:        &Estudiante{Nombre: "Test"},
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

func TestPreguntaResponderDoble(t *testing.T) {
	pregunta := Pregunta{
		Pregunta:          "¿Cual es el dia de la independencia?",
		Respuestas:        []string{"11/11", "25/12", "16/9"},
		RespuestaCorrecta: 3,
		estudiante:        &Estudiante{Nombre: "Test"},
	}

	if pregunta.resuelta {
		t.Error("FALLO: La pregunta no puede estar resuelta por defecto")
	}

	const RESPUESTA_CORRECTA = "16/9"
	pregunta.Responder(RESPUESTA_CORRECTA)
	pregunta.Responder(RESPUESTA_CORRECTA)

	if !pregunta.resuelta {
		t.Error("FALLO: El metodo responder no funciona correctamente")
	}
}

func TestPreguntaResponderEstudiante(t *testing.T) {
	estudiante := Estudiante{
		Nombre:              "Test",
		RespuestasCorrectas: 0,
	}
	pregunta := Pregunta{
		Pregunta:          "¿Cual es el dia de la independencia?",
		Respuestas:        []string{"11/11", "25/12", "16/9"},
		RespuestaCorrecta: 3,
		estudiante:        &estudiante,
	}

	const RESPUESTA_CORRECTA = "16/9"
	pregunta.Responder(RESPUESTA_CORRECTA)
	if estudiante.RespuestasCorrectas != 1 {
		t.Error("FALLO: Responder la pregunta no hace cambios en el estudiante!")
	} else if !pregunta.resuelta {
		t.Error("FALLO: Responder la pregunta no hace cambios en la pregunta")
	}

	const RESPUESTA_INCORRECTA = "11/11"
	pregunta.Responder(RESPUESTA_INCORRECTA)
	if estudiante.RespuestasCorrectas != 0 {
		t.Error("FALLO: Volver a responder mal la pregunta deberia decrementar una respuesta correcta")
	} else if pregunta.resuelta {
		t.Error("FALLO: Responder la pregunta no hace cambios en la pregunta, incluso cuando se respondio mal")
	}
}

func TestPreguntaResponderEstudianteNoResuelta(t *testing.T) {
	estudiante := Estudiante{
		Nombre:              "Test",
		RespuestasCorrectas: 0,
	}
	pregunta := Pregunta{
		Pregunta:          "¿Cual es el dia de la independencia?",
		Respuestas:        []string{"11/11", "25/12", "16/9"},
		RespuestaCorrecta: 3,
		estudiante:        &estudiante,
	}

	const RESPUESTA_INCORRECTA = "11/11"
	pregunta.Responder(RESPUESTA_INCORRECTA)
	if estudiante.RespuestasCorrectas != 0 {
		t.Error("FALLO: Volver a responder mal la pregunta deberia decrementar una respuesta correcta")
	} else if pregunta.resuelta {
		t.Error("FALLO: Responder la pregunta no hace cambios en la pregunta, incluso cuando se respondio mal")
	}
}
