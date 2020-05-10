// Package controlador s'encarrega
// de controlar les entrades ddel blog
// i implementa les opcions del menú de
// l'usuari.
package controlador

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/DanielDiaz4401/ProjecteGo1/domini/entrada"
	"github.com/DanielDiaz4401/ProjecteGo1/domini/entradesblog"
	"github.com/DanielDiaz4401/ProjecteGo1/domini/temps"
)

//Controlador controla el blog
type Controlador struct {
	entrades entradesblog.EntradesBlog
}

// New crea un nou controlador
func New() Controlador {
	return Controlador{entradesblog.New()}
}

// Dates demana a l'usuari un dia
// inicial i un dia final per mostrar
// les entrades entre aquests dos dies
func (c *Controlador) Dates() {
	reader := bufio.NewReader(os.Stdin)
	var inici, fi string
	fmt.Println("Dia inicial?")
	inici, _ = reader.ReadString('\n')
	inici = strings.ReplaceAll(inici, "\n", "")
	fmt.Println("Dia final?")
	fi, _ = reader.ReadString('\n')
	fi = strings.ReplaceAll(fi, "\n", "")
	fmt.Println(c.entrades.EntreDates(temps.New(inici), temps.New(fi)))
}

// EliminaEntrada elimina l'entrada amb
// identificador igual al valor del argument
// num, i mostra l'entrada eliminada.
func (c *Controlador) EliminaEntrada(num int) {
	c.entrades.Elimina(num)
}

// Index mostra l'index de les entrades del blog.
func (c *Controlador) Index() {
	fmt.Println(c.entrades.Index())
}

// MostraEntrada l'entrada amb id==num
func (c *Controlador) MostraEntrada(num int) {
	entrada, err := c.entrades.Agafa(num)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Println(entrada.ToString())
}

// MostraEntrades mostra totes les entrades de
// l'usuari.
func (c *Controlador) MostraEntrades() {
	fmt.Println(c.entrades.ToString())
}

//NovaEntrada Demana el titol, el text, i el
// dia i hora a l'usuari, crea una entrada nova,
// i l'afegeix al blog en la posició que li
// correspon segons el criteri d'ordenació actual.
func (c *Controlador) NovaEntrada() {
	reader := bufio.NewReader(os.Stdin)
	var titol, text, data string
	fmt.Println("Titol?")
	titol, _ = reader.ReadString('\n')
	titol = strings.ReplaceAll(titol, "\n", "")
	fmt.Println("Text?")
	text, _ = reader.ReadString('\n')
	text = strings.ReplaceAll(text, "\n", "")
	fmt.Println("Data? (intro per assignar data actual)")
	data, _ = reader.ReadString('\n')
	data = strings.ReplaceAll(data, "\n", "")
	c.entrades.AfegeixOrdenat(entrada.New(titol, text, data))
}

// Ordena pregunta pel criteri d'ordenació
// i ordena les entrades seguint aquest criteri
func (c *Controlador) Ordena() {
	reader := bufio.NewReader(os.Stdin)
	var criteri string
	fmt.Println("Criteri? (num/titol)")
	criteri, _ = reader.ReadString('\n')
	criteri = strings.ReplaceAll(criteri, "\n", "")
	err := c.entrades.SetCriteri(criteri)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	c.entrades.Ordena()
}
