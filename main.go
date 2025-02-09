package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	"github.com/charmbracelet/log"
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

	// Creating tables if they don't exist
	createTables()

	// Seeding data
	seedData()

	// Initialize random seed (call only once)
	rand.Seed(time.Now().UnixNano())

	// Assign a random weapon to the player
	weaponType, weaponDamage := getRandomWeapon()
	player := Player{WeaponType: weaponType, Damage: weaponDamage * rng(), HP: 100}

	// Start the game
	attack(player)
}

// createTables ensures the tables exist before inserting data
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

// seedData populates the tables with initial data
func seedData() {
	addMonster("Dragon", 150, 30)
	addMonster("Human", 50, 5)
	addMonster("Cat", 5, 0)

	addWeapon("Sword", 10)
	addWeapon("Spear", 7)
	addWeapon("Axe", 12)
}

func rng() int {
	return rand.Intn(6) + 1 // Random number between 1 and 6
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
	return "Fists", 2 // Default weapon
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

func attack(player Player) {
	monster := getRandomMonster()
	if monster == nil {
		fmt.Println(termenv.String("No monsters found!").Foreground(termenv.ANSIYellow))
		return
	}

	fmt.Println(termenv.String(fmt.Sprintf(" A wild %s appears with %d HP!", monster.MonsterType, monster.HP)).
		Foreground(termenv.ANSICyan))

	fmt.Println(termenv.String(fmt.Sprintf(" You wield a %s dealing %d damage and have %d HP.", player.WeaponType, player.Damage, player.HP)).
		Foreground(termenv.ANSIGreen))

	for monster.HP > 0 && player.HP > 0 {
		// Player attacks
		playerDamage := player.Damage + rng()
		monster.HP -= playerDamage
		fmt.Println(termenv.String(fmt.Sprintf("󰓥 You hit the %s for %d damage! It now has %d HP left.", monster.MonsterType, playerDamage, monster.HP)).
			Foreground(termenv.ANSIBlue))

		if monster.HP <= 0 {
			fmt.Println(termenv.String(fmt.Sprintf(" You defeated the %s!", monster.MonsterType)).
				Foreground(termenv.ANSIGreen).Bold())
			endGame()
			return
		}

		// Monster attacks
		monsterDamage := monster.Damage + rng()
		player.HP -= monsterDamage
		fmt.Println(termenv.String(fmt.Sprintf("󰓥 The %s hits you for %d damage! You now have %d HP left.", monster.MonsterType, monsterDamage, player.HP)).
			Foreground(termenv.ANSIRed).Italic())

		if player.HP <= 0 {
			fmt.Println(termenv.String(" You died!").Foreground(termenv.ANSIBrightRed).Bold())
			endGame()
			return
		}

		time.Sleep(time.Second)
	}
}

func endGame() {
	_, err := db.Exec("DROP TABLE IF EXISTS monsters; DROP TABLE IF EXISTS weapons;")
	if err != nil {
		log.Fatal("Error cleaning up database:", err)
	}
}
