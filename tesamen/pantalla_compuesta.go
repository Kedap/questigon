package main

import (
	"fmt"
	"strings"
)

type PantallaCompuesta struct {
	*PantallaSimple
	titulo           string
	preguntasTotales int
}

func (p *PantallaCompuesta) Cabezero() {
	fmt.Printf("\x1bc")
	fmt.Printf("\t\t\t--==[ %s ]==--\n\n%s\n\n\t\t\t%s\n\n", p.TituloExamen, p.Descripcion, strings.ToUpper(p.titulo))
}
