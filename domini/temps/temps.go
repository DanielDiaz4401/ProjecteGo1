// Package temps controla el
//moment en el que es va
//crear l'entrada
package temps

import (
	"strings"
	"time"
)

// Algo es algo
type Algo interface {
	Get() string
	ComparaTemps(o Temps) int
}

// Temps esta expressat
// en hores (hh:mm:dd)
// i en dia (aaaa-mm-dd)
type Temps struct {
	dia  string
	hora string
}

// New crea un nou Temps.
// Hi ha tres opcions:
// 	-Posar el dia i l'hora.
//  -Posar només el dia
//  -No posar res i asignar el
//	dia actual.
func New(data []string) Temps {
	switch len(data) {
	case 0:
		return ara()
	case 1:
		return Temps{dia: data[0]}
	default:
		return Temps{data[0], data[1]}
	}
}

// ara retorna el temps actual en
// tipus Temps
func ara() Temps {
	aux := time.Now()
	total := aux.Format("2006-01-02 15:04:05")
	vector := strings.Split(total, " ")
	dia := vector[0]
	hora := vector[1]
	return Temps{dia, hora}
}

// Get retorna un string del temps amb
// format aaaa-dd-mm hh:mm:ss
func (t *Temps) Get() string {
	return t.dia + " " + t.hora
}

func (t *Temps) comparaDia(o Temps) int {
	return strings.Compare(t.dia, o.dia)
}

// ComparaTemps compara els temps
func (t *Temps) ComparaTemps(o Temps) int {
	dia := t.comparaDia(o)
	if dia == 0 {
		return strings.Compare(t.hora, o.hora)
	}
	return dia
}
