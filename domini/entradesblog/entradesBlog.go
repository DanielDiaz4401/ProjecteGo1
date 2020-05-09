// Package entradesblog gestiona
// tota la llista d'entrades
// de l'usuari i estableix un
// ordre segons el criteri
package entradesblog

import (
	"errors"
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
type EntradesBlog []Entrada

var criteri string

// New crea una llista buida d'entrades
func New() EntradesBlog {
	return []Entrada{}
}

// Agafa l'entrada amb id==num.
// Si existeix, la retorna.
// Si no existeix, retorna err
func (e *EntradesBlog) Agafa(num int) (Entrada, error) {
	for _, entrada := range *e {
		if entrada.GetID() == num {
			return entrada, nil
		}
	}
	return Entrada{}, errors.New("No existeix l'entrada")
}

// Elimina l'entrada amb id==num.
// Si existeix, la retorna.
// Si no existeix, retorna err
func (e *EntradesBlog) Elimina(num int) (Entrada, error) {
	for i, entrada := range *e {
		if entrada.GetID() == num {
			*e = remove(*e, i)
			return entrada, nil
		}
	}
	return Entrada{}, errors.New("No existeix l'entrada")
}

// Ejemplo de polimorfismo en go
func remove(slice []Entrada, s int) []Entrada {
	return append(slice[:s], slice[s+1:]...)
}

// EntreDates retorna un string amb les
// compreses entre les dues dates.
// "Dates{
//		entrada
// }"
func (e *EntradesBlog) EntreDates(inici, fi temps.Temps) string {
	aux := "Dates{\n"
	for _, entrada := range *e {
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
	for _, entrada := range *e {
		aux += "\t" + strconv.Itoa(int(entrada.GetID())) + " " + entrada.GetTitol() + "\n"
	}
	return aux + "}"
}

// ToString retorna un string amb les
// entrades ordenades.
// "Totes{
//		entrada
// }"
func (e *EntradesBlog) ToString() string {
	aux := "Totes{\n"
	for _, entrada := range *e {
		aux += "\t" + strconv.Itoa(int(entrada.GetID())) + " " + entrada.GetTitol() + entrada.GetQuan().Get() + "\n"
	}
	return aux + "}"
}
