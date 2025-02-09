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

func seedData() {
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
	buffs := []string{"Increase HP (+2) & Reduce Damage (-1)", "Increase Damage (+5) & Reduce HP (-5)", "Random Weapon"}
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

func changeWeapon(player *Player) {
	rows, err := db.Query("SELECT weaponType, damage FROM weapons")
	if err != nil {
		log.Fatal("Error fetching weapons:", err)
	}
	defer rows.Close()

	var weapons []string
	var weaponMap = make(map[string]int)
	for rows.Next() {
		var weaponType string
		var damage int
		err := rows.Scan(&weaponType, &damage)
		if err != nil {
			log.Fatal("Error scanning weapon row:", err)
		}
		weapons = append(weapons, weaponType)
		weaponMap[weaponType] = damage
	}

	prompt := promptui.Select{
		Label: "Select a New Weapon",
		Items: weapons,
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal("Prompt failed:", err)
	}

	player.WeaponType = result
	player.Damage = weaponMap[result]
	fmt.Println(termenv.String(fmt.Sprintf(" Weapon Changed! You now wield a %s dealing %d damage.", player.WeaponType, player.Damage)).
		Foreground(termenv.ANSIGreen))
}

func fight(player *Player) {
	for player.HP > 0 {
		monster := getRandomMonster()
		if monster == nil {
			fmt.Println(termenv.String("No monsters found!").Foreground(termenv.ANSIYellow))
			return
		}

		// Display the monster and player information
		fmt.Println(termenv.String(fmt.Sprintf("A wild %s appears with %d HP!", monster.MonsterType, monster.HP)).
			Foreground(termenv.ANSICyan))
		fmt.Println(termenv.String(fmt.Sprintf("You wield a %s dealing %d damage and have %d HP.", player.WeaponType, player.Damage, player.HP)).
			Foreground(termenv.ANSIGreen))

		playerDefending := false
		monsterDefending := false

		for monster.HP > 0 && player.HP > 0 {
			playerAction := promptAction()

			if playerAction == "Defend" {
				fmt.Println(termenv.String(" You block the attack!").Foreground(termenv.ANSIYellow))
				playerDefending = true
			} else if playerAction == "Heal" {
				player.HP = min(player.HP+20, 100)
				fmt.Println(termenv.String(fmt.Sprintf(" You heal! Your HP is now %d.", player.HP)).Foreground(termenv.ANSIGreen))
				playerDefending = false
			} else if playerAction == "Attack" {
				playerDamage := player.Damage + rng()
				if monsterDefending {
					fmt.Println(termenv.String(fmt.Sprintf(" The %s blocked your attack!", monster.MonsterType)).
						Foreground(termenv.ANSIYellow))
					monsterDefending = false // Reset defense after blocking
				} else {
					monster.HP -= playerDamage
					fmt.Println(termenv.String(fmt.Sprintf("󰓥 You attack the %s for %d damage! It now has %d HP.", monster.MonsterType, playerDamage, monster.HP)).
						Foreground(termenv.ANSIBlue))
				}
			}

			monsterAction := enemyTurn(monster)

			if monsterAction == "Defend" {
				fmt.Println(termenv.String(fmt.Sprintf(" The %s prepares to block!", monster.MonsterType)).
					Foreground(termenv.ANSIYellow))
				monsterDefending = true
			} else if monsterAction == "Heal" {
				maxHP := 100 // Default max HP
				if monster.MonsterType == "Dragon" {
					maxHP = 150
				}
				monster.HP = min(monster.HP+15, maxHP)
				fmt.Println(termenv.String(fmt.Sprintf(" The %s heals! It now has %d HP.", monster.MonsterType, monster.HP)).
					Foreground(termenv.ANSIGreen))
				monsterDefending = false
			} else {
				monsterDamage := monster.Damage + rng()
				if playerDefending {
					fmt.Println(termenv.String(" You blocked the enemy's attack!").Foreground(termenv.ANSIYellow))
					playerDefending = false // Reset defense after blocking
				} else {
					player.HP -= monsterDamage
					fmt.Println(termenv.String(fmt.Sprintf("󰓥 The %s attacks you for %d damage! You now have %d HP.", monster.MonsterType, monsterDamage, player.HP)).
						Foreground(termenv.ANSIRed))
				}
			}

			if player.HP <= 0 {
				fmt.Println(termenv.String(" You died!").Foreground(termenv.ANSIBrightRed).Bold())
				endGame()
				return
			}

			time.Sleep(time.Second)
		}

		fmt.Println(termenv.String(fmt.Sprintf("The %s has been defeated!", monster.MonsterType)).
			Foreground(termenv.ANSIGreen).Bold())
		buffsAction(player)
	}
}

func buffsAction(player *Player) {

	baff1 := getRandomBuff()
	baff2 := getRandomBuff()
	baff3 := getRandomBuff()

	prompt := promptui.Select{
		Label: "Select a Buff/Weapon (Upgrade)",
		Items: []string{baff1, baff2, baff3},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal("Prompt failed: %v", err)
	}

	if result == "Increase HP (+2) & Reduce Damage (-1)" {
		player.HP += 2
		if player.Damage > 1 {
			player.Damage -= 1
		} else {
			fmt.Println(termenv.String(" Damage cannot be reduced further!").Foreground(termenv.ANSIRed))
		}
		fmt.Println(termenv.String(fmt.Sprintf(" Buff Applied! HP: %d, Damage: %d", player.HP, player.Damage)).
			Foreground(termenv.ANSIGreen))
	} else if result == "Increase Damage (+5) & Reduce HP (-5)" {
		player.Damage += 5
		if player.HP > 5 {
			player.HP -= 5
		} else {
			player.HP = 0
		}
		fmt.Println(termenv.String(fmt.Sprintf(" Buff Applied! Damage: %d, HP: %d", player.Damage, player.HP)).
			Foreground(termenv.ANSIGreen))
	} else if result == "Random Weapon" {
		changeWeapon(player)
	} else {
		fmt.Println(termenv.String(" No Buff Applied.").Foreground(termenv.ANSIYellow))
	}
}

func promptAction() string {
	prompt := promptui.Select{
		Label: "Select an action",
		Items: []string{"Attack", "Defend", "Heal"},
	}

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
