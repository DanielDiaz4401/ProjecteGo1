// Package entrada inclou els recursos
// per crear i gestionar entrades
// de forma independent
package entrada

import (
	"strings"

	"github.com/DanielDiaz4401/ProjecteGo1/domini/temps"
)

// Entrada gestiona una Entrada del Blog.
// Una Entrada té un id, un titol, un text,
// i el dia i hora en que s'ha creat.
type Entrada struct {
	id    int
	titol string
	text  string
	data  temps.Temps
}

var nEntrades int = 0

// New crea una entrada amb el titol, el text,
// i el dia i hora indicats i incrementa en
// 1 el valor del atribut static nEntrades.
func New(titol, text, data string) Entrada {
	nEntrades++
	return Entrada{nEntrades - 1, titol, text, temps.New(data)}
}

// GetID retorna la ID de l'entrada
func (e Entrada) GetID() int {
	return e.id
}

// GetTitol retorna el titol de l'entrada
func (e Entrada) GetTitol() string {
	return e.titol
}

// GetQuan retorna la data de l'entrada
func (e Entrada) GetQuan() temps.Temps {
	return e.data
}

// Compare compara les dues entrades segons
// el criteri, que pot ser "num" o "titol"
func Compare(una, altra Entrada, criteri string) int {
	if criteri == "titol" {
		return strings.Compare(una.titol, altra.titol)
	}
	if criteri == "num" {
		return una.id - altra.id
	}
	return 0
}

// ToString retorna un string en el format
// "id titol text dia hora"
func (e Entrada) ToString() string {
	return string(e.id) + e.titol + e.text + e.data.Get()
}
