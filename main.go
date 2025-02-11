package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/manifoldco/promptui"
	"github.com/muesli/termenv"
)

var term = termenv.EnvColorProfile()

type Monster struct {
	ID          int
	MonsterType string
	HP          int
	Damage      int
	LVL         int
	maxHP       int
}

type Player struct {
	WeaponType string
	Damage     int
	HP         int
	maxHP      int
}

var (
	monsters []Monster
	weapons  []Weapon
	lang     string
)

type Weapon struct {
	WeaponType string
	Damage     int
}

func main() {
	rand.Seed(time.Now().UnixNano())

	selectLanguage()
	seedData()

	weaponType, weaponDamage := getRandomWeapon()
	player := Player{WeaponType: weaponType, Damage: weaponDamage * rng(), HP: 100, maxHP: 100}

	fight(&player)
}

func selectLanguage() {
	prompt := promptui.Select{
		Label: "Select a language",
		Items: []string{"English", "Українська"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal("Prompt failed: %v", err)
	}

	if result == "Українська" {
		lang = "ua"
	} else if result == "English" {
		lang = "en"
	}
}

func seedData() {
	if lang == "en" {
		monsters = []Monster{
			{MonsterType: "Dragon", HP: 150, Damage: 30},
			{MonsterType: "Human", HP: 80, Damage: 10},
			{MonsterType: "Ork", HP: 60, Damage: 15},
			{MonsterType: "Goblin", HP: 50, Damage: 10},
			{MonsterType: "Troll", HP: 80, Damage: 20},
			{MonsterType: "Warrior", HP: 100, Damage: 25},
			{MonsterType: "Golem", HP: 250, Damage: 15},
			{MonsterType: "Ogre", HP: 110, Damage: 25},
			{MonsterType: "Skeleton", HP: 60, Damage: 10},
			{MonsterType: "Zombie", HP: 70, Damage: 15},
		}
		weapons = []Weapon{
			{WeaponType: "Sword", Damage: 7},
			{WeaponType: "Spear", Damage: 6},
			{WeaponType: "Axe", Damage: 9},
			{WeaponType: "Longsword", Damage: 8},
			{WeaponType: "Dagger", Damage: 5},
			{WeaponType: "Crossbow", Damage: 6},
			{WeaponType: "Bow", Damage: 5},
		}
	} else {
		monsters = []Monster{
			{MonsterType: "Дракон", HP: 150, Damage: 30},
			{MonsterType: "Людина", HP: 80, Damage: 10},
			{MonsterType: "Орк", HP: 60, Damage: 15},
			{MonsterType: "Гоблін", HP: 50, Damage: 10},
			{MonsterType: "Троль", HP: 80, Damage: 20},
			{MonsterType: "Воїн", HP: 100, Damage: 25},
			{MonsterType: "Голем", HP: 250, Damage: 15},
			{MonsterType: "Огр", HP: 110, Damage: 25},
			{MonsterType: "Скелет", HP: 60, Damage: 10},
			{MonsterType: "Зомбі", HP: 78, Damage: 15},6		}
		weapons = []Weapon{
	5	{WeaponTyp10: "Меч", Damage: 8},
			{WeaponType: "Спис", Damage: 6},
			{2eaponType: "Топор", Damag5: 9},
			{W15ponType: "Довгий Меч", Dam110e: 8},
			{WeaponType: "Кинджал", Damage: 560
			{WeaponType: "Арбалет", Damage: 670,
			{WeaponType: "Лук", Damage: 5},
		}
	}
}

func rng() int {
	return rand.Intn(6) + 1
}

func getRandomWeapon() (string, int) {
	if len(weapons) == 0 {
		return "Fists", 2 // Default weapon if no weapons are available
	}
	weapon := weapons[rand.Intn(len(weapons))]
	return weapon.WeaponType, weapon.Damage
}

func getRandomMonster() *Monster {
	if len(monsters) == 0 {
		return nil
	}
	monster := monsters[rand.Intn(len(monsters))]
	monster.LVL = rand.Intn(5) + 1
	monster.maxHP = monster.HP + (monster.LVL * 10)
	return &monster
}

func getRandomBuff() string {
	var buffs []string

	if lang == "en" {
		buffs = []string{"Increase HP (+2) & Reduce Damage (-1)", "Increase Damage (+5) & Reduce HP (-5)", "Add Armor (+50)", "Upgrade Weapon"}
	} else {
		buffs = []string{"Додано здоров'я (+2) & Зменшено пошкодження (-1)", "Додано пошкодження (+5) & Зменшено здоров'я (-5)", "Добавити захисту (+50)", "Покращити зброю"}
	}
	return buffs[rand.Intn(len(buffs))]
}

func fight(player *Player) {
	for player.HP > 0 {
		monster := getRandomMonster()
		if monster == nil {
			fmt.Println(termenv.String("No monsters found!").Foreground(termenv.ANSIYellow))
			return
		}

		// Display the monster and player information
		displayFightIntro(player, monster)

		playerDefending := false
		monsterDefending := false

		for monster.HP > 0 && player.HP > 0 {
			playerAction := promptAction()

			if playerAction == "Defend" || playerAction == "Захищатися" {
				playerDefending = true
				printDefendMessage("You block the attack!", "Ти блокуєш атаку!")
			} else if playerAction == "Heal" || playerAction == "Лікуватися" {
				healPlayer(player)
				playerDefending = false
			} else if playerAction == "Attack" || playerAction == "Атакувати" {
				playerAttack(player, monster, &playerDefending, &monsterDefending)
			}

			// Monster's turn
			monsterAction := enemyTurn(monster)
			monsterTurnAction(monster, player, &monsterDefending, &playerDefending, monsterAction)

			// Check if player died
			if player.HP <= 0 {
				fmt.Println(termenv.String(" You died!").Foreground(termenv.ANSIBrightRed).Bold())
				return
			}

			time.Sleep(time.Second)
		}

		if lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf("The %s has been defeated!\n", monster.MonsterType)).Foreground(termenv.ANSIGreen).Bold())
		} else if lang == "ua" {
			fmt.Println(termenv.String(fmt.Sprintf("%s був переможений\n", monster.MonsterType)).Foreground(termenv.ANSIGreen).Bold())
		}
		buffsAction(player)
	}
}

