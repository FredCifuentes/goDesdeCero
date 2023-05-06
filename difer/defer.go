package difer

import (
	"fmt"
	"log"
)

func UsandoDefer() {
	fmt.Println("mensaje 1")

	defer fmt.Println("dentro de defer")

	fmt.Println("mensaje 2")
}

func EjemploPanic() {
	defer func ()  {
		reco:=recover()
		if reco!= nil {
	log.Fatalf("Ocurrio un error  que genero panic y se aborto el programa\n %v",reco)
		}
	}()
	a := 1

	if a == 1 {
		panic("Se encontro un error")
	}
}
