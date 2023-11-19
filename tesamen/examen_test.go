package main

import "testing"

func TestNuevoExamenPantallaInicio(t *testing.T) {
	examen, _ := NuevoExamen("examen.json")

	if examen.primeraPantalla == nil {
		t.Error("FALLO: La primera pantalla esta vacia...")
		t.FailNow()
	} else if examen.primeraPantalla.ObtenerSiguiente() == nil {
		t.Error("FALLO: La pantalla siguiente de la primera pantalla esta vacia...")
		t.FailNow()
	}
}

func TestNuevoExamenPantallaInstrucciones(t *testing.T) {
	examen, _ := NuevoExamen("examen.json")
	instrucciones := examen.primeraPantalla.ObtenerSiguiente()

	if instrucciones == nil {
		t.Error("FALLO: No tiene instrucciones...")
		t.FailNow()
	} else if instrucciones.ObtenerSiguiente() == nil {
		t.Error("FALLO: No hay nada despues de las instrucciones...")
		t.FailNow()
	}
}

func TestNuevoExamenTutoriales(t *testing.T) {
	examen, _ := NuevoExamen("examen.json")
	primeraPregunta := examen.primeraPantalla.ObtenerSiguiente().ObtenerSiguiente()
	segundaPregunta := primeraPregunta.ObtenerSiguiente()

	if primeraPregunta == nil {
		t.Error("FALLO: No tiene la primera pregunta del tutorial...")
		t.FailNow()
	} else if primeraPregunta.ObtenerSiguiente() == nil {
		t.Error("FALLO: No hay nada despues de la primera pregunta del tutorial...")
		t.FailNow()
	}

	if segundaPregunta == nil {
		t.Error("FALLO: No tiene la segunda pregunta del tutorial...")
		t.FailNow()
	} else if segundaPregunta.ObtenerSiguiente() == nil {
		t.Error("FALLO: No hay nada despues de la segunda pregunta del tutorial...")
		t.FailNow()
	} else if segundaPregunta.ObtenerAnterior() == nil {
		t.Error("FALLO: No hay nada antes de las segunda pregunta del tutorial...")
		t.FailNow()
	}

}

func TestNuevoExamenCalificacion(t *testing.T) {
	examen, _ := NuevoExamen("examen.json")
	tutorialCalif := examen.primeraPantalla.ObtenerSiguiente().ObtenerSiguiente().ObtenerSiguiente().ObtenerSiguiente()

	if tutorialCalif == nil {
		t.Error("FALLO: No tiene la calificacion del tutorial...")
		t.FailNow()
	} else if tutorialCalif.ObtenerSiguiente() == nil {
		t.Error("FALLO: No hay nada despues de la calificacion del tutorial...")
		t.FailNow()
	} else if tutorialCalif.ObtenerAnterior() == nil {
		t.Error("FALLO: No hay nada antes de la calificacion del tutorial...")
		t.FailNow()
	}
}

func TestNuevoExamenConfirmacion(t *testing.T) {
	examen, _ := NuevoExamen("examen.json")
	confirmacion := examen.primeraPantalla.ObtenerSiguiente().ObtenerSiguiente().
		ObtenerSiguiente().ObtenerSiguiente().ObtenerSiguiente()

	if confirmacion == nil {
		t.Error("FALLO: No tiene la confirmacion del tutorial...")
		t.FailNow()
	} else if confirmacion.ObtenerSiguiente() == nil {
		t.Error("FALLO: No hay nada despues de la confirmacion del tutorial...")
		t.FailNow()
	}
}
