package code

import "fmt"

func CompDay(c *Character){
	
    if c.CurrentCompDay >= c.MaxCompDay/2 {
		fmt.Printf("\033[36m%d / %d\033[0m", c.CurrentCompDay, c.MaxCompDay)
	}
	if c.CurrentCompDay < c.MaxCompDay/2 && c.CurrentCompDay > 0 {
		fmt.Printf("\033[31m%d / %d\033[0m", c.CurrentCompDay, c.MaxCompDay)
	}
}
