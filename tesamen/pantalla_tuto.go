package main

import "fmt"

type PantallaTutorial struct {
	*PantallaCompuesta
	instrucciones string
	msgFinal      string
}

func (p *PantallaTutorial) cuerpo() {
	fmt.Println(p.instrucciones)
	fmt.Println(p.msgFinal)
}
