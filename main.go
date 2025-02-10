package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/charmbracelet/log"
	"github.com/manifoldco/promptui"
	_ "github.com/mattn/go-sqlite3"
	"github.com/muesli/termenv"
)

var db *sql.DB
var term = termenv.EnvColorProfile()

type Monster struct {
	ID          int
	MonsterType string
	HP          int
	Damage      int
}

type Player struct {
	WeaponType string
	Damage     int
	HP         int
}

func main() {
	var err error
	db, err = sql.Open("sqlite3", "./.yasg.db")
	if err != nil {
		log.Fatal("Error opening database:", err)
	}
	defer db.Close()

	createTables()
	seedData()
	rand.Seed(time.Now().UnixNano())

	weaponType, weaponDamage := getRandomWeapon()
	player := Player{WeaponType: weaponType, Damage: weaponDamage * rng(), HP: 100}

	selectLanguage()
	fight(&player)
}

func createTables() {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS monsters (
			id INTEGER PRIMARY KEY,
			monsterType TEXT,
			hp INTEGER,
			damage INTEGER
		);
		CREATE TABLE IF NOT EXISTS weapons (
			id INTEGER PRIMARY KEY,
			weaponType TEXT,
			damage INTEGER
		);
	`)
	if err != nil {
		log.Fatal("Error creating tables:", err)
	}
}

var lang string

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
		addMonster("Dragon", 150, 30)
		addMonster("Human", 50, 10)
		addMonster("Ork", 40, 15)
		addMonster("Goblin", 20, 5)
		addMonster("Troll", 60, 20)
		addMonster("Warrior", 100, 15)
		addMonster("Golem", 200, 20)
		addMonster("Ogre", 80, 25)
		addMonster("Skeleton", 30, 10)

		addWeapon("Sword", 7)
		addWeapon("Spear", 6)
		addWeapon("Axe", 9)
		addWeapon("Longsword", 8)
		addWeapon("Dagger", 5)
		addWeapon("Crossbow", 6)
		addWeapon("Bow", 5)
	} else {
		addMonster("Дракон", 150, 30)
		addMonster("Людина", 50, 10)
		addMonster("Орк", 40, 15)
		addMonster("Гоблін", 20, 5)
		addMonster("Троль", 60, 20)
		addMonster("Воїн", 100, 15)
		addMonster("Голем", 200, 20)
		addMonster("Огр", 80, 25)
		addMonster("Скелет", 30, 10)

		addWeapon("Меч", 7)
		addWeapon("Спис", 6)
		addWeapon("Топор", 9)
		addWeapon("Довгий Меч", 8)
		addWeapon("Кинджал", 5)
		addWeapon("Арбалет", 6)
		addWeapon("Лук", 5)
	}
}

func rng() int {
	return rand.Intn(6) + 1
}

func addWeapon(weaponType string, damage int) {
	stmt, err := db.Prepare("INSERT INTO weapons (weaponType, damage) VALUES (?, ?)")
	if err != nil {
		log.Fatal("Error preparing weapon insert:", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(weaponType, damage)
	if err != nil {
		log.Fatal("Error inserting weapon:", err)
	}
}

func addMonster(monsterType string, hp int, damage int) {
	stmt, err := db.Prepare("INSERT INTO monsters (monsterType, hp, damage) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal("Error preparing monster insert:", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(monsterType, hp, damage)
	if err != nil {
		log.Fatal("Error inserting monster:", err)
	}
}

func deleteMonster(id int) {
	_, err := db.Exec("DELETE FROM monsters WHERE id = ?", id)
	if err != nil {
		log.Fatal("Error deleting defeated monster:", err)
	}
}

func getRandomBuff() string {
	var buffs []string

	if lang == "en" {
		buffs = []string{"Increase HP (+2) & Reduce Damage (-1)", "Increase Damage (+5) & Reduce HP (-5)"}
	} else {
		buffs = []string{"Додано здоров'я (+2) & Зменшено пошкодження (-1)", "Додано пошкодження (+5) & Зменшено здоров'я (-5)"}
	}
	return buffs[rand.Intn(len(buffs))]
}

func getRandomWeapon() (string, int) {
	rows, err := db.Query("SELECT weaponType, damage FROM weapons ORDER BY RANDOM() LIMIT 1")
	if err != nil {
		log.Fatal("Error fetching weapon:", err)
	}
	defer rows.Close()

	var weaponType string
	var damage int
	if rows.Next() {
		err := rows.Scan(&weaponType, &damage)
		if err != nil {
			log.Fatal("Error scanning weapon row:", err)
		}
		return weaponType, damage
	}
	return "Fists", 2
}

func getRandomMonster() *Monster {
	rows, err := db.Query("SELECT id, monsterType, hp, damage FROM monsters ORDER BY RANDOM() LIMIT 1")
	if err != nil {
		log.Fatal("Error fetching monster:", err)
	}
	defer rows.Close()

	var monster Monster
	if rows.Next() {
		err := rows.Scan(&monster.ID, &monster.MonsterType, &monster.HP, &monster.Damage)
		if err != nil {
			log.Fatal("Error scanning monster row:", err)
		}
		return &monster
	}
	return nil
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
				endGame()
				return
			}

			time.Sleep(time.Second)
		}

		// Monster defeated, apply buffs
		fmt.Println(termenv.String(fmt.Sprintf("The %s has been defeated!", monster.MonsterType)).
			Foreground(termenv.ANSIGreen).Bold())

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
	player.HP = min(player.HP+15, 100)
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
		monster.HP = min(monster.HP+15, 200) // Fixed monster healing limit
		if lang == "en" {
			fmt.Println(termenv.String(fmt.Sprintf(" The %s heals! It now has %d HP.", monster.MonsterType, monster.HP)).Foreground(termenv.ANSIGreen))
		} else {
			fmt.Println(termenv.String(fmt.Sprintf(" Монстр вилікувався! Тепер він має %d здоров'я.", monster.HP)).Foreground(termenv.ANSIGreen))
		}
		*monsterDefending = false
	} else {
		monsterDamage := monster.Damage + rng()
		if *playerDefending {
			printDefendMessage("You blocked the enemy's attack!", "Ти заблокував атаку ворога!")
			*playerDefending = false // Reset defense after blocking
		} else {
			player.HP -= monsterDamage
			if lang == "en" {
				fmt.Println(termenv.String(fmt.Sprintf("󰓥 The %s attacks you for %d damage! You now have %d HP.", monster.MonsterType, monsterDamage, player.HP)).Foreground(termenv.ANSIRed))
			} else {
				fmt.Println(termenv.String(fmt.Sprintf("󰓥 Ти атакував %s з силою %d! Тепер в тебе %d здоров'я.", monster.MonsterType, monster.Damage, player.HP)).Foreground(termenv.ANSIBlue))
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
			Label: "Виберіть бонус або зброю (покращення)",
			Items: []string{baff1, baff2, "Випадкова зброя"},
		}
	} // <-- Moved closing brace here

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal("Prompt failed:", err)
	}

	// Apply buffs based on user selection
	if result == "Increase HP (+2) & Reduce Damage (-1)" || result == "Додано здоров'я (+2) & Зменшено пошкодження (-1)" {
		player.HP += 2
		if player.Damage > 1 {
			player.Damage -= 1
		} else {
			fmt.Println(termenv.String(" Damage cannot be reduced further!").Foreground(termenv.ANSIRed))
		}
		fmt.Println(termenv.String(fmt.Sprintf(" Buff Applied! HP: %d, Damage: %d", player.HP, player.Damage)).Foreground(termenv.ANSIGreen))

	} else if result == "Increase Damage (+5) & Reduce HP (-5)" || result == "Додано пошкодження (+5) & Зменшено здоров'я (-5)" {
		player.Damage += 5
		if player.HP > 5 {
			player.HP -= 5
		} else {
			player.HP = 0
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

func endGame() {
	_, err := db.Exec("DROP TABLE IF EXISTS monsters; DROP TABLE IF EXISTS weapons;")
	if err != nil {
		log.Fatal("Error cleaning up database:", err)
	}
}
