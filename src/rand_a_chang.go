package code

import (
    "fmt"
    "math/rand"
    "time"
)

type Mission struct {
    Name       string
    Difficulty string
    HPDamage   int
    ExpReward  int
}

func LaunchMission(c *Character) {

    missions := []Mission{
        {"Créer un CV", "Facile", 5, 50},
        {"Trouver un recruteur", "Moyenne", 10, 80},
        {"Aller à un entretien d’embauche", "Moyenne", 15, 100},
        {"Participer à un atelier de recherche d’emploi", "Difficile", 20, 120},
        {"Simulation d’entretien téléphonique", "Difficile", 25, 150},
    }

    mission := missions[rand.Intn(len(missions))]
	// genre len(mission) pour le 1er c'est 5 hp (en moins (si on perd (c'est toujour random)))
	//rand.Intn c'est le nb random qui est choisie(c'est ce qui fait le random)
	//mission[index] bah juste la premiere la deuxiem etc

    fmt.Println("\033[36m══════════════════════════════════════\033[0m")
    fmt.Printf("🎯 Mission : %s (Difficulté : %s)\n", mission.Name, mission.Difficulty)
    fmt.Println("En cours... ⏳")
    time.Sleep(2 * time.Second)

	//check voir les dommage et les apply
    c.CurrentCompDay -= mission.HPDamage
    if c.CurrentCompDay < 0 {
        c.CurrentCompDay = 0
    }

    fmt.Printf("⚔️  Vous avez perdu %d PV. (PV restant : %d/%d)\n", mission.HPDamage, c.CurrentCompDay, c.MaxCompDay)

    c.Experience += mission.ExpReward
    fmt.Printf("🏆 Vous gagnez %d points d'expérience ! Total : %d\n", mission.ExpReward, c.Experience)

    if IsDead(c) {
        fmt.Println("⚠️ Vous avez perdu connaissance après la mission...")
    }

    fmt.Println("\033[36m══════════════════════════════════════\033[0m")
}
