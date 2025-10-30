package code

import "fmt"

func IsDead(c *Character) bool {
	redBold := "\033[1;31m"
	reset := "\033[0m"

	if c.CurrentCompDay <= 0 {
		fmt.Printf("\n💀 Malheureusement, vous êtes un %sBOULET%s et avez perdu tout vos jours d'indemnités. Vous avez été expulsé ! 💀\n", redBold, reset)
		c.CurrentCompDay = c.MaxCompDay / 2
		fmt.Printf("⚡ Mais %sMELENCHON%s a été généreux et vous a redonné %d/%d jours d'indemnisation.\n", redBold, reset, c.CurrentCompDay, c.MaxCompDay)
		return true
	}
	return false
}
