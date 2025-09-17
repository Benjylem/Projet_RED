package code

import (
    "fmt"
    "math/rand"
    "time"
)

type Item struct {
    Nom    string
    Effet  string
    Rarete string
}

// Fonction qui vérifie la limite d’inventaire (max 10 objets)
func addToInventory(inventory []Item, newItem Item) []Item {
    if len(inventory) >= 10 {
        fmt.Println("⚠️ Inventaire plein ! Impossible d’ajouter :", newItem.Nom)
        return inventory
    }
    return append(inventory, newItem)
}

func accessInventory(inventory []Item) {
    fmt.Println("=== Inventaire du joueur ===")
    if len(inventory) == 0 {
        fmt.Println("Votre inventaire est vide.")
        return
    }

    for i, item := range inventory {
        fmt.Printf("%d. %s (%s) → Effet : %s\n", i+1, item.Nom, item.Rarete, item.Effet)
    }
}

var shop = []Item{
    {"Tickets resto", "+20 EXP pendant 20 missions", "Commun"},
    {"CV premium", "+20% chance de réussite", "Rare"},
    {"Pause clope", "+10 jours d’indemnisation", "Peu commun"},
    {"Formation Pôle Emploi", "+30 EXP pendant 20 missions", "Peu commun"},
    {"Prime d’activité", "+10% EXP pendant 5 missions", "Rare"},
    {"Recommandation d’un pote", "+15% chance réussite pour 3 missions", "Peu commun"},
    {"Café du conseiller", "+20 jours d’indemnisation", "Rare"},
    {"Objet mystère", "Bonus ou malus aléatoire", "Commun"},
}

// Acheter un objet
func buyItem(inventory []Item, choice int) []Item {
    if choice < 0 || choice >= len(shop) {
        fmt.Println("Choix invalide.")
        return inventory
    }
    selected := shop[choice]
    fmt.Printf("Vous avez acheté : %s (%s)\n", selected.Nom, selected.Rarete)
    return addToInventory(inventory, selected)
}

// Récompense de mission
func missionReward(inventory []Item) []Item {
    rand.Seed(time.Now().UnixNano())
    chance := rand.Intn(100)
    if chance < 30 {
        reward := shop[rand.Intn(len(shop))]
        fmt.Printf("🎉 Vous avez gagné un objet en mission : %s (%s)\n", reward.Nom, reward.Rarete)
        inventory = addToInventory(inventory, reward)
    } else {
        fmt.Println("Aucun objet gagné cette fois.")
    }
    return inventory
}

func main() {
    var inventory []Item

    // Exemple : remplir l’inventaire avec des achats
    for i := 0; i < 12; i++ {
        inventory = buyItem(inventory, i%len(shop))
    }

    // Affichage final
    accessInventory(inventory)
}

// Structure Item
type Item struct {
	Name        string
	Rarity      string
	Price       int
	Description string
}

// Structure Joueur
type Player struct {
	Exp       int
	Inventory []Item
	MaxItems  int
}

// Fonction pour afficher le marchand
func showMerchant(player *Player, shop []Item) {
	fmt.Println("=== Bienvenue au Marchand Pole Emploi ===")
	fmt.Println("Voici les objets disponibles :")
	for i, item := range shop {
		fmt.Printf("%d. %s [%s] - Prix: %d EXP\n   Effet: %s\n", i+1, item.Name, item.Rarity, item.Price, item.Description)
	}

	var choice int
	fmt.Println("\nEntrez le numéro de l'objet que vous souhaitez acheter (0 pour quitter) :")
	fmt.Scan(&choice)

	if choice == 0 {
		fmt.Println("Vous quittez le marchand.")
		return
	}

	if choice < 1 || choice > len(shop) {
		fmt.Println("Choix invalide.")
		return
	}

	selectedItem := shop[choice-1]

	// Vérification inventaire plein
	if len(player.Inventory) >= player.MaxItems {
		fmt.Println("Inventaire plein, impossible d’acheter cet objet.")
		return
	}

	// Vérification EXP
	if player.Exp < selectedItem.Price {
		fmt.Printf("Pas assez d’expérience pour acheter %s.\n", selectedItem.Name)
		return
	}

	// Achat réussi
	player.Exp -= selectedItem.Price
	player.Inventory = append(player.Inventory, selectedItem)
	fmt.Printf("Le joueur a récupéré %s (%s) et a perdu %d EXP de son compteur.\n",
		selectedItem.Name, selectedItem.Description, selectedItem.Price)
}

// Programme principal pour tester
func main() {
	// Joueur de test
	player := Player{
		Exp:       500,
		Inventory: []Item{},
		MaxItems:  10,
	}

	// Objets disponibles dans le shop
	shop := []Item{
		{"Ticket resto", "Commun", 150, "+20 EXP pendant 20 jours"},
		{"Formation Pole Emploi", "Rare", 300, "+30 EXP pendant 20 jours"},
		{"PC portable d’occasion", "Épique", 600, "Augmente les chances de réussite des missions"},
		{"Mystery Box", "Très variable", 100, "Contient un bonus ou un malus aléatoire"},
	}

	// Lancer l’interface du marchand
	showMerchant(&player, shop)

	// Afficher inventaire après achat
	fmt.Println("\nInventaire actuel du joueur :")
	for _, item := range player.Inventory {
		fmt.Printf("- %s [%s]\n", item.Name, item.Rarity)
	}
	fmt.Printf("EXP restant : %d\n", player.Exp)
}
