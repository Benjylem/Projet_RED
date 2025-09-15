package code

import (
	"fmt"
	"strings"
)

func TakePot(c *Character){
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
