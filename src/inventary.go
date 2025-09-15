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
		fmt.Println("\n=== Inventory ===")
		ShowInventory(c)
		fmt.Println("[P] Utiliser une potion")
		fmt.Println("[R] Retour")
		fmt.Print("Choix : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		if input == "p" {
			TakePot(c)
		} else if input == "r" {
			return
		} else {
			fmt.Println("Choix invalide.")
		}
	}
}

func AccesMerchant(c *Character, reader *bufio.Reader) {
	for {
		fmt.Println("\n=== Bienvenu à la CAF ===")
		fmt.Println("[1] Potion de vie (gratuit)")
		fmt.Println("[R] Retour")
		fmt.Println("Choice : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "1":
			AddInventory(c, "Potion")
			fmt.Println("Vous avez reçu une potion de vie")
		case "r":
			return
		default:
			fmt.Println("Wrong choice")
		}
	}
}