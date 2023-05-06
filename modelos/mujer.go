package modelos

type Mujer struct {
	Hombre
}

func (m *Mujer) Respirar() {
	m.Respirando = true
}

func (m *Mujer) Cantar() {
	m.Cantando = true
}
func (m *Mujer) Pensar() {
	m.Pensando = true
}
func (m *Mujer) Sexo() string {
	return "Mujer"
}
