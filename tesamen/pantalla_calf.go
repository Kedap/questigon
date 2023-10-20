package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/fatih/color"
)

type PantallaCalif struct {
	*PantallaCompuesta
	estudiante *Estudiante
}

func (p *PantallaCalif) Cabezero() {
	fmt.Printf("\x1bc")
	estiloTitulo := color.New(color.FgYellow).Add(color.Bold)
	fmt.Printf("\t\t\t--==[ %s ]==--\n\n%s\n\n\t\t\t", p.TituloExamen, p.Descripcion)
	estiloTitulo.Print(strings.ToUpper(p.titulo))
	promedio := (p.estudiante.preguntasCorrectas * 100) / p.preguntasTotales
	fmt.Printf("\n\n\t%s\t%d\n\t\t%d/%d\t\t\t(%d%%)\n\n\t\t",
		p.estudiante.Nombre, p.estudiante.Grupo,
		p.estudiante.preguntasCorrectas,
		p.preguntasTotales, promedio)
	estiloCalif := color.New()
	if promedio >= 70 {
		estiloCalif.Add(color.FgGreen).Add(color.Bold)
		estiloCalif.Println("APROBADO!!!")
	} else {
		estiloCalif.Add(color.FgRed).Add(color.Bold)
		estiloCalif.Println("NO APROBADO")
	}
	for {
		time.Sleep(time.Minute)
	}
}
