package code

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)

func Menu(c *Charactere) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\033[H\033[2J")

		fmt.Println("======================================")
		fmt.Println("     Bienvenue à Pôle Emploi Quest   ")
		fmt.Println("======================================")
		fmt.Println("[1] Afficher les information du personnage")
		fmt.Println("[2] Accéder à l'inventaire")
		fmt.Println("[3] Chercher une mission")
		fmt.Println("[4] Consulter les jour d'indemnisation")
		fmt.Println("[5] Aller à la CAF (Merchand)")
		fmt.Println("[6] Quitter le jeu")
		fmt.Println("======================================")
		fmt.Print("Entrez votre choix : ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			DisplayInfo(c)
		case "2":
			AccessINVENTAIRE(c, reader)
		case "3":
			//launchMission()
		case "4":
			//checkIndemnisation()
		case "5":
			AccesMerchant(c, reader)
		case "6":
			fmt.Println("\nMerci d'avoir joué ! À bientôt 👋")
			return
		default:
			fmt.Println("\nChoix invalide, réessayez.")
		}
		fmt.Println("\nAppuyez sur Entrée pour revenir au menu.")
		reader.ReadString('\n')
	}
}
// func displayInfo() {
// 	fmt.Println("\n[INFO] Fonction d'affichage du personnage à implémenter.")
// }

// func launchMission() {
// 	fmt.Println("\n[MISSION] Fonction de recherche de mission à implémenter.")
// }

// func checkIndemnisation() {
// 	fmt.Println("\n[INDEMNISATION] Fonction de suivi des jours à implémenter.")
// }
