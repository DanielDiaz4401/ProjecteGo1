// Package entradesblog gestiona
// tota la llista d'entrades
// de l'usuari i estableix un
// ordre segons el criteri
package entradesblog

import(
	"github.com/DanielDiaz4401/ProjecteGo1/domini/entrada"
)

// Llista d'entrades de l'usuari
type EntradesBlog []Entrada

var criteri string

func New() EntradesBlog {
	return []Entrada{}
}

// Agafa l'entrada amb id==num i,
// si existeix, la retorna.
func (e *EntradesBlog) Agafa(num int) Entrada {
	for _,var := range e {
		if var.GetID()==num {
			return var
		}
	}
	return nil
}

// Elimina l'entrada amb id==num i,
// si existeix, la retorna
func (e *EntradesBlog) Elimina(num int) Entrada {
	for i, var := range e {
		if var.GetID()==num {
			remove(e,i)
			return var
		}
	}
	return nil
}

// EntreDates retorna un string amb les
// compreses entre les dues dates.
// "Dates{
//		entrada
// }"
func (e *EntradesBlog) EntreDates(inici,fi temps.Temps) string{
	aux := "Dates{\n"
	for _, var := range e {
		aux +="\t" var.GetID()+" "+var.GetTitol+"\n" 
	}
	return aux+"}"
}

// Index retorna un string amb les
// entrades ordenades.
// "Index{
//		id titol
// }"
func (e *EntradesBlog) Index() string{
	aux := "Index{\n"
	for _, var := range e {
		aux +="\t" var.GetID()+" "+var.GetTitol+"\n" 
	}
	return aux+"}"
}

// ToString retorna un string amb les
// entrades ordenades.
// "Totes{
//		entrada
// }"
func (e *EntradesBlog) ToString() string{
	var aux string
	aux := "Totes{\n"
	for _, var := range e {
		aux += var.GetID()+" "+var.GetTitol()+var.GetQuan()"\n" 
	}
	return aux+"}"
}