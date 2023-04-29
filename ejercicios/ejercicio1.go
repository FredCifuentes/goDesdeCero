package ejercicios

import (
	"strconv"
)

func CoversionEntero(valor string) (int, string) {

	numero := valor

	if sv, err := strconv.Atoi(numero); sv > 100 {
		if err != nil {
			return 0, "Hubo un error"
		}
		return sv, "Es mayor a 100"

	} else {
		return sv, "Es menor a 100"
	}

}
