package code

import (
    "fmt"
	"strings"
	"unicode"

)

// Définition de la structure Equipment
type Equipment struct {
	EquipmentHead string
	ChestEquipment string
	FootEquipment string
}
// Définition de la structure Character
type Character struct {
    Name      string
    Class     string
    Level     int
    Skill     []string
    MaxCompDay     int
    CurrentCompDay int
	Experience int
    Inventory []string
	Equipment Equipment

}

// Fonction d'initialisation qui adapte les caractéristiques selon la classe
func InitCharacter(name, class string) Character {
    switch class {
    case "sud":
        return Character{
            Name:      name,
            Class:     class,
            Level:     0,
            Skill:     []string{"colère"},
            MaxCompDay:     350,
            CurrentCompDay: 10,
			Experience: 0,
            Inventory: []string{},
			Equipment: Equipment{
				EquipmentHead: "beret",
				ChestEquipment: "chemise",
				FootEquipment: "mocassins",
			},
        }
    case "est":
        return Character{
            Name:      name,
            Class:     class,
            Level:     0,
            Skill:     []string{"colère"},
            MaxCompDay:     300,
            CurrentCompDay: 150,
			Experience: 0,
            Inventory: []string{},
			Equipment: Equipment{
				EquipmentHead: "casquette",
				ChestEquipment: "sweat",
				FootEquipment: "basket",
			},
        }
    case "nord":
        return Character{
            Name:      name,
            Class:     class,
            Level:     0,
            Skill:     []string{"colère"},
            MaxCompDay:     400,
            CurrentCompDay: 200,
			Experience: 0,
            Inventory: []string{},
			Equipment: Equipment{
				EquipmentHead: "bob",
				ChestEquipment: "marcel",
				FootEquipment: "claquettes-chaussettes",
			},
        }
    default:
        // Classe invalide -> personnage par défaut
        return Character{
            Name:      name,
            Class:     "inconnu",
            Level:     0,
            MaxCompDay:     100,
            CurrentCompDay: 100,
			Experience: 0,
            Inventory: []string{},
        }
    }
}

// Vérifie que le nom contient uniquement des lettres
func IsAlpha(s string) bool {
    for _, r := range s {
        if !unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

// Fonction de création de personnage par l'utilisateur
func CharacterCreation() Character {
    var name, class string

    // Demander le nom
    for {
        fmt.Print("Entrez le nom de votre personnage (lettres uniquement) : ")
        fmt.Scanln(&name)
        if IsAlpha(name) {
            break
        }
       fmt.Println("\033[31m❌ Le nom ne doit contenir que des lettres.\033[0m") // Rouge
    }

    name = strings.Title(strings.ToLower(name))

    for {
        fmt.Print("Choisissez une classe (sud / est / nord) : ")
        fmt.Scanln(&class)
        class = strings.ToLower(class)
        if class == "sud" || class == "est" || class == "nord" {
            break
        }
        fmt.Println("\033[31m❌ Classe invalide, réessayez.\033[0m")
    }

    return InitCharacter(name, class)
}

// Nouvelle fonction displayInfo
func DisplayInfo(c Character) {
    fmt.Println("\033[36m══════════════════════════════════════\033[0m")
    fmt.Println("\033[1;33m       🎉 PERSONNAGE CRÉÉ 🎉        \033[0m")
    fmt.Println("\033[1;33m       PARTEZ A L'AVENTURE           \033[0m")
    fmt.Println("\033[36m══════════════════════════════════════\033[0m")

    fmt.Printf("\033[32mNom :\033[0m %s\n", c.Name)
    fmt.Printf("\033[32mClasse :\033[0m %s\n", c.Class)
    fmt.Printf("\033[32mNiveau :\033[0m %d\n", c.Level)
    fmt.Printf("\033[32mSkill :\033[0m %s\n", c.Skill)
    fmt.Printf("\033[32mCompensation Day :\033[0m %d / %d\n", c.CurrentCompDay, c.MaxCompDay)
    fmt.Printf("\033[32mExpérience :\033[0m %d\n", c.Experience)

    fmt.Println("\n\033[36m--- Équipement ---\033[0m")
    fmt.Printf("  🪖 Tête : %s\n", c.Equipment.EquipmentHead)
    fmt.Printf("  👕 Corps : %s\n", c.Equipment.ChestEquipment)
    fmt.Printf("  👟 Pieds : %s\n", c.Equipment.FootEquipment)

    fmt.Println("\n\033[36m--- Inventaire ---\033[0m")
    for i, item := range c.Inventory {
        fmt.Printf("  %d. %s\n", i+1, item)
    }

    fmt.Println("\033[36m══════════════════════════════════════\033[0m")
}
