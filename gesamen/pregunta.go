package main

type Pregunta struct {
	Pregunta          string
	Respuestas        []string
	RespuestaCorrecta int
	resuelta          bool
	estudiante        *Estudiante
}

func (p *Pregunta) Responder(respuesta string) {
	if respuesta == p.Respuestas[p.RespuestaCorrecta-1] && !p.resuelta {
		p.estudiante.RespuestasCorrectas++
		p.resuelta = true
	} else if respuesta != p.Respuestas[p.RespuestaCorrecta-1] && p.resuelta {
		p.estudiante.RespuestasCorrectas--
		p.resuelta = false
	}
}
