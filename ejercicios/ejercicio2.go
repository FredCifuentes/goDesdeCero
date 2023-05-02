package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var err error

func SolicitarNumero() {

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Recibiendo Datos desde teclado numero 1")
	for {
		if scanner.Scan() {
			numero1, err = strconv.Atoi(scanner.Text())
			if err != nil {
				continue
			} else {
				break
			}
		}
	}

	fmt.Println("Tabla de Multiplicar")

	for i := 1; i <= 10; i++ {
		fmt.Println(numero1, "x", i, "=", numero1*i)
	}

}
