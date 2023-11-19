package main

type Pregunta struct {
	Pregunta          string
	Respuestas        []string
	RespuestaCorrecta int
	resuelta          bool
}

func (p *Pregunta) Responder(respuesta string) {
	p.resuelta = respuesta == p.Respuestas[p.RespuestaCorrecta-1]
}
