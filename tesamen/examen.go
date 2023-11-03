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

/*
La estructura Examen almacena información sobre el examen, incluyendo su
nombre, una lista de preguntas (Preguntas), y la primera pantalla que se
mostrará al comenzar el examen (primeraPantalla).

Definición de la estructura Examen que representa un examen.
*/
type Examen struct {
	Nombre          string
	Preguntas       []PPregunta
	primeraPantalla Pantalla
}

/*
La función NuevoExamen se encarga de crear una instancia de Examen a partir de
un archivo JSON, configurando sus componentes, incluyendo pantallas de inicio,
preguntas, tutoriales y pantallas de calificación

Función NuevoExamen crea y configura un nuevo examen a partir de un archivo JSON.
*/
func NuevoExamen(ruta string) (Examen, Controlador) {
	_, err := os.Stat(ruta)
	if os.IsNotExist(err) {
		fmt.Println("El archivo", ruta, "no existe :/")
		os.Exit(1)
	}

	nuevoEstudiante := Estudiante{} // Crea una instancia de Estudiante.

	jsonFile, err := os.Open("examen.json") // Abre el archivo JSON del examen.
	if err != nil {
		fmt.Println("Ocurrió un error al leer el examen:c")
		os.Exit(1)
	}
	defer jsonFile.Close()

	valorBytes, err := io.ReadAll(jsonFile) // Lee el contenido del archivo JSON.
	if err != nil {
		fmt.Println("Ocurrió un error al leer el examen:c")
		os.Exit(1)
	}

	verificarExamen(valorBytes) // Verifica el formato del examen JSON.

	var nuevoExamen Examen                   // Crea una instancia de Examen.
	json.Unmarshal(valorBytes, &nuevoExamen) // Decodifica el JSON en la instancia de Examen.

	preguntasTotales := len(nuevoExamen.Preguntas) // Obtiene la cantidad de preguntas en el examen.
	if preguntasTotales == 0 {
		fmt.Println("El archivo", ruta, "no tiene ninguna pregunta :/")
		time.Sleep(time.Second)
		os.Exit(1)
	}

	// Configura una pantalla de calificación para el examen.
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

	// Convierte el examen en una secuencia de pantallas y configura la pantalla de confirmación.
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

	// Configura una pantalla de calificación para el tutorial.
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

	// Configura las preguntas de tutorial.
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
	nuevoControlador := Controlador{
		pantallaSubscriptora: &instrucciones,
	}

	// Configura la primera pantalla de inicio del examen.
	nuevoExamen.primeraPantalla = &PantallaIncio{
		PantallaSimple: &PantallaSimple{
			TituloExamen: nuevoExamen.Nombre,
			PSiguiente:   &instrucciones,
		},
		Estudiante: &nuevoEstudiante,
	}

	return nuevoExamen, nuevoControlador
}

/*
La función toma las preguntas de un examen, las convierte en pantallas de
preguntas y configura las relaciones entre las pantallas, como la pantalla
siguiente y anterior

jsonAPantallas convierte las preguntas del examen en pantallas de preguntas y las configura.
*/
func (e *Examen) jsonAPantallas(est *Estudiante) {
	preguntasTotales := len(e.Preguntas) - 1

	// Itera a través de las preguntas y crea las pantallas correspondientes.
	for i, v := range e.Preguntas {
		var siguiente, anterior Pantalla

		// Configura la pantalla siguiente y anterior según la posición de la pregunta.
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

		// Crea una nueva pantalla de pregunta a partir de la pregunta original.
		nuevaPregunta := PPregunta{
			PantallaCompuesta: &PantallaCompuesta{
				PantallaSimple: &PantallaSimple{
					TituloExamen: e.Nombre,
					Descripcion: fmt.Sprintf("%d de %d",
						i+1,
						preguntasTotales+1),
					PSiguiente: siguiente,
					PAnterior:  anterior,
				},
				titulo:           v.Pregunta,
				preguntasTotales: preguntasTotales + 1,
			},
			Pregunta:          v.Pregunta,
			Respuestas:        v.Respuestas,
			RespuestaCorrecta: v.RespuestaCorrecta,
			estudiante:        est,
			instrucciones:     "\t<- Ir a la anterior  -> Ir a la siguiente\n\t/\\ Seleccionar arriba \\/ Seleccionar abajo \n\tENTER elegir opción",
		}
		e.Preguntas[i] = nuevaPregunta
	}
}

/*
La función calcula el hash SHA-1 del archivo de examen y lo compara con un
valor esperado específico para el sistema operativo en uso. Si los valores no
coinciden, se imprime un mensaje de que el archivo fue modificado y el programa
se cierra. Esta función se utiliza para garantizar que el archivo del examen no
haya sido alterado.

verificarExamen verifica la integridad del archivo de examen mediante un hash SHA-1.
*/
func verificarExamen(b []byte) {
	var sha1_esperado string
	// Determina el valor esperado de SHA-1 basado en el sistema operativo.
	if runtime.GOOS == "windows" {
		sha1_esperado = "339558623d89d82b95cb9a971664990730c230c5"
	} else {
		sha1_esperado = "da8569639ca976d21dc2f276f9208ffd6fb2ae92"
	}
	// Calcula el hash SHA-1 del archivo y lo compara con el valor esperado.
	if sha1_esperado != fmt.Sprintf("%x", sha1.Sum(b)) {
		fmt.Println("El archivo fue modificado :/\nTu examen no es valido")
		time.Sleep(5 * time.Second)
		os.Exit(1)
	}
}
