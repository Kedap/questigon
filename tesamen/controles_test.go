package main

import "testing"

func TestEliminarPant(t *testing.T) {
	controlador := Controlador{
		pantallaSubscriptora: &PantallaSimple{},
	}

	controlador.EliminarPant()
	var esperado Pantalla = nil

	if controlador.pantallaSubscriptora != esperado {
		t.Error("No se pudo borrar la pantalla subscriptora del controlador")
	}
}

func TestIntercambiarPant(t *testing.T) {
	controlador := Controlador{
		pantallaSubscriptora: &PantallaSimple{},
	}
	nuevaPantalla := &PantallaSimple{
		TituloExamen: "Test",
		Descripcion:  "Test",
	}

	controlador.IntercambiarPant(nuevaPantalla)

	if controlador.pantallaSubscriptora != nuevaPantalla {
		t.Error("No se intercambiar la pantalla subscriptora del controlador")
	}
}
