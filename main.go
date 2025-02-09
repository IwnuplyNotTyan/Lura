package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

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
		log.Fatal(err)
	}
	defer db.Close()

	// Creating separate tables for monsters and weapons
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS monsters (id INTEGER PRIMARY KEY, monsterType TEXT, hp INTEGER, damage INTEGER)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS weapons (id INTEGER PRIMARY KEY, weaponType TEXT, damage INTEGER)")
	if err != nil {
		log.Fatal(err)
	}

	addMonster("dragon", 150, 30)
	addMonster("human", 50, 5)
	addMonster("cat", 5, 0)

	addWeapon("sword", 10)
	addWeapon("spear", 7)
	addWeapon("axe", 12)

	// Assigning a random weapon to the player
	weaponType, weaponDamage := getRandomWeapon()
	player := Player{WeaponType: weaponType, Damage: weaponDamage * rng(), HP: 100}

	attack(player)
}

func rng() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1 // Random number between 1 and 6
}

func addWeapon(weaponType string, damage int) {
	stmt, err := db.Prepare("INSERT INTO weapons (weaponType, damage) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(weaponType, damage)
	if err != nil {
		log.Fatal(err)
	}
}

func addMonster(monsterType string, hp int, damage int) {
	stmt, err := db.Prepare("INSERT INTO monsters (monsterType, hp, damage) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(monsterType, hp, damage)
	if err != nil {
		log.Fatal(err)
	}
}

func getRandomWeapon() (string, int) {
	rows, err := db.Query("SELECT weaponType, damage FROM weapons ORDER BY RANDOM() LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var weaponType string
	var damage int
	if rows.Next() {
		rows.Scan(&weaponType, &damage)
		return weaponType, damage
	}
	return "fists", 2 // Default weapon if no weapons exist
}

func getRandomMonster() *Monster {
	rows, err := db.Query("SELECT id, monsterType, hp, damage FROM monsters ORDER BY RANDOM() LIMIT 1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var monster Monster
	if rows.Next() {
		rows.Scan(&monster.ID, &monster.MonsterType, &monster.HP, &monster.Damage)
		return &monster
	}
	return nil
}

func attack(player Player) {
	monster := getRandomMonster()
	if monster == nil {
		fmt.Println("No monsters found!")
		return
	}

	fmt.Printf("A wild %s appears with %d HP!\n", monster.MonsterType, monster.HP)
	fmt.Printf("You wield a %s dealing %d damage and have %d HP.\n", player.WeaponType, player.Damage, player.HP)

	for monster.HP > 0 && player.HP > 0 {
		// Player attacks
		playerDamage := player.Damage + rng()
		monster.HP -= playerDamage
		fmt.Printf("You hit the %s for %d damage! It now has %d HP left.\n", monster.MonsterType, playerDamage, monster.HP)

		if monster.HP <= 0 {
			fmt.Printf("You defeated the %s!\n", monster.MonsterType)
			return
		}

		// Monster attacks
		monsterDamage := monster.Damage + rng()
		player.HP -= monsterDamage
		fmt.Printf("The %s hits you for %d damage! You now have %d HP left.\n", monster.MonsterType, monsterDamage, player.HP)

		if player.HP <= 0 {
			fmt.Println("You died!")
			return
		}

		time.Sleep(time.Second)
	}
}
