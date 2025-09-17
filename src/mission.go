package main

import (
    "fmt"
    "math/rand"
    "time"
)

// Mission représente une mission possible
type Mission struct {
    Nom         string
    Exp         int
    Difficulte  string
    EchecEffet  string
}

// Probabilités de réussite selon la difficulté
var probaReussite = map[string]int{
    "Facile":    75,
    "Moyenne":   45,
    "Difficile": 15,
}

func main() {
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

    fmt.Println("=== Menu des Missions ===")
    for i, m := range missions {
        fmt.Printf("%d. %s (Difficulté: %s)\n", i+1, m.Nom, m.Difficulte)
    }

    var choix int
    fmt.Print("\nChoisissez une mission (1-10) : ")
    fmt.Scanln(&choix)

    if choix < 1 || choix > len(missions) {
        fmt.Println("Choix invalide.")
        return
    }

    mission := missions[choix-1]
    fmt.Printf("\nVous avez choisi : %s\n", mission.Nom)

    chance := probaReussite[mission.Difficulte]
    tirage := rand.Intn(100) + 1

    fmt.Printf("Calcul de la réussite... (Probabilité : %d%%)\n\n", chance)
    time.Sleep(3 * time.Second)

    if tirage <= chance {
        fmt.Println(":white_check_mark: Mission réussie ! Vous gagnez", mission.Exp, "EXP.")
    } else {
        fmt.Println(":x: Échec de la mission. Effet négatif :", mission.EchecEffet)
    }
}