package main

import "fmt"

type PantallaTutorial struct {
	*PantallaCompuesta
	instrucciones string
	msgFinal      string
}

func (p *PantallaTutorial) cuerpo() {
	fmt.Println(p.instrucciones, "\nEl examen tiene", p.preguntasTotales, "preguntas")
	fmt.Println(p.msgFinal)
}
