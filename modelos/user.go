package modelos

import "time"

type Usuario struct {
	id               int
	CreatedAt        time.Time
	Status           bool
	Nombre           string
	Apellido_p       string
	Apellido_m       string
	Fecha_nacimiento string
	Edad             int
}

func (this *Usuario) AgregarUsuario(id int, createdAt time.Time, staus bool, nombre string, apellidoP string, apellidoM string, fechaNacimiento string, edad int) {
	this.id = id
	this.CreatedAt = createdAt
	this.Status = staus
	this.Nombre = nombre
	this.Apellido_m = apellidoM
	this.Apellido_p = apellidoP
	this.Fecha_nacimiento = fechaNacimiento
	this.Edad = edad
}
