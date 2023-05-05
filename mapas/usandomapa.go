package mapas

import "fmt"

func MostrarMapa() {
	paises := make(map[string]string)

	paises["Mexico"] = "DF"
	paises["Argentina"] = "la paz"

	campeones := map[string]int{"barcelona": 10, "real madrid": 12}

	fmt.Println(paises)
	fmt.Println(campeones)

	for equipo, puntaje := range campeones {
		fmt.Printf("Equipo %s puntaje %d \n", equipo, puntaje)
	}
}