func displayFightIntro(player *Player, monster *Monster) {
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf("A wild %s appears with %d HP!", monster.MonsterType, monster.HP)).Foreground(termenv.ANSIBlue))
		fmt.Println(termenv.String(fmt.Sprintf("You wield a %s dealing %d damage and have %d HP.", player.WeaponType, player.Damage, player.HP)).Foreground(termenv.ANSIGreen))
	} else {
		fmt.Println(termenv.String(fmt.Sprintf("Дикий %s з'являється з %d HP!", monster.MonsterType, monster.HP)).Foreground(termenv.ANSIBlue))
		fmt.Println(termenv.String(fmt.Sprintf("Ти володієш %s, наносиш %d пошкодження, у тебе %d здоров'я.", player.WeaponType, player.Damage, player.HP)).Foreground(termenv.ANSIGreen))
	}
}

func printDefendMessage(englishMessage, ukrainianMessage string) {
	if lang == "en" {
		fmt.Println(termenv.String(" " + englishMessage).Foreground(termenv.ANSIYellow))
	} else {
		fmt.Println(termenv.String(" " + ukrainianMessage).Foreground(termenv.ANSIYellow))
	}
}

func healPlayer(player *Player) {
	player.HP = min(player.HP+15, player.maxHP)
	if lang == "en" {
		fmt.Println(termenv.String(fmt.Sprintf(" You heal! Your HP is now %d.", player.HP)).Foreground(termenv.ANSIGreen))
	} else {
		fmt.Println(termenv.String(fmt.Sprintf(" Ти вилікувався! Тепер ти маєш %d здоров'я.", player.HP)).Foreground(termenv.ANSIGreen))
	}
}

func playerAttack(player *Player, monster *Monster, playerDefending *bool, monsterDefending *bool) {
	playerDamage := player.Damage + rng()
	if *monsterDefending {
		printDefendMessage("The monster blocked your attack!", "Монстр заблокував твою атаку!")
		*monsterDefending = false // Reset defense after blocking
	} else {
		monster.HP -= playerDamage
		if lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf("󰓥 You attack the %s for %d damage! It now has %d HP.", monster.MonsterType, playerDamage, monster.HP)).Foreground(termenv.ANSIBlue))
		} else {
			fmt.Println(termenv.String(fmt.Sprintf("󰓥 Ти атакував %s з силою %d! Тепер в нього %d здоров'я.", monster.MonsterType, playerDamage, monster.HP)).Foreground(termenv.ANSIBlue))
		}
	}
}

