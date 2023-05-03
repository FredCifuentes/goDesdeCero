package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var err error
var texto string

func SolicitarNumero() string {

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
		texto += fmt.Sprintf("%d x %d = %d\n", numero1, i, numero1*i)
	}
	//fmt.Println(texto)
	return texto

}
