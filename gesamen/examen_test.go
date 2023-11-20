package main

import "testing"

func TestNuevoExamen(t *testing.T) {
	estudiante := Estudiante{
		Nombre: "Test",
	}
	examen, err := NuevoExamen("examen.json", &estudiante)

	if err != nil {
		t.Errorf("FALLO: %s", err)
	} else if len(examen.Preguntas) == 0 {
		t.Error("FALLO: No puedes tener 0 preguntas")
	} else if examen.Nombre == "" {
		t.Error("FALLO: El examen no puede no tener nombre!")
	}

	const NOMBRE_PRIMERA_PREGUNTA = "¿Cuál de las siguientes es una proposición simple?"
	const RESPUESTA_PRIMERA_PREGUNTA = 1
	const TERCERA_RESPUESTA_PRIMERA_PREGUNTA = "O María estudia o Juan estudia"
	primera_preguanta := examen.Preguntas[0]
	if primera_preguanta.Pregunta != NOMBRE_PRIMERA_PREGUNTA {
		t.Errorf("FALLO: Como nombre de la primera pregunta se esperaba: %s y se obtuvo: %s",
			NOMBRE_PRIMERA_PREGUNTA, primera_preguanta.Pregunta)
	} else if primera_preguanta.RespuestaCorrecta != RESPUESTA_PRIMERA_PREGUNTA {
		t.Errorf("FALLO: Como respuesta de la primera pregunta se esperaba: %d y se obtuvo: %d",
			RESPUESTA_PRIMERA_PREGUNTA, primera_preguanta.RespuestaCorrecta)
	} else if primera_preguanta.Respuestas[2] != TERCERA_RESPUESTA_PRIMERA_PREGUNTA {
		t.Errorf("FALLO: Como tercera respuesta de la primera pregunta se esperaba: %s y se obtuvo: %s",
			TERCERA_RESPUESTA_PRIMERA_PREGUNTA, primera_preguanta.Respuestas[2])
	}

	for _, pregunta := range examen.Preguntas {
		if pregunta.estudiante == nil {
			t.Error("FALLO: Todas las preguntas deberían contener un estudiante")
		}
	}
}
