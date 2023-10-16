package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/nsf/termbox-go"
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
	PSiguiente   Pantalla
	PAnterior    Pantalla
	//Conrtoles
}
type Pantalla interface {
	Cabezero()
	cuerpo()
	ObtenerSiguiente() Pantalla
	ObtenerAnterior() Pantalla
	renderizarPregunta()
}

func (p *PantallaSimple) cuerpo()                    {}
func (p *PantallaSimple) renderizarPregunta()        {}
func (p *PantallaSimple) ObtenerSiguiente() Pantalla { return p.PSiguiente }
func (p *PantallaSimple) ObtenerAnterior() Pantalla  { return p.PAnterior }
func (p *PantallaSimple) Cabezero() {
	fmt.Printf("\x1bc")
	fmt.Printf("\t\t\t--==[ %s ]==--\n\n\t\t%s\n\n", p.TituloExamen, p.Descripcion)
}

func mostrarPantalla(p Pantalla) {
	switch p.(type) {
	case *PantallaIncio:
		p.Cabezero()
		p.cuerpo()
		mostrarPantalla(p.ObtenerSiguiente())
	case *PPregunta:
		p.renderizarPregunta()
	default:
		err := termbox.Init()
		if err != nil {
			fmt.Println("Oh no, ocurrio el error", err)
			os.Exit(1)
		}
		defer termbox.Close()
		canalEventos := make(chan termbox.Event)
		go func() {
			for {
				canalEventos <- termbox.PollEvent()
			}
		}()
		p.Cabezero()
		p.cuerpo()
		for {
			select {
			case ev := <-canalEventos:
				if ev.Type == termbox.EventKey {
					switch {
					case ev.Ch == 68:
						anterior := p.ObtenerAnterior()
						if anterior != nil {
							termbox.Close()
							mostrarPantalla(anterior)
						}
					case ev.Ch == 67:
						siguiente := p.ObtenerSiguiente()
						if siguiente != nil {
							termbox.Close()
							mostrarPantalla(siguiente)
						}
					}
				}
			}
		}
	}
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
	Pregunta              string
	Respuestas            []string
	RespuestaCorrecta     int
	estudiante            Estudiante
	respuestaSeleccionada int
	respuestaElegida      int
	instrucciones         string
}
type Pregunta interface {
	responder()
}

func (p *PPregunta) responder() {
	if p.respuestaElegida == p.RespuestaCorrecta {
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
func (p *PPregunta) renderizarPregunta() {
	p.Cabezero()
	p.cuerpo()
	err := termbox.Init()
	if err != nil {
		fmt.Println("Oh no, ocurrio el error al inicializar los controladores:", err)
		os.Exit(1)
	}
	defer termbox.Close()
	canalEventos := make(chan termbox.Event)
	go func() {
		for {
			canalEventos <- termbox.PollEvent()
		}
	}()
	p.Cabezero()
	p.cuerpo()
	fmt.Println("\n\n", p.instrucciones)
	for {
		select {
		case ev := <-canalEventos:
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Ch == 68:
					anterior := p.ObtenerAnterior()
					if anterior != nil {
						termbox.Close()
						mostrarPantalla(anterior)
					}
				case ev.Ch == 67:
					siguiente := p.ObtenerSiguiente()
					if siguiente != nil {
						termbox.Close()
						mostrarPantalla(siguiente)
					}
				case ev.Ch == 65:
					if p.respuestaSeleccionada-1 != -1 {
						p.respuestaSeleccionada--
					}
				case ev.Ch == 66:
					if p.respuestaSeleccionada+1 != len(p.Respuestas) {
						p.respuestaSeleccionada++
					}
				case ev.Key == termbox.KeyEnter:
					p.respuestaElegida = p.respuestaSeleccionada
					p.responder()
				}
				time.Sleep(60 * time.Microsecond)
				p.renderizarPregunta()
			}
		}
	}
}

func main() {
	daniel := Estudiante{}
	jsonFile, err := os.Open("examen.json")
	if err != nil {
		fmt.Println("Ocurrio un error al leer el examen:c")
		os.Exit(1)
	}
	defer jsonFile.Close()
	valorBytes, _ := io.ReadAll(jsonFile)
	var examen JsonExamen
	json.Unmarshal(valorBytes, &examen)
	examen.Preguntas[2].PantallaCompuesta = &PantallaCompuesta{
		PantallaSimple: &PantallaSimple{
			TituloExamen: "Prueba en Go",
			Descripcion:  "Prueba 1 de 2",
			PAnterior:    &examen.Preguntas[1],
		},
		titulo:           "¿Que dia es hoy?",
		preguntasTotales: 10,
	}
	examen.Preguntas[2].instrucciones = "\t<- Ir a la anterior  -> Ir a la siguiente\n\t/\\ Seleccionar arriba \\/ Seleccionar abajo \n\tENTER elegir opción"
	examen.Preguntas[1].PantallaCompuesta = &PantallaCompuesta{
		PantallaSimple: &PantallaSimple{
			TituloExamen: "Prueba en Go",
			Descripcion:  "Prueba 2 de 2",
			PSiguiente:   &examen.Preguntas[2],
		},
		titulo:           "¿Que dia fue el eclipse solar?",
		preguntasTotales: 10,
	}
	examen.Preguntas[1].instrucciones = "\t<- Ir a la anterior  -> Ir a la siguiente\n\t/\\ Seleccionar arriba \\/ Seleccionar abajo \n\tENTER elegir opción"
	const TEXTO_TUTORIAL = `
          <- (Flecha izquierda) Ir a la anterior

          -> (Flecha derecha) Ir a la siguiente

          /\ (Flecha arriba) Seleccionar arriba


          \/ (Flecha abajo) Seleccionar abajo


          ENTER Elegir la opción
        `
	instrucciones := PantallaTutorial{
		PantallaCompuesta: &PantallaCompuesta{
			PantallaSimple: &PantallaSimple{
				TituloExamen: "Prueba en Go",
				Descripcion:  "Tutorial",
				PSiguiente:   &examen.Preguntas[1],
			},
			titulo: "instrucciones",
		},
		instrucciones: TEXTO_TUTORIAL,
		msgFinal:      "Las siguientes preguntas solo seran de prueba",
	}
	primera := PantallaIncio{
		PantallaSimple: &PantallaSimple{
			TituloExamen: "Pruba en Go",
			PSiguiente:   &instrucciones,
		},
		Estudiante: &daniel,
	}
	mostrarPantalla(&primera)
}
