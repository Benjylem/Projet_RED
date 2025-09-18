package code

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Menu(c *Character) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\033[H\033[2J")

		fmt.Println("\033[36m══════════════════════════════════════\033[0m")
		fmt.Println("\033[1;33m        🏢 Bienvenue à Pôle Emploi Quest 🏢        \033[0m")
		fmt.Println("\033[36m══════════════════════════════════════\033[0m")

fmt.Println("\033[32m[0]\033[0m 📜 Niveau")
fmt.Println("\033[32m[1]\033[0m 📜 Afficher les informations du personnage")
fmt.Println("\033[32m[2]\033[0m 🎒 Accéder à l'inventaire")
fmt.Println("\033[32m[3]\033[0m ⚔️  Chercher une mission")
fmt.Println("\033[32m[4]\033[0m 📅 Consulter les jours d'indemnisation")
fmt.Println("\033[32m[5]\033[0m 🏪 Epicerie")
fmt.Println("\033[32m[6]\033[0m 🏪 Aller à la CAF (Marchand)")
fmt.Println("\033[35m[7] 🎉 Qui sont les artistes?\033[0m")
fmt.Println("\033[31m[8] ❌ Quitter le jeu\033[0m")
fmt.Println("\033[36m══════════════════════════════════════\033[0m")


		fmt.Print("\033[1;34mEntrez votre choix : \033[0m")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "0":
			Experience(c)
		case "1":
			DisplayInfo(*c)
		case "2":
			AccessInventory(c, reader)
		case "3":
			LaunchMission(c)
		case "4":
			CompDay(c)
		case "5":
			ForgeMenu(c)
		case "6":
			AccesMerchant(c, reader)
		case "7":
			fmt.Println("\nLes artistes sont : \nABBA \n& \nSteven Spielberg") //Mission annexe 6
		case "8":
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
