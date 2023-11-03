package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

/*
Esta parte del código representa la pantalla de inicio en la que el estudiante
ingresa su nombre y grupo. El nombre se almacena en mayúsculas, y se pregunta
al usuario si la información es correcta. Si la respuesta es "n", la pantalla
se muestra nuevamente.

Definición de la estructura PantallaInicio que extiende la estructura PantallaSimple.
*/
type PantallaIncio struct {
	*PantallaSimple             // Referencia a una pantalla simple.
	Estudiante      *Estudiante // Referencia al objeto Estudiante asociado.
}

// Implementación del método cuerpo para PantallaInicio.
func (p *PantallaIncio) cuerpo() {
	fmt.Print("\t\tIngresa tu nombre empezando por apellido\n\n> ") // Imprime un mensaje para ingresar el nombre.
	lector := bufio.NewScanner(os.Stdin)                            // Crea un escáner para leer desde la entrada estándar.
	lector.Scan()                                                   // Lee la entrada del usuario.
	err := lector.Err()                                             // Comprueba si hubo algún error durante la lectura.
	if err != nil {
		println("Ocurrió un error al leer la entrada :c\nPulse ENTER para salir del programa") // Muestra un mensaje de error.
		fmt.Scanln()                                                                           // Espera a que el usuario presione Enter para salir del programa.
		os.Exit(1)                                                                             // Sale del programa con código de error 1.
	}
	p.Estudiante.Nombre = strings.ToUpper(lector.Text()) // Almacena el nombre del estudiante en mayúsculas.

	fmt.Print("\n\nIngresa tu grupo\n> ") // Solicita al usuario ingresar el grupo.
	var grupo string
	estiloError := color.New(color.FgRed).Add(color.Bold)
	fmt.Scanln(&grupo) // Lee el número de grupo ingresado por el usuario.
	nGrupo, err := strconv.Atoi(grupo)
	if err != nil {
		estiloError.Println("El grupo debe de ser un numero entero chico :/")
		estiloError.Println("Vuelve a intentarlo")
		time.Sleep(3 * time.Second)
		mostrarPantalla(p)
		return
	}
	p.Estudiante.Grupo = nGrupo
	fmt.Printf("\nEres %s del grupo %d\n", p.Estudiante.Nombre, p.Estudiante.Grupo) // Muestra la información ingresada.

	fmt.Print("¿Es correcto? [S/n]: ") // Pregunta al usuario si la información es correcta.
	var opc string
	fmt.Scanln(&opc) // Lee la respuesta del usuario.
	opc = opc[0:1]   // Obtener el primer caracter del string
	opc = strings.ToLower(opc)
	if opc == "n" { // Si la respuesta no es "s" (significa no es correcta), vuelve a mostrar la misma pantalla.
		estiloError.Println("Volveras a ingresar tus datos")
		estiloError.Println("Esta vez ingresalos bien!")
		time.Sleep(3 * time.Second)
		mostrarPantalla(p)
	} else if opc != "s" {
		estiloError.Println("Tu opción no es valida, solo se acepta S o n")
		estiloError.Println("Vuelve a intentarlo")
		time.Sleep(3 * time.Second)
		mostrarPantalla(p)
	}
}
