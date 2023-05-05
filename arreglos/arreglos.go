package arreglos

import "fmt"

var array [10]int

var array2 [3]int = [3]int{5, 7, 9}

var matriz [2][2]int

func MuestroArrelos() {

	fmt.Println("LLenando arrglos")
	for i := 0; i < 10; i++ {
		array[i] = i
	}
	fmt.Println("Imprmiendo arreglos")
	for i := 0; i < 10; i++ {
		fmt.Println(array[i])
	}
	tabla2 := [6]string{"fredy", "Daniel"}

	fmt.Println(tabla2)
}
