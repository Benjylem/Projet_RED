package code

import (
	"fmt"
)

type Charactere struct {
	Nom	string
	Class      string
	Level      int
	MaxHP      int
	CurrentHP  int
	INVENTAIRE  []string
	DaysIndemnisation int
}

func InitCharacter(nom, class string, level, maxHP, currentHP int, inventaire []string,  daysIndem int) Charactere {
	return Charactere{
		Nom:      nom,
		Class:     class,
		Level:     level,
		MaxHP:     maxHP,
		CurrentHP: currentHP,
		INVENTAIRE: inventaire,
		DaysIndemnisation: daysIndem,
	}
}

func DisplayInfo(c *Charactere) {
	fmt.Println("\n===== FICHE PERSONNAGE =====")
	fmt.Printf("👤 Nom      : %s\n", c.Nom)
	fmt.Printf("🌍 Classe   : %s\n", c.Class)
	fmt.Printf("📈 Niveau   : %d\n", c.Level)
	fmt.Printf("❤️ PV       : %d/%d\n", c.CurrentHP, c.MaxHP)
	fmt.Printf("🎒 Inventaire: %v\n", c.INVENTAIRE)
	fmt.Println("============================")
}
