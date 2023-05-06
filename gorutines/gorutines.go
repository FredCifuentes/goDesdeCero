package gorutines

import (
	"fmt"
	"strings"
	"time"
)

func MinombreLento(nombre string,canal1 chan bool) {
	letras := strings.Split(nombre, "")

	for _, v := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(v)

	}

	canal1<- true

}
