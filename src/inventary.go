package code

import (
	"fmt"
	"bufio"
	"strings"
)

func AddINVENTAIRE(c *Charactere, item string) {
	item = strings.TrimSpace(item)
	c.INVENTAIRE = append(c.INVENTAIRE, item)
}

func RemoveINVENTAIRE(c *Charactere, index int) {
	if index >= 0 && index < len(c.INVENTAIRE) {
		c.INVENTAIRE = append(c.INVENTAIRE[:index], c.INVENTAIRE[index+1:]...)
	}
}

func ShowInventaire(c *Charactere) {
	if len(c.INVENTAIRE) == 0 {
		fmt.Println("Inventaire vide.")
	} else {
		for i, item := range c.INVENTAIRE {
			fmt.Printf("[%d] %q\n", i+1, item)
		}
	}
}

func AccessINVENTAIRE(c *Charactere, reader *bufio.Reader) {
	for {
		fmt.Println("\n=== INVENTAIRE ===")
		ShowInventaire(c)
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

func AccesMerchant(c *Charactere, reader *bufio.Reader) {
	for {
		fmt.Println("\n=== Bienvenu à la CAF ===")
		fmt.Println("[1] Potion de vie (gratuit)")
		fmt.Println("[R] Retour")
		fmt.Println("Choice : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "1":
			AddINVENTAIRE(c, "Potion")
			fmt.Println("Vous avez reçu une potion de vie")
		case "r":
			return
		default:
			fmt.Println("Wrong choice")
		}
	}
}