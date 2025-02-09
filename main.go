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

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS app (id INTEGER PRIMARY KEY, monsterType TEXT, hp INTEGER, damage INTEGER, weaponType TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	// Example usage:
	addMonster("dragon", 150, 30)
	addMonster("human", 50, 5)
	addMonster("cat", 5, 0)

	player := Player{WeaponType: "sword", Damage: 10 * rng(), HP: 100}
	attack(player)
}

func rng() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1 // Random number between 1 and 6
}

func addMonster(monsterType string, hp int, damage int) {
	stmt, err := db.Prepare("INSERT INTO app (monsterType, hp, damage) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(monsterType, hp, damage)
	if err != nil {
		log.Fatal(err)
	}
}

func getRandomMonster() *Monster {
	rows, err := db.Query("SELECT id, monsterType, hp, damage FROM app ORDER BY RANDOM() LIMIT 1")
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

	for monster.HP > 0 {
		damage := player.Damage + rng()
		monster.HP -= damage
		fmt.Printf("You hit the %s for %d damage! It now has %d HP left.\n", monster.MonsterType, damage, monster.HP)
		if monster.HP <= 0 {
			fmt.Printf("You defeated the %s!\n", monster.MonsterType)
			return
		}
		time.Sleep(time.Second)
	}
}
