// Package entradesblog gestiona
// tota la llista d'entrades
// de l'usuari i estableix un
// ordre segons el criteri
package entradesblog

import (
	"errors"
	"sort"
	"strconv"

	"github.com/DanielDiaz4401/ProjecteGo1/domini/entrada"
	"github.com/DanielDiaz4401/ProjecteGo1/domini/temps"
)

// Entrada és un alies d'entrada
// per poder manipular les entrades
// en els métodes que ho requereixin
type Entrada = entrada.Entrada

// EntradesBlog es d'entrades
// de l'usuari
type EntradesBlog struct {
	entrades []entrada.Entrada
}

var criteri string

// New crea una llista buida d'entrades
func New() EntradesBlog {
	return EntradesBlog{[]entrada.Entrada{}}
}

// AfegeixOrdenat afegeix una Entrada e a la posició
// que li correspon a la llista d'entrades,
// segons el criteri actual.
func (e *EntradesBlog) AfegeixOrdenat(in entrada.Entrada) {
	for i, aux := range e.entrades {
		if i < len(e.entrades) {
			if entrada.Compare(in, aux, criteri) < 0 {
				e.entrades = add(e.entrades, i, in)
				return
			}
		}
	}
	e.entrades = append(e.entrades, in)
}

func add(slice []entrada.Entrada, i int, afegir entrada.Entrada) []entrada.Entrada {
	slice = append(slice, afegir) // Augmento la capacitat ce slice
	copy(slice[i+1:], slice[i:])  // Copio el que hi ha a i una posició a la dreta
	slice[i] = afegir             // Ara que tinc un lloc lliure, afegeixo en aquesta posicio l'entrada
	return slice
}

// Agafa l'entrada amb id==num.
// Si existeix, la retorna.
// Si no existeix, retorna err
func (e *EntradesBlog) Agafa(num int) (entrada.Entrada, error) {
	for _, entrada := range e.entrades {
		if entrada.GetID() == num {
			return entrada, nil
		}
	}
	return entrada.Entrada{}, errors.New("No existeix l'entrada")
}

// Elimina l'entrada amb id==num.
// Si existeix, la retorna.
// Si no existeix, retorna err
func (e *EntradesBlog) Elimina(num int) (entrada.Entrada, error) {
	for i, entrada := range e.entrades {
		if entrada.GetID() == num {
			e.entrades = remove(e.entrades, i)
			return entrada, nil
		}
	}
	return entrada.Entrada{}, errors.New("No existeix l'entrada")
}

// Ejemplo de polimorfismo en go
func remove(slice []entrada.Entrada, i int) []entrada.Entrada {
	return append(slice[:i], slice[i+1:]...)
}

// EntreDates retorna un string amb les
// compreses entre les dues dates.
// "Dates{
//		entrada
// }"find . -name '*.php' | xargs wc -l
func (e *EntradesBlog) EntreDates(inici, fi temps.Temps) string {
	aux := "Dates{\n"
	for _, entrada := range e.entrades {
		if entrada.GetQuan().ComparaTemps(inici) > 0 {
			if entrada.GetQuan().ComparaTemps(fi) < 0 {
				aux += "\t" + strconv.Itoa(int(entrada.GetID())) + " " + entrada.GetTitol() + "\n"
			}
		}
	}
	return aux + "}"
}

// Index retorna un string amb les
// entrades ordenades.
// "Index{
//		id titol
// }"
func (e *EntradesBlog) Index() string {
	aux := "Index{\n"
	for _, entrada := range e.entrades {
		aux += "\t" + strconv.Itoa(int(entrada.GetID())) + " " + entrada.GetTitol() + "\n"
	}
	return aux + "}"
}

// Ordena les entrades segons el criteri especificatfind . -name '*.php' | xargs wc -lfind . -name '*.php' | xargs wc -lfind . -name '*.php' | xargs wc -lfind . -name '*.php' | xargs wc -lfind . -name '*.php' | xargs wc -l
func (e *EntradesBlog) Ordena() {
	if criteri == "num" {
		sort.Slice(e.entrades, func(i, j int) bool {
			return e.entrades[i].GetID() < e.entrades[j].GetID()
		})
	} else {
		sort.Slice(e.entrades, func(i, j int) bool {
			return e.entrades[i].GetTitol() < e.entrades[j].GetTitol()
		})
	}
}

// SetCriteri estableix el criteri per ordenar
// en "tit" (titol) o "num" (id). Retorna error
// si el string no correspon amb cap opció
func (e *EntradesBlog) SetCriteri(crit string) error {
	switch crit {
	case "titol":
		criteri = "titol"
		return nil
	case "num":
		criteri = "num"
		return nil
	default:
		return errors.New("Error al establir el criteri")
	}
}

// ToString retorna un string amb les
// entrades ordenades.
// "Totes{
//		entrada
// }"
func (e *EntradesBlog) ToString() string {
	aux := "Totes{\n"
	for _, entrada := range e.entrades {
		aux += "\t" + strconv.Itoa(int(entrada.GetID())) + " " + entrada.GetTitol() + " " + entrada.GetQuan().Get() + "\n"
	}
	return aux + "}"
}
