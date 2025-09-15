package main

import (
	"code/src"
)
//faut que je creer une struct tydef Charactere ou je met le nom etc pour pouvoir le rappeller plus facilement
//dans chaque func avec diff player if i wanna change of stat for a player 
//genre a finir mais ::}

func main() {
	c1 := code.InitCharacter(
		"Jean-Michel CV", // Nom
		"Elfe",           // Classe
		1,                // Niveau
		100,              // PV max
		40,               // PV actuel
		[]string{}, // Inventaire
		180,
	)
	code.Menu(&c1)
}