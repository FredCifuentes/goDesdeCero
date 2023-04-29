package main

import (
	"fmt"
	"runtime"

	"github.com/fredcifuentes/goDesdeCero/variables"
)

func main() {
	Estado, Nombre := variables.ConvertirTexto(1200)
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

	}
}
