package code

import (
	"bufio"
	"fmt"
	"strings"
)

func CheckMaxItem(c *Character) bool {
	return len(c.Inventory) < 5
}

func AddSkill(c *Character, item string) {
	for i, l := range c.Skill {
		if l == item {
			fmt.Println("\033[31m❌ Choix invalide. Réessayez !\033[0m")
		}
		i++
	}
	c.Skill = append(c.Skill, item)
}

func AddInventory(c *Character, item string) {
	item = strings.TrimSpace(item)
	if CheckMaxItem(c) {
		c.Inventory = append(c.Inventory, item)
		fmt.Printf("\033[32m✅ %q ajouté à votre inventaire !\033[0m\n", item)
	} else {
		fmt.Println("\033[31m❌ Inventaire plein ! Vous ne pouvez pas avoir plus de 5 items.\033[0m")
	}
}

func AddInventoryUpgraded1(c *Character, item string) {
	item = strings.TrimSpace(item)
	if len(c.Inventory) < 8 {
		c.Inventory = append(c.Inventory, item)
		fmt.Printf("\033[32m✅ %q ajouté à votre inventaire !\033[0m\n", item)
	} else {
		fmt.Println("\033[31m❌ Inventaire plein ! Vous ne pouvez pas avoir plus de 8 items.\033[0m")
	}
}
func RemoveInventory(c *Character, index int) {
	if index >= 0 && index < len(c.Inventory) {
		c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
	}
}

func ShowInventory(c *Character) {
	if len(c.Inventory) == 0 {
		fmt.Println("Inventaire vide.")
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
		fmt.Println("\n\033[32m[P]\033[0m 💊 Oasis")
		fmt.Println("\033[35m[O]\033[0m ☠️  Café réchauffé")
		fmt.Println("\033[35m[E]\033[0m 🛡️  Équiper un habit")
		fmt.Println("\033[31m[R]\033[0m ↩️  Retour")
		fmt.Println("\033[36m══════════════════════════════════════\033[0m")
		fmt.Print("\033[1;34mVotre choix : \033[0m")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "p" {
			GoodPot(c)
		} else if input == "o" {
			PoisonPot(c)
		} else if input == "e" {
			EquipMenu(c, reader)
		} else if input == "d" {
			Skill(c)
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
		fmt.Printf("✨ XP Disponible : %d\n", c.Experience)
		fmt.Println("\033[32m[1]\033[0m 💊 Oasis (coût : 30 XP)")
		fmt.Println("\033[35m[2]\033[0m ☠️ Café réchauffé (bonus : 20 XP)")
		fmt.Println("\033[35m[3]\033[0m ☠️ Formation skill : Diplômatie (coût : 20 XP, bonus : 10 CompDay)")
		fmt.Println("\033[33m[4]\033[0m 👜 Plus grosse sacoche (coût : 50 XP)")
		fmt.Println("\033[31m[R]\033[0m ↩️  Retour")
		fmt.Println("\033[36m══════════════════════════════════════\033[0m")
		fmt.Print("\033[1;34mVotre choix : \033[0m")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "1":
			if c.Experience >= 30 {
				c.Experience -= 30
				AddInventory(c, "Oasis")
				fmt.Println("\033[32m✅ Vous avez bu un oasis !\033[0m")
			} else {
				fmt.Println("\033[31m❌ Pas assez d'XP !\033[0m")
			}
		case "2":
			if c.Experience >= 50 {
				c.Experience += 20
				AddInventory(c, "Café réchaufé")
				fmt.Println("\033[32m✅ Vous avez bu un café réchauffé !\033[0m")
			} else {
				fmt.Println("\033[31m❌ Pas assez d'XP !\033[0m")
			}
		case "3":
			if c.Experience >= 20 {
				c.Experience -= 20
				c.CurrentCompDay += 10
				AddSkill(c, "diplomatie")
				fmt.Println("\033[32m✅ Vous êtes plus calme !\033[0m")
			} else {
				fmt.Println("\033[31m❌ Pas assez d'XP !\033[0m")
			}
		case "4":
			if c.Experience >= 50 {
				c.Experience -= 50
				UpgradeInventorySlot(c)
				fmt.Println("\033[32m✅ Votre inventaire à augmenté de taille !\033[0m")
			} else {
				fmt.Println("\033[31m❌ Pas assez d'XP !\033[0m")
			}
		case "r":
			return
		default:
			fmt.Println("\033[31m❌ Choix invalide. Réessayez !\033[0m")
		}
	}
}
func UpgradeInventorySlot(c *Character) {
	// Placeholder for inventory upgrade functionality
	fmt.Println("Inventaire amélioré !")
}
