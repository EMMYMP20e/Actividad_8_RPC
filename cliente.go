package main

import (
	"bufio"
	"fmt"
	"net/rpc"
	"os"
	"os/exec"
)

type Datos struct {
	Nombre       string
	Materia      string
	Calificacion float64
}

func main() {
	c, err := rpc.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println(err)
		return
	}
	var input string
	opc := 1
	for opc != 0 {
		opc = menu()
		switch opc {
		case 1:
			var nombre, materia string
			var calificacion float64
			fmt.Scanln(&input)
			fmt.Print("Nombre del Alumno: ")
			fmt.Scanln(&nombre)
			fmt.Print("Materia: ")
			fmt.Scanln(&materia)
			fmt.Print("Calificación: ")
			fmt.Scan(&calificacion)
			fmt.Scanln(&input)
			data := Datos{
				Nombre:       nombre,
				Materia:      materia,
				Calificacion: calificacion,
			}
			var reply bool
			err = c.Call("Servidor.AgregarDatos", data, &reply)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Datos agregados")
			}
			pausa()
		case 2:
			var nombre string
			fmt.Scanln(&input)
			fmt.Print("Nombre de alumno:")
			fmt.Scanln(&nombre)
			var reply float64
			err = c.Call("Servidor.PromedioAlumno", nombre, &reply)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Promedio de Alumno: ", reply)
			}
			pausa()
		case 3:
			var reply float64
			fmt.Scanln(&input)
			err = c.Call("Servidor.PromedioGeneral", 0, &reply)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Promedio General: ", reply)
			}
			pausa()
		case 4:
			var materia string
			fmt.Scanln(&input)
			fmt.Print("Nombre de la materia:")
			fmt.Scanln(&materia)
			var reply float64
			err = c.Call("Servidor.PromedioMateria", materia, &reply)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println("Promedio Materia: ", reply)
			}
			pausa()
		}
	}
}

func pausa() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("[Presiona Enter para Continuar]")
	scanner.Scan()
}

func menu() int {
	var opc int
	limpiarPantalla()
	fmt.Println("1.- Agregar Calificación")
	fmt.Println("2.- Mostrar promedio de un alumno")
	fmt.Println("3.- Mostrar promedio general")
	fmt.Println("4.- Mostrar promedio de una materia")
	fmt.Println("0.- Salir")
	fmt.Print("Opción: ")
	fmt.Scan(&opc)
	limpiarPantalla()
	return opc
}

func limpiarPantalla() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
