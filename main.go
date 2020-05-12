// Package main es crea perqué Go només
// executa de packages amb aquest nom
package main

import (
	"github.com/DanielDiaz4401/ProjecteGo1/iu/interficieusuari"
)

func main() {
	iu := interficieusuari.New()
	iu.Cicle()
}