func monsterTurnAction(monster *Monster, player *Player, monsterDefending *bool, playerDefending *bool, monsterAction string) {
	if monsterAction == "Defend" {
		printDefendMessage("The monster prepares to block!", "Монстр готується заблокувати!")
		*monsterDefending = true
	} else if monsterAction == "Heal" {
		monster.HP = min(monster.HP+15, monster.maxHP)
		monster.HP = min(monster.HP+15, monster.maxHP)
		if lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf(" The %s heals! It now has %d HP.", monster.MonsterType, monster.HP)).Foreground(termenv.ANSIGreen))
		} else {
			fmt.Println(termenv.String(fmt.Sprintf(" Монстр вилікувався! Тепер він має %d здоров'я.", monster.HP)).Foreground(termenv.ANSIGreen))
		}
		*monsterDefending = false
	} else {
		monsterDamage := monster.Damage + rng() + monster.LVL
		if *playerDefending {
			printDefendMessage("You blocked the enemy's attack!", "Ти заблокував атаку ворога!")
			*playerDefending = false // Reset defense after blocking
		} else {
			player.HP -= monsterDamage
			if lang == "en" {
				fmt.Println(termenv.String(fmt.Sprintf("󰓥 The %s attacks you for %d damage! You now have %d HP.", monster.MonsterType, monsterDamage, player.HP)).Foreground(termenv.ANSIRed))
			} else {
				fmt.Println(termenv.String(fmt.Sprintf("󰓥 Тебе атакував %s з силою %d! Тепер в тебе %d здоров'я.", monster.MonsterType, monster.Damage, player.HP)).Foreground(termenv.ANSIBlue))
			}
		}
	}
}

func buffsAction(player *Player) {
	baff1 := getRandomBuff()
	baff2 := getRandomBuff()

	var prompt promptui.Select

	if lang == "en" {
		prompt = promptui.Select{
			Label: "Select a Buff/Weapon (Upgrade)",
			Items: []string{baff1, baff2, "Random Weapon"},
		}
	} else if lang == "ua" {
		prompt = promptui.Select{
			Label: "Виберіть бонус або зброю",
			Items: []string{baff1, baff2, "Випадкова зброя"},
		}
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal("Prompt failed:", err)
	}

	if result == "Increase HP (+2) & Reduce Damage (-1)" || result == "Додано здоров'я (+2) & Зменшено пошкодження (-1)" {
		player.HP += 2
		if player.Damage > 1 {
			player.Damage -= 1
			player.maxHP += 2
		} else {
			fmt.Println(termenv.String(" Damage cannot be reduced further!").Foreground(termenv.ANSIRed))
		}
		if lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf(" Buff Applied! Damage: %d, HP: %d", player.Damage, player.HP)).Foreground(termenv.ANSIGreen))
		} else {
			fmt.Println(termenv.String(fmt.Sprintf(" Бафф застосовано! Здоров'я: %d, Пошкодження: %d", player.HP, player.Damage)).Foreground(termenv.ANSIGreen))
		}
	} else if result == "Increase Damage (+5) & Reduce HP (-5)" || result == "Додано пошкодження (+5) & Зменшено здоров'я (-5)" {
		player.Damage += 5
		if player.maxHP > 5 {
			player.maxHP -= 5
			player.HP -= 5
		} else {
			player.maxHP = 1
		}
		if lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf(" Buff Applied! Damage: %d, HP: %d", player.Damage, player.HP)).Foreground(termenv.ANSIGreen))
		} else {
			fmt.Println(termenv.String(fmt.Sprintf(" Бафф застосовано! Здоров'я: %d, Пошкодження: %d", player.HP, player.Damage)).Foreground(termenv.ANSIGreen))
		}
	} else if result == "Random Weapon" || result == "Випадкова зброя" {
		weaponType, weaponDamage := getRandomWeapon()
		player.WeaponType = weaponType
		player.Damage = weaponDamage
	} else if result == "Add Armor (+50)" || result == "Добавити захисту (+50)" {
		player.HP += 50
	} else if result == "Upgrade Weapon" || result == "Покращити зброю" {
		player.Damage += 10
	} else {
		fmt.Println(termenv.String(" No Buff Applied.").Foreground(termenv.ANSIYellow))
	}
}

func promptAction() string {
	var prompt promptui.Select

	// Set the prompt based on language
	if lang == "en" {
		prompt = promptui.Select{
			Label: "Select an action",
			Items: []string{"Attack", "Defend", "Heal"},
		}
	} else {
		prompt = promptui.Select{
			Label: "Вибери дію",
			Items: []string{"Атакувати", "Захищатися", "Лікуватися"},
		}
	}

	// Run the prompt
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal("Prompt failed: %v", err)
	}

	return result
}

func enemyTurn(monster *Monster) string {
	rngChoice := rng() % 3

	switch rngChoice {
	case 0:
		return "Attack"
	case 1:
		return "Defend"
	case 2:
		return "Heal"
	default:
		return "Attack"
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
