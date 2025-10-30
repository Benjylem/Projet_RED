package code

import (
	"fmt"
	"math/rand"
	"time"
)

// Enemy represents a combat opponent
type Enemy struct {
	Name      string
	MaxHP     int
	CurrentHP int
	MinDamage int
	MaxDamage int
	Loot      []string
	ExpReward int
}

// Attack types with different effects
var attackTypes = []string{
	"Coup de poing",
	"Coup de pied",
	"Attaque tournoyante",
	"Charge",
	"Frappe critique",
}

// List of possible enemies with varying difficulty
var enemies = []Enemy{
	{
		Name:      "Conseiller Pôle Emploi Agressif",
		MaxHP:     50,
		MinDamage: 5,
		MaxDamage: 15,
		Loot:      []string{"Plume de Corbeau", "Cuir de Sanglier"},
		ExpReward: 50,
	},
	{
		Name:      "Inspecteur du Travail",
		MaxHP:     80,
		MinDamage: 10,
		MaxDamage: 20,
		Loot:      []string{"Fourrure de Loup", "Peau de Troll"},
		ExpReward: 100,
	},
	{
		Name:      "Manager Toxique",
		MaxHP:     120,
		MinDamage: 15,
		MaxDamage: 30,
		Loot:      []string{"Peau de Troll", "Fourrure de Loup", "Cuir de Sanglier"},
		ExpReward: 150,
	},
	{
		Name:      "Recruteur Impitoyable",
		MaxHP:     100,
		MinDamage: 12,
		MaxDamage: 25,
		Loot:      []string{"Fourrure de Loup", "Plume de Corbeau", "Peau de Troll"},
		ExpReward: 120,
	},
	{
		Name:      "Patron de Start-up",
		MaxHP:     150,
		MinDamage: 20,
		MaxDamage: 35,
		Loot:      []string{"Peau de Troll", "Peau de Troll", "Fourrure de Loup"},
		ExpReward: 200,
	},
}

// LaunchCombat initiates a random combat encounter
func LaunchCombat(c *Character) {
	rand.Seed(time.Now().UnixNano())

	// Select a random enemy
	enemy := enemies[rand.Intn(len(enemies))]
	enemy.CurrentHP = enemy.MaxHP

	fmt.Println("\033[36m══════════════════════════════════════════════\033[0m")
	fmt.Printf("⚔️  \033[1;31mUn %s apparaît !\033[0m\n", enemy.Name)
	fmt.Printf("💀 HP Ennemi : %d/%d\n", enemy.CurrentHP, enemy.MaxHP)
	fmt.Println("\033[36m══════════════════════════════════════════════\033[0m")

	time.Sleep(2 * time.Second)

	// Combat loop
	for enemy.CurrentHP > 0 && c.CurrentCompDay > 0 {
		// Player turn
		fmt.Println("\n\033[33m--- Votre tour ---\033[0m")
		fmt.Printf("💚 Vos CompDay : %d/%d\n", c.CurrentCompDay, c.MaxCompDay)

		playerAttack := attackTypes[rand.Intn(len(attackTypes))]
		playerDamage := rand.Intn(20) + 10 // Random damage between 10-30

		fmt.Printf("🗡️  Vous utilisez : \033[32m%s\033[0m\n", playerAttack)
		time.Sleep(1 * time.Second)
		fmt.Printf("💥 Vous infligez %d dégâts !\n", playerDamage)

		enemy.CurrentHP -= playerDamage
		if enemy.CurrentHP < 0 {
			enemy.CurrentHP = 0
		}
		fmt.Printf("💀 HP Ennemi : %d/%d\n", enemy.CurrentHP, enemy.MaxHP)

		time.Sleep(1 * time.Second)

		// Check if enemy is defeated
		if enemy.CurrentHP <= 0 {
			fmt.Println("\n\033[32m══════════════════════════════════════════════\033[0m")
			fmt.Printf("✅ \033[1;32mVictoire ! %s a été vaincu !\033[0m\n", enemy.Name)
			fmt.Println("\033[32m══════════════════════════════════════════════\033[0m")

			// Distribute random rewards
			distributeRewards(c, &enemy)
			break
		}

		// Enemy turn
		fmt.Println("\n\033[31m--- Tour de l'ennemi ---\033[0m")
		enemyAttack := attackTypes[rand.Intn(len(attackTypes))]
		enemyDamage := rand.Intn(enemy.MaxDamage-enemy.MinDamage+1) + enemy.MinDamage

		fmt.Printf("⚔️  %s utilise : \033[31m%s\033[0m\n", enemy.Name, enemyAttack)
		time.Sleep(1 * time.Second)
		fmt.Printf("💥 Vous subissez %d dégâts !\n", enemyDamage)

		c.CurrentCompDay -= enemyDamage
		if c.CurrentCompDay < 0 {
			c.CurrentCompDay = 0
		}
		fmt.Printf("💚 Vos CompDay : %d/%d\n", c.CurrentCompDay, c.MaxCompDay)

		time.Sleep(1 * time.Second)

		// Check if player is defeated
		if c.CurrentCompDay <= 0 {
			fmt.Println("\n\033[31m══════════════════════════════════════════════\033[0m")
			fmt.Printf("❌ \033[1;31mDéfaite ! Vous avez été vaincu par %s...\033[0m\n", enemy.Name)
			fmt.Println("\033[31m══════════════════════════════════════════════\033[0m")
			IsDead(c)
			break
		}

		time.Sleep(1 * time.Second)
	}
}

// distributeRewards gives random rewards to the player after winning combat
func distributeRewards(c *Character, enemy *Enemy) {
	// Give experience
	c.Experience += enemy.ExpReward
	fmt.Printf("🎉 Vous gagnez %d XP !\n", enemy.ExpReward)

	// Give random loot
	numLoot := rand.Intn(len(enemy.Loot)) + 1 // At least 1, up to all loot items
	fmt.Printf("\n🎁 Récompenses obtenues :\n")

	for i := 0; i < numLoot; i++ {
		if i >= len(enemy.Loot) {
			break
		}
		lootItem := enemy.Loot[rand.Intn(len(enemy.Loot))]

		if CheckMaxItem(c) {
			c.Inventory = append(c.Inventory, lootItem)
			fmt.Printf("   ✅ %s\n", lootItem)
		} else {
			fmt.Printf("   ❌ %s (inventaire plein !)\n", lootItem)
		}
	}

	// Random bonus: chance to get extra CompDay restoration
	if rand.Intn(100) < 30 { // 30% chance
		bonusHP := rand.Intn(20) + 10
		c.CurrentCompDay += bonusHP
		if c.CurrentCompDay > c.MaxCompDay {
			c.CurrentCompDay = c.MaxCompDay
		}
		fmt.Printf("\n💚 Bonus ! Vous récupérez %d jours d'indemnisation !\n", bonusHP)
	}
}
