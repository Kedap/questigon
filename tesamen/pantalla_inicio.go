package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type PantallaIncio struct {
	*PantallaSimple
	Estudiante *Estudiante
}

func (p *PantallaIncio) cuerpo() {
	fmt.Print("\t\tIngresa tu nombre empezando por apellido\n\n> ")
	lector := bufio.NewScanner(os.Stdin)
	lector.Scan()
	err := lector.Err()
	if err != nil {
		println("Ocurrió un error al leer la entrada :c\nPulse ENTER para salir del programa")
		fmt.Scanln()
		os.Exit(1)
	}
	p.Estudiante.Nombre = strings.ToUpper(lector.Text())
	fmt.Print("\n\nIngresa tu grupo\n> ")
	fmt.Scanln(&p.Estudiante.Grupo)
	fmt.Printf("\nEres %s del grupo %d\n", p.Estudiante.Nombre, p.Estudiante.Grupo)
	fmt.Print("¿Es correcto? [S/n]: ")
	var opc string
	fmt.Scanln(&opc)
	if strings.ToLower(opc) != "s" {
		mostrarPantalla(p)
	}
}
