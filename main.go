package main

import (
	"fmt"

	"github.com/fredcifuentes/goDesdeCero/variables"
)

func main() {
	Estado, Nombre := variables.ConvertirTexto(1200)
	fmt.Println(Estado)
	fmt.Println(Nombre)
}
