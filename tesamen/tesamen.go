package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/fatih/color"
)

type JsonExamen struct {
	Nombre    string
	Preguntas []PPregunta
}

type Estudiante struct {
	Nombre                   string
	Grupo                    int
	preguntasCorrectas       int
	preguntasPruebaCorrectas int
}

func (e *Estudiante) ResponderBien() {
	e.preguntasCorrectas++
}
func (e *Estudiante) ResponderMal() {
	e.preguntasCorrectas--
}
func (e *Estudiante) ResponderPruebaBien() {
	e.preguntasPruebaCorrectas++
}
func (e *Estudiante) ResponderPruebaMal() {
	e.preguntasPruebaCorrectas--
}

type PantallaSimple struct {
	TituloExamen string
	Descripcion  string
	PSiguiente   *Pantalla
	PAnterior    *Pantalla
	//Conrtoles
}
type Pantalla interface {
	Cabezero()
	cuerpo()
}

func (p *PantallaSimple) cuerpo() {}
func (p *PantallaSimple) Cabezero() {
	fmt.Printf("\x1bc")
	fmt.Printf("\t\t\t--==[ %s ]==--\n\n\t\t%s\n\n", p.TituloExamen, p.Descripcion)
}
func mostrarPantalla(p Pantalla) {
	p.Cabezero()
	p.cuerpo()
}

type PantallaIncio struct {
	*PantallaSimple
	Estudiante *Estudiante
}

func (p *PantallaIncio) cuerpo() {
	fmt.Print("\t\tIngresa tu nombre empezando por apellido\n\n> ")
	lector := bufio.NewReader(os.Stdin)
	nombre, err := lector.ReadString('\n')
	if err != nil {
		println("Ocurrio un error al leer la entrada :c\nPulse ENTER para salir del programa")
		fmt.Scanln()
		os.Exit(1)
	}
	nombre = strings.ToUpper(nombre)
	p.Estudiante.Nombre = strings.Trim(nombre, "\n")
	fmt.Print("\n\nIngresa tu grupo\n> ")
	fmt.Scanln(&p.Estudiante.Grupo)
	fmt.Printf("\nNombre: %sGrupo: %d\n", nombre, p.Estudiante.Grupo)
	fmt.Print("¿Es correcto? [S/n]: ")
	var opc string
	fmt.Scanln(&opc)
	if strings.ToLower(opc) != "s" {
		mostrarPantalla(p)
	}
}

type PantallaCompuesta struct {
	*PantallaSimple
	titulo           string
	preguntasTotales int
}

func (p *PantallaCompuesta) Cabezero() {
	fmt.Printf("\x1bc")
	fmt.Printf("\t\t\t--==[ %s ]==--\n\n%s\n\n\t\t\t%s\n\n", p.TituloExamen, p.Descripcion, strings.ToUpper(p.titulo))
}

type PantallaTutorial struct {
	*PantallaCompuesta
	instrucciones string
	msgFinal      string
}

func (p *PantallaTutorial) cuerpo() {
	fmt.Println(p.instrucciones, "\nEl examen tiene", p.preguntasTotales, "preguntas")
	fmt.Println(p.msgFinal)
}

type PPregunta struct {
	*PantallaCompuesta
	Pregunta          string
	Respuestas        []string
	RespuestaCorrecta int
	//Controladores
	estudiante            *Estudiante
	respuestaSeleccionada int
	respuestaElegida      int
}
type Pregunta interface {
	responder(respuesta int)
}

func (p *PPregunta) responder(respuesta int) {
	if respuesta == p.RespuestaCorrecta {
		p.estudiante.ResponderBien()
	} else {
		p.estudiante.ResponderMal()
	}
}

func (p *PPregunta) Cabezero() {
	fmt.Printf("\x1bc")
	fmt.Printf("\t\t\t--==[ %s ]==--\n\n%s\n\n\t\t\t", p.TituloExamen, p.Descripcion)
	estiloPregunta := color.New(color.FgWhite).Add(color.Italic).Add(color.Bold)
	estiloPregunta.Println(p.titulo)
}
func (p *PPregunta) cuerpo() {
	estiloRespuesta := color.New(color.FgWhite).Add(color.BgMagenta).Add(color.Bold)
	for i, e := range p.Respuestas {
		var respuesta string
		if i == p.respuestaElegida {
			respuesta = "[*] " + e
		} else {
			respuesta = "[ ] " + e
		}
		if i == p.respuestaSeleccionada {
			estiloRespuesta.Println(respuesta)
		} else {
			fmt.Println(respuesta)
		}
	}
}

func main() {
	daniel := Estudiante{}
	primera := PantallaIncio{
		PantallaSimple: &PantallaSimple{
			TituloExamen: "Pruba en Go",
		},
		Estudiante: &daniel,
	}
	mostrarPantalla(&primera)
	const TEXTO_TUTORIAL = `
          <- (Flecha izquierda) Ir a la anterior

          -> (Flecha derecha) Ir a la siguiente

          /\ (Flecha arriba) Seleccionar arriba
          ||

          || (Flecha abajo) Seleccionar abajo
          \/

          ENTER Elegir la opción
        `
	instrucciones := PantallaTutorial{
		PantallaCompuesta: &PantallaCompuesta{
			PantallaSimple: &PantallaSimple{
				TituloExamen: "Prueba en Go",
				Descripcion:  "Tutorial",
			},
			titulo: "instrucciones",
		},
		instrucciones: TEXTO_TUTORIAL,
		msgFinal:      "Las siguientes preguntas solo seran de prueba",
	}
	mostrarPantalla(&instrucciones)
	fmt.Scanln()
	fmt.Println("Leyendo examen...")
	jsonFile, err := os.Open("examen.json")
	if err != nil {
		fmt.Println("Ocurrio un error al leer el examen:c")
		os.Exit(1)
	}
	defer jsonFile.Close()
	valorBytes, _ := io.ReadAll(jsonFile)
	var examen JsonExamen
	json.Unmarshal(valorBytes, &examen)
	fmt.Println(examen)
	fmt.Scanln()
	// examen.Preguntas[0] = PPregunta{
	// 	PantallaCompuesta: &PantallaCompuesta{
	// 		PantallaSimple: &PantallaSimple{
	// 			TituloExamen: "Prueba en Go",
	// 			Descripcion:  "1 de 10",
	// 		},
	// 		titulo:           "¿Que dia fue el eclipse solar?",
	// 		preguntasTotales: 10,
	// 	},
	// 	estudiante: &daniel,
	// }
	examen.Preguntas[1].PantallaCompuesta = &PantallaCompuesta{
		PantallaSimple: &PantallaSimple{
			TituloExamen: "Prueba en Go",
			Descripcion:  "1 de 10",
		},
		titulo:           "¿Que dia fue el eclipse solar?",
		preguntasTotales: 10,
	}
	mostrarPantalla(&examen.Preguntas[1])
}
