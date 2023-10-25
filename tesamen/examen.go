package main

import (
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
)

// Definitivamente es el examen
type Examen struct {
	Nombre          string
	Preguntas       []PPregunta
	primeraPantalla Pantalla
}

func NuevoExamen(ruta string) Examen {
	_, err := os.Stat(ruta)
	if os.IsNotExist(err) {
		fmt.Println("El archivo", ruta, "no existe :/")
		os.Exit(1)
	}
	nuevoEstudiante := Estudiante{}
	jsonFile, err := os.Open("examen.json")
	if err != nil {
		fmt.Println("Ocurrió un error al leer el examen:c")
		os.Exit(1)
	}
	defer jsonFile.Close()
	valorBytes, err := io.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Ocurrió un error al leer el examen:c")
		os.Exit(1)
	}
	verificarExamen(valorBytes)
	var nuevoExamen Examen
	json.Unmarshal(valorBytes, &nuevoExamen)
	preguntasTotales := len(nuevoExamen.Preguntas)
	if preguntasTotales == 0 {
		fmt.Println("El archivo", ruta, "no tiene ninguna pregunta :/")
		os.Exit(1)
	}
	pCalif := PantallaCalif{
		PantallaCompuesta: &PantallaCompuesta{
			PantallaSimple: &PantallaSimple{
				TituloExamen: nuevoExamen.Nombre,
			},
			titulo:           "Calificaciones",
			preguntasTotales: preguntasTotales,
		},
		estudiante: &nuevoEstudiante,
	}
	nuevoExamen.jsonAPantallas(&nuevoEstudiante)
	pConfir := PantallaConfirmacion{
		PantallaCompuesta: &PantallaCompuesta{
			PantallaSimple: &PantallaSimple{
				TituloExamen: nuevoExamen.Nombre,
				Descripcion:  "Confirmación",
				PSiguiente:   &pCalif,
				PAnterior:    &nuevoExamen.Preguntas[preguntasTotales-1],
			},
			titulo: "\tSeguro",
		},
	}
	nuevoExamen.Preguntas[preguntasTotales-1].PSiguiente = &pConfir
	tutorialCalf := TutorialCalificacion{
		PantallaCalif: &PantallaCalif{
			PantallaCompuesta: &PantallaCompuesta{
				PantallaSimple: &PantallaSimple{
					TituloExamen: nuevoExamen.Nombre,
					PSiguiente:   &nuevoExamen.Preguntas[0],
				},
				titulo:           "Calificaciones (TUTORIAL)",
				preguntasTotales: 2,
			},
			estudiante: &nuevoEstudiante,
		},
	}
	segundaTutorial := TutorialPregunta{
		PPregunta: &PPregunta{
			PantallaCompuesta: &PantallaCompuesta{
				PantallaSimple: &PantallaSimple{
					TituloExamen: nuevoExamen.Nombre,
					Descripcion:  "Examen de TUTORIAL 2 de 2",
				},
				titulo:           "Un ejemplo de verbo es...",
				preguntasTotales: 2,
			},
			Pregunta:          "Un ejemplo de verbo es...",
			Respuestas:        []string{"Pregunta", "Martillo", "Correr"},
			RespuestaCorrecta: 3,
			estudiante:        &nuevoEstudiante,
			instrucciones:     "\t<- Ir a la anterior  -> Ir a la siguiente\n\t/\\ Seleccionar arriba \\/ Seleccionar abajo \n\tENTER elegir opción",
		},
	}
	primeraTutorial := TutorialPregunta{
		PPregunta: &PPregunta{
			PantallaCompuesta: &PantallaCompuesta{
				PantallaSimple: &PantallaSimple{
					TituloExamen: nuevoExamen.Nombre,
					Descripcion:  "Examen de TUTORIAL 1 de 2",
					PSiguiente:   &segundaTutorial,
				},
				titulo:           "¿Cual es el día de la independencia?",
				preguntasTotales: 2,
			},
			Pregunta:          "¿Cual es el día de la independencia?",
			Respuestas:        []string{"11 de noviembre", "16 de septiembre", "2 de marzo"},
			RespuestaCorrecta: 2,
			estudiante:        &nuevoEstudiante,
			instrucciones:     "\t<- Ir a la anterior  -> Ir a la siguiente\n\t/\\ Seleccionar arriba \\/ Seleccionar abajo \n\tENTER elegir opción",
		},
	}
	segundaTutorial.PAnterior = &primeraTutorial
	tutorialConfr := PantallaConfirmacion{
		PantallaCompuesta: &PantallaCompuesta{
			PantallaSimple: &PantallaSimple{
				TituloExamen: nuevoExamen.Nombre,
				PSiguiente:   &tutorialCalf,
				PAnterior:    &segundaTutorial,
			},
			titulo: "\tSeguro",
		},
	}
	segundaTutorial.PSiguiente = &tutorialConfr
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
				TituloExamen: nuevoExamen.Nombre,
				Descripcion:  "Tutorial",
				PSiguiente:   &primeraTutorial,
			},
			titulo: "Instrucciones",
		},
		instrucciones: TEXTO_TUTORIAL,
		msgFinal:      "LAS SIGUIENTES PREGUNTAS SOLO SERÁN DE PRUEBA",
	}
	nuevoExamen.primeraPantalla = &PantallaIncio{
		PantallaSimple: &PantallaSimple{
			TituloExamen: nuevoExamen.Nombre,
			PSiguiente:   &instrucciones,
		},
		Estudiante: &nuevoEstudiante,
	}

	return nuevoExamen
}

func (e *Examen) jsonAPantallas(est *Estudiante) {
	preguntasTotales := len(e.Preguntas) - 1
	for i, v := range e.Preguntas {
		var siguiente, anterior Pantalla
		if i == preguntasTotales {
			siguiente = nil
		} else {
			siguiente = &e.Preguntas[i+1]
		}
		if i == 0 {
			anterior = nil
		} else {
			anterior = &e.Preguntas[i-1]
		}
		nuevP := PPregunta{
			PantallaCompuesta: &PantallaCompuesta{
				PantallaSimple: &PantallaSimple{TituloExamen: e.Nombre,
					Descripcion: fmt.Sprintf("%d de %d",
						i+1,
						preguntasTotales+1),
					PSiguiente: siguiente,
					PAnterior:  anterior},
				titulo:           v.Pregunta,
				preguntasTotales: preguntasTotales + 1},
			Pregunta:          v.Pregunta,
			Respuestas:        v.Respuestas,
			RespuestaCorrecta: v.RespuestaCorrecta,
			estudiante:        est,
			instrucciones:     "\t<- Ir a la anterior  -> Ir a la siguiente\n\t/\\ Seleccionar arriba \\/ Seleccionar abajo \n\tENTER elegir opción",
		}
		e.Preguntas[i] = nuevP
	}
}

func verificarExamen(b []byte) {
	var sha1_esperado string
	if runtime.GOOS == "windows" {
		sha1_esperado = "be4d43e56f4e2bf82f513588023d2280ca63778d"
	} else {
		sha1_esperado = "661cb8c527194d7bfe38c99d9dc758b3b64121ea"
	}
	if sha1_esperado != fmt.Sprintf("%x", sha1.Sum(b)) {
		fmt.Println("El archivo fue modificado :/\nTu examen no es valido")
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
}
