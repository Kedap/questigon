package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Pregunta struct {
	Pregunta          string
	Respuestas        []string
	RespuestaCorrecta int
}

type Examen struct {
	Nombre      string
	Preguntas   []Pregunta
	rutaArchivo string
}

func main() {
	fmt.Print("Escribe el lugar en donde quieres guardar tu nuevo examen (por defecto 'examen.json')\n> ")
	var ruta string
	fmt.Scanln(&ruta)
	if ruta == "" {
		fmt.Println("Seleccionando 'examen.json'")
		ruta = "examen.json"
	}
	nExamen := Examen{
		rutaArchivo: ruta,
	}
	fmt.Print("Ingresa el nombre de tu examen\n> ")
	lector := bufio.NewReader(os.Stdin)
	nombre, err := lector.ReadString('\n')
	if err != nil {
		fmt.Println("Ocurrio un error al leer la entrada :/")
		os.Exit(1)
	}
	nExamen.Nombre = strings.Trim(nombre, "\n")
	fmt.Print("¿Cuantas preguntas quieres que tenga tu examen?\n> ")
	var nPreguntas int
	fmt.Scanln(&nPreguntas)
	fmt.Print("¿Cuantas respuestas quieres que tengan tus preguntas de tu examen?\n> ")
	var nRespuestas int
	fmt.Scanln(&nRespuestas)

	for i := 0; i < nPreguntas; i++ {
		fmt.Printf("Ingresa tu pregunta %d\n", i+1)
		nPregunta, err :=
			lector.ReadString('\n')
		if err != nil {
			fmt.Println("Ocurrio un error al leer la entrada :/")
			os.Exit(1)
		}
		pregunta := Pregunta{
			Pregunta: strings.Trim(nPregunta, "\n"),
		}
		for j := 0; j < nRespuestas; j++ {
			fmt.Printf("Ingresa la respuesta %d para la pregunta %d\n", j+1, i+1)
			respuesta, err :=
				lector.ReadString('\n')
			if err != nil {
				fmt.Println("Ocurrio un error al leer la entrada :/")
				os.Exit(1)
			}
			pregunta.Respuestas = append(pregunta.Respuestas, strings.Trim(respuesta, "\n"))
		}
		fmt.Printf("Coloca el número de la respuesta correcta para la pregunta %d (1-%d): ", i+1, nRespuestas)
		var correcta int
		fmt.Scanln(&correcta)
		pregunta.RespuestaCorrecta = correcta
		nExamen.Preguntas = append(nExamen.Preguntas, pregunta)
	}

	datos, _ := json.Marshal(nExamen)
	f, err := os.Create(nExamen.rutaArchivo)
	if err != nil {
		fmt.Println("Ocurrio un grave error al intentar crear el archivo", nExamen.rutaArchivo)
		os.Exit(1)
	}
	defer f.Close()
	_, err = f.Write(datos)
	if err != nil {
		fmt.Println("Ocurrio un grave error al intentar escribir el archivo", nExamen.rutaArchivo)
		os.Exit(1)
	}
	fmt.Printf("Se termino de escribir el examen '%s' en el archivo %s", nExamen.Nombre, nExamen.rutaArchivo)
}
