// Package interficieusuari gestiona el menú principal i invoca al
// controlador per implementar cada una de les opcions.
package interficieusuari

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/DanielDiaz4401/ProjecteGo1/controlador"
)

// InterficieUsuari pot "heredar"
// de controlador en la primera versió
type InterficieUsuari struct {
	controlador.Controlador
}

// New crea la Interficie d'usuari al main
func New() InterficieUsuari {
	return InterficieUsuari{controlador.New()}
}

// Cicle principal d'execució:
// mostra opcions, llegeix opcio
// i executa opcio.
func (i *InterficieUsuari) Cicle() {
	reader := bufio.NewReader(os.Stdin)
	i.MostraOpcions()
	op, _ := reader.ReadString('\n')
	op = strings.ReplaceAll(op, "\n", "")
	i.ExecutaOpcio(op)
	for strings.ToLower(string(op[0])) != "f" {
		i.MostraOpcions()
		op, _ = reader.ReadString('\n')
		op = strings.ReplaceAll(op, "\n", "")
		i.ExecutaOpcio(op)
	}
}

// ExecutaOpcio cada valor d'opcio
func (i *InterficieUsuari) ExecutaOpcio(op string) {
	switch string(op[0]) {
	case "i":
		i.Index()
	case "m":
		aux, err := buscaNumero(op)
		if err != nil {
			fmt.Printf("S'ha produit un error: %v\n", err)
			return
		}
		i.MostraEntrada(aux)
	case "n":
		i.NovaEntrada()
	case "e":
		aux, err := buscaNumero(op)
		if err != nil {
			fmt.Printf("S'ha produit un error: %v\n", err)
			return
		}
		i.EliminaEntrada(aux)
	case "o":
		i.Ordena()
	case "d":
		i.Dates()
	case "t":
		i.MostraEntrades()
	case "f":
		println("Adeu!")
	}
}

func buscaNumero(frase string) (int, error) {
	aux := strings.Split(frase, " ")
	if len(aux) != 2 {
		return 0, errors.New("Error de format")
	}
	toParse := strings.Split(aux[1], "\n")
	return strconv.Atoi(toParse[0])
}

// MostraOpcions del menú principal.
func (i *InterficieUsuari) MostraOpcions() {
	fmt.Println("Opcions:\n" +
		"\tindex\n" +
		"\tmostra <num>\n" +
		"\tnova\n" +
		"\telimina <num>\n" +
		"\tordena\n" +
		"\tdates\n" +
		"\ttotes\n" +
		"\tfi\n" +
		"Introdueix una opció:\n")
}
