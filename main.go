package main

import (
	"fmt"

	"github.com/fredcifuentes/goDesdeCero/gorutines"
)

func main() {
	/*ejercicios.SolicitarNumero()

	iteraciones.IteracionesFor()

		teclado.RecibirDatosEntrada()

		numero, cadena := ejercicios.CoversionEntero("120")
		fmt.Println(numero)
		fmt.Println(cadena)*/

	/*Estado, Nombre := variables.ConvertirTexto(1200)
	fmt.Println(Estado)
	fmt.Println(Nombre)

	if os := runtime.GOOS; os == "Linux." {
		fmt.Println("es linux")
	} else {
		fmt.Println("Es windows")
	}

	switch os := runtime.GOOS; os {
	case "Linux.":
		fmt.Println("es Linux")
	default:
		fmt.Println("Es Windows en el swithc")

	}*/

	//files.SumarTabla()

	//files.LeoArchivo()
	//funciones.Calculos()

	//funciones.Exponencia(2)

	//arreglos.MuestroArrelos()
	/*Juan := new(modelos.Hombre)
	Juana := new(modelos.Mujer)
	implinterfaces.HumanosRespirando(Juan)

	implinterfaces.HumanosRespirando(Juana)*/

	//difer.EjemploPanic()
	canal1 := make(chan bool)
	go gorutines.MinombreLento("Fredy", canal1)

	fmt.Println("esoty aqui")
	<-canal1

}
