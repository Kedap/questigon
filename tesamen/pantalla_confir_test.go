package main

import "testing"

func TestNecesitaControlador(t *testing.T) {
	pc := PantallaConfirmacion{}
	if pc.NecesitaControlador() {
		t.Error("La pantalla de confirmacion no necesita controlador")
	}
}
