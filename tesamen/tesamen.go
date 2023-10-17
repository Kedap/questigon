package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type JsonExamen struct {
	Nombre    string
	Preguntas []PPregunta
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
	primeraCol := PantallaCalif{
		PantallaCompuesta: &PantallaCompuesta{
			PantallaSimple: &PantallaSimple{
				TituloExamen: "Prueba en Go",
			},
			titulo:           "Calificaciones",
			preguntasTotales: 2,
		},
		estudiante: &daniel,
	}
	seguro := PantallaConfirmacion{
		PantallaCompuesta: &PantallaCompuesta{
			PantallaSimple: &PantallaSimple{
				TituloExamen: "Prueba en Go",
				PSiguiente:   &primeraCol,
				PAnterior:    &examen.Preguntas[2],
			},
			titulo: "\tSeguro",
		},
	}
	examen.Preguntas[2].PantallaCompuesta = &PantallaCompuesta{
		PantallaSimple: &PantallaSimple{
			TituloExamen: "Prueba en Go",
			Descripcion:  "Prueba 2 de 2",
			PAnterior:    &examen.Preguntas[1],
			PSiguiente:   &seguro,
		},
		titulo:           "¿Que dia es hoy?",
		preguntasTotales: 10,
	}
	examen.Preguntas[2].instrucciones = "\t<- Ir a la anterior  -> Ir a la siguiente\n\t/\\ Seleccionar arriba \\/ Seleccionar abajo \n\tENTER elegir opción"
	examen.Preguntas[2].estudiante = &daniel
	examen.Preguntas[1].PantallaCompuesta = &PantallaCompuesta{
		PantallaSimple: &PantallaSimple{
			TituloExamen: "Prueba en Go",
			Descripcion:  "Prueba 1 de 2",
			PSiguiente:   &examen.Preguntas[2],
		},
		titulo:           "¿Que dia fue el eclipse solar?",
		preguntasTotales: 10,
	}
	examen.Preguntas[1].instrucciones = "\t<- Ir a la anterior  -> Ir a la siguiente\n\t/\\ Seleccionar arriba \\/ Seleccionar abajo \n\tENTER elegir opción"
	examen.Preguntas[1].estudiante = &daniel
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
