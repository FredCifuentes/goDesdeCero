package implinterfaces

import (
	"fmt"

	"github.com/fredcifuentes/goDesdeCero/interfaces"
)

func HumanosRespirando(hu interfaces.Humano) {
	hu.Cantar()
	hu.Pensar()
	hu.Respirar()
	hu.Sexo()
	fmt.Printf("humano: %s\n", hu.Sexo())
}
