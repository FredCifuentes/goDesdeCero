package usuarios

import (
	"fmt"
	"time"

	"github.com/fredcifuentes/goDesdeCero/modelos"
)

func AltaUsuario() {
	Usuario := new(modelos.Usuario)
	Usuario.AgregarUsuario(1, time.Now(), true, "fredy", "cifuentes", "robledo", "17-04-1997", 26)
	fmt.Println(Usuario)
}
