package modelos

type Hombre struct {
	Edad       int
	Altura     float32
	Peso       float32
	Respirando bool
	Pensando bool
	Cantando bool
	vivo bool
	


}

func (h *Hombre) Respirar(){
	h.Respirando=true
}

func (h *Hombre) Cantar(){
	h.Cantando=true
}
func (h *Hombre) Pensar(){
	h.Pensando=true
}
func (h *Hombre) Sexo()string{
	return "Hombre"
}
