package files

import (
	//"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fredcifuentes/goDesdeCero/ejercicios"
)

var ruta string = "./files/txt/tabla.txt"

func CrearTabla() {
	texto := ejercicios.SolicitarNumero()

	archivo, err := os.Create(ruta)

	if err != nil {
		fmt.Println("error al crear un archivo", err)
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()
}

func SumarTabla() {
	texto := ejercicios.SolicitarNumero()
	if !Append(ruta, texto) {
		fmt.Println("Error al concatenar")
	}
}

func Append(ruta string, texto string) bool {
	archivo, err := os.OpenFile(ruta, os.O_WRONLY|os.O_APPEND, 0664)

	if err != nil {
		fmt.Println("Error durante el apend", err)
		return false
	}
	_, err = archivo.WriteString(texto)
	if err != nil {
		fmt.Println("Error al Write String ", err)
	}
	archivo.Close()
	return true
}

func LeoArchivo() {
	archivo, err := ioutil.ReadFile(ruta)
	if err != nil {
		fmt.Println("Error al Escribir en el archivo", err)
		return
	}

	fmt.Println(string(archivo))

}
