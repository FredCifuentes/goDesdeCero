package main

import (
	"fmt"

	"github.com/fredcifuentes/goDesdeCero/ejercicios"
)

func main() {

	numero, cadena := ejercicios.CoversionEntero("120")
	fmt.Println(numero)
	fmt.Println(cadena)

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

}
