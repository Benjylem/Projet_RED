package code
import (
	"fmt"
    "math/rand"
    "time"
)

type Mission struct {
    Nom         string
    Exp         int
    Difficulte  string
    EchecEffet  string
    Loot        string
}

var probaReussite = map[string]int{
    "Facile":    75,
    "Moyenne":   45,
    "Difficile": 15,
}

func LaunchMission(c *Character) {
    rand.Seed(time.Now().UnixNano())

    missions := []Mission{
    {"Remplir son CV", 1, "Facile", "-10 jours d’indemnisation", "Plume de Corbeau"},
    {"Aller à un entretien simulé", 20, "Moyenne", "-10 jours d’indemnisation", "Fourrure de Loup"},
    {"Rechercher des offres sur internet", 25, "Facile", "-10 jours d’indemnisation", "Cuir de Sanglier"},
    {"Atelier “Préparer lettre de motivation”", 70, "Moyenne", "-10 jours d’indemnisation", "Fourrure de Loup"},
    {"Cours de formation courte", 70, "Moyenne", "-10 jours d’indemnisation", "Peau de Troll"},
    {"Simulation d’entretien téléphonique", 70, "Moyenne", "-10 jours d’indemnisation", "Plume de Corbeau"},
    {"Accompagner un autre migrant", 80, "Moyenne", "-10 jours d’indemnisation", "Cuir de Sanglier"},
    {"Remplir un formulaire administratif", 2, "Facile", "-10 jours d’indemnisation", "Plume de Corbeau"},
    {"Participer à un atelier numérique", 100, "Moyenne", "-10 jours d’indemnisation", "Fourrure de Loup"},
    {"Coaching entretien individuel", 100, "Difficile", "-10 jours d’indemnisation", "Peau de Troll"},
}


    fmt.Println("\033[36m═════════════════════════════════════\033[0m")
    fmt.Println("\t\033[1m📋 MENU DES MISSIONS\033[0m")
    fmt.Println("\033[36m═════════════════════════════════════\033[0m")
    for i, m := range missions {
        fmt.Printf("%d. 🎯 \033[33m%s\033[0m (Difficulté: \033[35m%s\033[0m)\n", i+1, m.Nom, m.Difficulte)
    }

    var choix int
    fmt.Print("\n➡️  Choisissez une mission (1-10) : ")
    _, err := fmt.Scanln(&choix)

    if err != nil {
        fmt.Println("❌ Erreur de saisie.")
		return
    }

    if choix < 1 || choix > len(missions) {
        fmt.Println("❌ Choix invalide.")
        return
    }

    mission := missions[choix-1]

    fmt.Println("\n\033[36m══════════════════════════════════════════════\033[0m")
    fmt.Printf("🎯 Mission sélectionnée : \033[33m%s\033[0m (Difficulté : \033[35m%s\033[0m)\n", mission.Nom, mission.Difficulte)
    fmt.Println("⏳ En cours... Veuillez patienter...")
    fmt.Println("\033[36m══════════════════════════════════════════════\033[0m")

    time.Sleep(3 * time.Second)
    chance := probaReussite[mission.Difficulte]
    tirage := rand.Intn(100) + 1

	fmt.Printf("\033[36m══════════════════════════════════════════════\033[0m\n")
    fmt.Printf("🎲 Probabilité de réussite : \033[32m%d%%\033[0m", chance)

    if tirage <= chance {
        fmt.Println("✅ \033[32mMission réussie !\033[0m\n🎉 Vous gagnez", mission.Exp, "EXP.")
        c.Experience += mission.Exp
        if mission.Loot != "" {
            if CheckMaxItem(c) {
                c.Inventory = append(c.Inventory, mission.Loot)
                fmt.Printf("🎁 Vous recevez : %s\n", mission.Loot)
            } else {
			fmt.Printf("🎁 Vous devriez recevoir : %s, mais votre inventaire est plein!\n", mission.Loot)
			}
        }
        } else {
        fmt.Println("❌ \033[31mÉchec de la mission.\033[0m\n💀 Effet négatif :", mission.EchecEffet)
        c.CurrentCompDay -= 10
        if c.CurrentCompDay < 0 {
            c.CurrentCompDay = 0
        }
        IsDead(c)
    }
    fmt.Println("\033[36m══════════════════════════════════════════════\033[0m")
}
