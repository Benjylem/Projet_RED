package code

import "fmt"

func Experience(c *Character){
	
    if c.Experience >= 0 && c.Experience < 50 {
		fmt.Println("Vous êtes niveau 0")
	}
	if c.Experience >= 50 && c.Experience < 250 {
		fmt.Println("Vous êtes niveau 1")
	}
	if c.Experience >= 250 && c.Experience < 600 {
		fmt.Println("Vous êtes niveau 2")
	}
	if c.Experience >= 600 && c.Experience < 1100 {
		fmt.Println("Vous êtes niveau 3")		
	}
	if c.Experience >= 1100 && c.Experience < 1750 {
		fmt.Println("Vous êtes niveau 4")		
	}
	if c.Experience >= 1750 && c.Experience < 2550 {
		fmt.Println("Vous êtes niveau 5")		
	}
	if c.Experience >= 2550 && c.Experience < 3500 {
		fmt.Println("Vous êtes niveau 6")		
	}
	if c.Experience >= 3500 && c.Experience < 4600 {
		fmt.Println("Vous êtes niveau 7")		
	}
	if c.Experience >= 4600 && c.Experience < 5850 {
		fmt.Println("Vous êtes niveau 8")		
	}
	if c.Experience >= 5850 && c.Experience < 7250 {
		fmt.Println("Vous êtes niveau 9")		
	}
	if c.Experience >= 7250 {
		fmt.Println("Vous êtes niveau 10, votre job de rêve est à portée de main")		
	}
}
