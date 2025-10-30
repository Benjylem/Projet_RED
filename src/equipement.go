package code

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

var recipes = map[string][]string{
	"Bob":                  {"Plume de Corbeau", "Cuir de Sanglier"},
	"Marcel":               {"Fourrure de Loup", "Fourrure de Loup", "Peau de Troll"},
	"Claquette chaussette": {"Fourrure de Loup", "Cuir de Sanglier"},
}

func ForgeMenu(c *Character) {
	for {
		fmt.Println("\033[36m══════════════════════════════════════\033[0m")
		fmt.Println("\033[1;33m 	🛠️ FORGERON\033[0m – Choisissez un équipement à fabriquer :")
		fmt.Println("\033[36m══════════════════════════════════════\033[0m")
		fmt.Printf("Il vous reste %d XP", c.CurrentCompDay)
		fmt.Println("\n[1] 👒 Bob (coût : 5 XP)")
		fmt.Println("[2] 👕 Marcel (coût : 5 XP)")
		fmt.Println("[3] 👟 Claquette chaussette (coût : 5 XP)")
		fmt.Println("[R] ↩️ Retour")
		fmt.Print("Votre choix : ")

		var choix string
		fmt.Scanln(&choix)

		switch choix {
		case "1":
			craftEquipment(c, "Bob")
		case "2":
			craftEquipment(c, "Marcel")
		case "3":
			craftEquipment(c, "Claquette chaussette")
		case "R", "r":
			return
		default:
			fmt.Println("❌ Choix invalide.")
		}
	}
}

func craftEquipment(c *Character, item string) {
	needed := recipes[item]
	for _, res := range needed {
		if !hasItem(c.Inventory, res) {
			fmt.Printf("❌ Il vous manque : %s\n", res)
			return
		}
	}
	if c.Experience < 150 {
		fmt.Println("❌ Pas assez d'expérience pour fabriquer cet objet (5 XP requis).")
		return
	}
	c.Experience -= 55
	for _, res := range needed {
		removeItem(&c.Inventory, res)
	}
	c.Inventory = append(c.Inventory, item)
	fmt.Printf("✅ Vous avez fabriqué : %s\n", item)
}

func hasItem(inv []string, item string) bool {
	for _, v := range inv {
		if v == item {
			return true
		}
	}
	return false
}

func removeItem(inv *[]string, item string) {
	for i, v := range *inv {
		if v == item {
			*inv = append((*inv)[:i], (*inv)[i+1:]...)
			return
		}
	}
}

func EquipMenu(c *Character, reader *bufio.Reader) {
	var equipables []string
	for _, item := range c.Inventory {
		trimmed := strings.TrimSpace(item)
		if trimmed == "Bob" ||
			trimmed == "Marcel" ||
			trimmed == "Claquette chaussette" {
			equipables = append(equipables, item)
		}
	}
	if len(equipables) == 0 {
		fmt.Println("❌ Aucun équipement à équiper.")
		return
	}
	fmt.Println("\n\033[36m══════════════════════════════════════\033[0m")
	fmt.Println("\033[1;33m       🛡️  ÉQUIPER UN HABIT 🛡️\033[0m")
	fmt.Println("\033[36m══════════════════════════════════════\033[0m")
	for i, item := range equipables {
		fmt.Printf("[%d] %s\n", i+1, item)
	}
	fmt.Print("\033[32m Choisissez un équipement :\033[0m")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	choix, err := strconv.Atoi(input)
	if err != nil || choix < 1 || choix > len(equipables) {
		fmt.Println("\033[31m❌ Choix invalide.\033[0m")
		return
	}
	EquipItem(c, equipables[choix-1])
}

func EquipItem(c *Character, item string) {
	switch item {
	case "Bob":
		if c.Equipment.EquipmentHead != "" {
			fmt.Printf("🗑️ %s a été détruit.\n", c.Equipment.EquipmentHead)
		}
		c.Equipment.EquipmentHead = item
		c.MaxCompDay += 10

	case "Marcel":
		if c.Equipment.ChestEquipment != "" {
			fmt.Printf("🗑️ %s a été détruit.\n", c.Equipment.ChestEquipment)
		}
		c.Equipment.ChestEquipment = item
		c.MaxCompDay += 25

	case "Claquette chaussette":
		if c.Equipment.FootEquipment != "" {
			fmt.Printf("🗑️ %s a été détruit.\n", c.Equipment.FootEquipment)
		}
		c.Equipment.FootEquipment = item
		c.MaxCompDay += 15
	}
	removeItem(&c.Inventory, item)
	fmt.Printf("✅ %s équipé ! HP max = %d\n", item, c.MaxCompDay)
}
