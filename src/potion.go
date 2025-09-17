package code

import (
	"fmt"
	"strings"
	"time"
)

func GoodPot(c *Character){
	index := -1
	for i, item := range c.Inventory {
		if strings.TrimSpace(strings.ToLower(item)) == "potion" {
		index = i
		break
		}
	}
	if index == -1 {
    	    fmt.Println("❌ Vous n'avez plus de potion.")
        	return
    }
	RemoveInventory(c, index)
	c.CurrentCompDay += 40
	if c.CurrentCompDay > c.MaxCompDay {
        c.CurrentCompDay = c.MaxCompDay
    }
	fmt.Printf("Vous avez bu une potion ! Pv actuelle %d/%d\n", c.CurrentCompDay, c.MaxCompDay)
}

func PoisonPot(c *Character) {
    index := -1
    for i, item := range c.Inventory {
		lower := strings.ToLower(strings.TrimSpace(item))
        if strings.Contains(lower, "potion") && strings.Contains(lower, "poison") {
            index = i
            break
        }
    }
    if index == -1 {
        fmt.Println("❌ Vous n'avez pas de potion de poison.")
        return
    }
    RemoveInventory(c, index)
    fmt.Println("☠️ Vous buvez la potion de poison !")
    for i := 1; i <= 3; i++ {
        time.Sleep(1 * time.Second)
        c.CurrentCompDay -= 10
        if c.CurrentCompDay < 0 {
            c.CurrentCompDay = 0
        }
        fmt.Printf("💢 Dégâts de poison... PV : %d/%d\n", c.CurrentCompDay, c.MaxCompDay)

        if IsDead(c) {
            break
        }
    }
}
