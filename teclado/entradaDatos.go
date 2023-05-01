package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var mensaje string
var err error

func RecibirDatosEntrada() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Recibiendo Datos desde teclado numero 1")

	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}
		fmt.Println("este es el valor", numero1)
	}

	scanner2 := bufio.NewScanner(os.Stdin)
	fmt.Println("Recibiendo Datos desde teclado numero 2")

	if scanner2.Scan() {
		numero2, err = strconv.Atoi(scanner2.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}
		fmt.Println("este es el valor", numero2)
	}

	fmt.Println("Recibiendo mensaje ")
	if scanner.Scan() {
		mensaje = scanner.Text()
		fmt.Println(mensaje)
	}

	fmt.Println(mensaje, numero1*numero2)

}
