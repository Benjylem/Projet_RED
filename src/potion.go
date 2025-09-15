package code

import (
	"fmt"
	"strings"
)

func TakePot(c *Charactere){
	index := -1
	for i, item := range c.INVENTAIRE {
		if strings.TrimSpace(strings.ToLower(item)) == "potion" {
		index = i
		break
		}
	}
	if index == -1 {
    	    fmt.Println("❌ Vous n'avez plus de potion.")
        	return
    }
	RemoveINVENTAIRE(c, index)
	c.CurrentHP += 40
	if c.CurrentHP > c.MaxHP {
        c.CurrentHP = c.MaxHP
    }
	fmt.Printf("Vous avez bu une potion ! Pv actuelle %d/%d\n", c.CurrentHP, c.MaxHP)
}
