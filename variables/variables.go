package variables

import "fmt"

func MuestroEnteros() {
	var entero int
	entero32 := int32(10)
	entero64 := int64(100)

	fmt.Println("valor de entero", entero)
	fmt.Println("valor de int 32", entero32)
	fmt.Println("valor de int 64", entero64)
}
