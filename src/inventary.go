package code

import (
	"fmt"
	"bufio"
	"strings"
)

func AddInventory(c *Character, item string) {
	item = strings.TrimSpace(item)
	c.Inventory = append(c.Inventory, item)
}

func RemoveInventory(c *Character, index int) {
	if index >= 0 && index < len(c.Inventory) {
		c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
	}
}

func ShowInventory(c *Character) {
	if len(c.Inventory) == 0 {
		fmt.Println("Inventory vide.")
	} else {
		for i, item := range c.Inventory {
			fmt.Printf("[%d] %q\n", i+1, item)
		}
	}
}

func AccessInventory(c *Character, reader *bufio.Reader) {
	for {
		fmt.Println("\033[36m══════════════════════════════════════\033[0m")
		fmt.Println("\033[1;33m           🎒 INVENTAIRE 🎒           \033[0m")
		fmt.Println("\033[36m══════════════════════════════════════\033[0m")
		ShowInventory(c)
		fmt.Println("\n\033[32m[P]\033[0m 💊 Utiliser une potion de vie")
		fmt.Println("\033[35m[O]\033[0m ☠️  Utiliser une potion de poison")
		fmt.Println("\033[31m[R]\033[0m ↩️  Retour")
		fmt.Println("\033[36m══════════════════════════════════════\033[0m")
		fmt.Print("\033[1;34mVotre choix : \033[0m")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "p" {
			TakePot(c)
		} else if input == "r" {
			return
		} else {
			fmt.Println("\033[31m❌ Choix invalide. Réessayez !\033[0m")
		}
	}
}

func AccesMerchant(c *Character, reader *bufio.Reader) {
	for {
		fmt.Println("\033[36m══════════════════════════════════════\033[0m")
		fmt.Println("\033[1;33m         🏪 BIENVENUE À LA CAF 🏪        \033[0m")
		fmt.Println("\033[36m══════════════════════════════════════\033[0m")
		fmt.Println("\033[32m[1]\033[0m 💊 Potion de vie (gratuit)")
		fmt.Println("\033[35m[2]\033[0m ☠️  Potion de poison (gratuit)")
		fmt.Println("\033[31m[R]\033[0m ↩️  Retour")
		fmt.Println("\033[36m══════════════════════════════════════\033[0m")
		fmt.Print("\033[1;34mVotre choix : \033[0m")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "1":
			AddInventory(c, "Potion")
			fmt.Println("\033[32m✅ Vous avez reçu une potion de vie !\033[0m")
		case "r":
			return
		default:
			fmt.Println("\033[31m❌ Choix invalide. Réessayez !\033[0m")
		}
	}
}