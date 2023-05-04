package funciones

import "fmt"

func usandoClosures(valor int) func() int {
	numero := valor
	secuencia := 0

	return func() int {
		secuencia++
		return numero * secuencia
	}
}

func Closures() {
	numero := 2

	mitabla := usandoClosures(numero)

	for i := 1; i <= 10; i++ {
		fmt.Println(mitabla())
	}
}
