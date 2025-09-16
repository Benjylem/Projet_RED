package code

import "fmt"

func IsDead(c *Character) bool {
    if c.CurrentCompDay <= 0 {
        fmt.Println("\n💀 Vous êtes mort ! 💀")
        c.CurrentCompDay = c.MaxCompDay / 2
        fmt.Printf("⚡ Vous êtes ressuscité avec %d/%d PV.\n", c.CurrentCompDay, c.MaxCompDay)
        return true
    }
    return false
}
