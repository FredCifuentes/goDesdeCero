package arreglos

import "fmt"

var tablaS []int = []int{1, 2, 3}

var arregloS [10]int = [10]int{1, 2, 6, 7, 4, 3, 4, 5, 2, 1}

func UsandoSlices() {
	fmt.Println(tablaS)

	porcion := arregloS[1:]
	porcion2 := arregloS[:6]

	fmt.Println(porcion)
	fmt.Println(porcion2)
}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("largo %d, capacidad %d", len(elementos), cap(elementos))

	nums := make([]int, 0, 0)

	for i := 0; i < 10; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("largo %d, capacidad %d", len(nums), cap(nums))

}
