// Package entradesblog gestiona
// tota la llista d'entrades
// de l'usuari i estableix un
// ordre segons el criteri
package entradesblog

// Llista d'entrades de l'usuari
type EntradesBlog []Entrada

var criteri string

func New() EntradesBlog {
	return []Entrada{}
}

// Agafa l'entrada amb id==num i,
// si existeix, la retorna.
func (e *EntradesBlog) Agafa(num int) Entrada {
	for _, var := range e {
		if var.GetID()==num {
			return var
		}
		return nil
	}
}

// Elimina l'entrada amb id==num i,
// si existeix, la retorna
func (e *EntradesBlog) Elimina(num int) Entrada {
	for i, var := range e {
		
	}
}
