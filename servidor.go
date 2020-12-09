package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
)

type Datos struct {
	Nombre       string
	Materia      string
	Calificacion float64
}

var materias = make(map[string]map[string]float64)
var alumnos = make(map[string]map[string]float64)

type Servidor struct{}

func (this *Servidor) AgregarDatos(d Datos, reply *bool) error {
	alumno := make(map[string]float64)
	alumno[d.Nombre] = d.Calificacion
	if el, ok := materias[d.Materia]; ok {
		if _, ok := el[d.Nombre]; ok {
			return errors.New("Calificación ya existente")
		} else {
			materias[d.Materia][d.Nombre] = d.Calificacion
		}
	} else {
		materias[d.Materia] = alumno
	}
	materia := make(map[string]float64)
	materia[d.Materia] = d.Calificacion
	if _, ok := alumnos[d.Nombre]; ok {
		alumnos[d.Nombre][d.Materia] = d.Calificacion
	} else {
		alumnos[d.Nombre] = materia
	}
	*reply = true
	return nil
}

func (this *Servidor) PromedioAlumno(nombre string, reply *float64) error {
	cont := 0.0
	suma := 0.0
	for _, calificacion := range alumnos[nombre] {
		suma = suma + calificacion
		cont++
	}
	promedio := suma / cont
	*reply = promedio
	return nil
}
func (this *Servidor) PromedioMateria(materia string, reply *float64) error {
	cont := 0.0
	suma := 0.0
	for _, calificacion := range materias[materia] {
		suma = suma + calificacion
		cont++
	}
	promedio := suma / cont
	*reply = promedio
	return nil
}
func (this *Servidor) PromedioGeneral(c int, reply *float64) error {
	suma := 0.0
	cont := 0.0
	for _, alumno := range materias {
		for _, calificacion := range alumno {
			suma = suma + calificacion
			cont++
		}
	}
	promedio := suma / cont
	*reply = promedio
	return nil
}
func server() {
	rpc.Register(new(Servidor))
	ln, err := net.Listen("tcp", ":9999")
	if err != nil {
		fmt.Println(err)
	}
	for {
		c, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go rpc.ServeConn(c)
	}
}

func main() {
	go server()

	var input string
	fmt.Scanln(&input)
}
