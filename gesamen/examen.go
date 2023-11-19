package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

type Examen struct {
	Nombre    string
	Preguntas []Pregunta
}

func NuevoExamen(ruta string) (*Examen, error) {
	_, err := os.Stat(ruta)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("El archivo %s no existe!", ruta)
	}
	archivoJson, err := os.Open(ruta)
	if err != nil {
		return nil, err
	}
	defer archivoJson.Close()
	valorBytes, err := io.ReadAll(archivoJson)
	if err != nil {
		return nil, err
	}

	var nuevoExamen Examen
	json.Unmarshal(valorBytes, &nuevoExamen)

	if len(nuevoExamen.Preguntas) == 0 {
		return nil, fmt.Errorf("El examen contiene 0 preguntas")
	} else if nuevoExamen.Nombre == "" {
		return nil, errors.New("FALLO: El examen no puede no tener nombre!")
	}

	return &nuevoExamen, nil
}
