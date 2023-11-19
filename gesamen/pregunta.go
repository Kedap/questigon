package main

type Pregunta struct {
	Pregunta          string
	Respuestas        []string
	RespuestaCorrecta int
	resuelta          bool
}

func (p *Pregunta) Responder(respuesta int) {
	p.resuelta = respuesta == p.RespuestaCorrecta
}
