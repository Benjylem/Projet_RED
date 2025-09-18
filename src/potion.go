package code

import (
	"fmt"
	"strings"
	"time"
)

func GoodPot(c *Character){
	index := -1
	for i, item := range c.Inventory {
		if strings.TrimSpace(strings.ToLower(item)) == "oasis" {
		index = i
		break
		}
	}
	if index == -1 {
    	    fmt.Println("❌ Vous n'avez plus d'oasis.")
        	return
    }
	RemoveInventory(c, index)
	c.CurrentCompDay += 40
	if c.CurrentCompDay > c.MaxCompDay {
        c.CurrentCompDay = c.MaxCompDay
    }
	fmt.Printf("Vous avez bu un oasis ! CompDay actuel %d/%d\n", c.CurrentCompDay, c.MaxCompDay)
}

func PoisonPot(c *Character) {
    index := -1
    for i, item := range c.Inventory {
		lower := strings.ToLower(strings.TrimSpace(item))
        if strings.Contains(lower, "oasis") && strings.Contains(lower, "Café réchauffé") {
            index = i
            break
        }
    }
    if index == -1 {
        fmt.Println("❌ Vous n'avez pas de Café réchauffé.")
        return
    }
    RemoveInventory(c, index)
    fmt.Println("☠️ Vous buvez le Café réchauffé !")
    for i := 1; i <= 3; i++ {
        time.Sleep(1 * time.Second)
        c.CurrentCompDay -= 10
        if c.CurrentCompDay < 0 {
            c.CurrentCompDay = 0
        }
        fmt.Printf("💢 Dégâts de Café réchauffé... CompDay : %d/%d\n", c.CurrentCompDay, c.MaxCompDay)

        if IsDead(c) {
            break
        }
    }
}
func Skill(c *Character){
	index := -1
	for i, item := range c.Inventory {
		if strings.TrimSpace(strings.ToLower(item)) == "Diplômatie" {
		index = i
		break
		}
	}
	if index == -1 {
    	    fmt.Println("❌ Vous n'avez plus de formation diplomatie disponible.")
        	return
    }
	RemoveInventory(c, index)
	c.CurrentCompDay += 20
	if c.CurrentCompDay > c.MaxCompDay {
        c.CurrentCompDay = c.MaxCompDay
    }
	fmt.Printf("Vous êtes plus calme ! ! CompDay actuel %d/%d\n", c.CurrentCompDay, c.MaxCompDay)
}
