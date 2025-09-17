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
}

var probaReussite = map[string]int{
    "Facile":    75,
    "Moyenne":   45,
    "Difficile": 15,
}

func LaunchMission() {
    rand.Seed(time.Now().UnixNano())

    missions := []Mission{
        {"Remplir son CV", 100, "Facile", "-10 jours d’indemnisation"},
        {"Aller à un entretien simulé", 100, "Moyenne", "-10 jours d’indemnisation"},
        {"Rechercher des offres sur internet", 100, "Facile", "-10 jours d’indemnisation"},
        {"Atelier “Préparer lettre de motivation”", 100, "Moyenne", "-10 jours d’indemnisation"},
        {"Cours de formation courte", 100, "Moyenne", "-10 jours d’indemnisation"},
        {"Simulation d’entretien téléphonique", 100, "Moyenne", "-10 jours d’indemnisation"},
        {"Accompagner un autre migrant", 100, "Moyenne", "-10 jours d’indemnisation"},
        {"Remplir un formulaire administratif", 100, "Facile", "-10 jours d’indemnisation"},
        {"Participer à un atelier numérique", 100, "Moyenne", "-10 jours d’indemnisation"},
        {"Coaching entretien individuel", 100, "Difficile", "-10 jours d’indemnisation"},
    }

    fmt.Println("\033[36m═════════════════════════════════════\033[0m")
    fmt.Println("\t\033[1m📋 MENU DES MISSIONS\033[0m")
    fmt.Println("\033[36m═════════════════════════════════════\033[0m")
    for i, m := range missions {
        fmt.Printf("%d. 🎯 \033[33m%s\033[0m (Difficulté: \033[35m%s\033[0m)\n", i+1, m.Nom, m.Difficulte)
    }

    var choix int
    fmt.Print("\n➡️  Choisissez une mission (1-10) : ")
    fmt.Scanln(&choix)

    if choix < 1 || choix > len(missions) {
        fmt.Println("❌ Choix invalide.")
        return
    }

    mission := missions[choix-1]

    // Styled mission start
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
    } else {
        fmt.Println("❌ \033[31mÉchec de la mission.\033[0m\n💀 Effet négatif :", mission.EchecEffet)
    }

    fmt.Println("\033[36m══════════════════════════════════════════════\033[0m")
}
