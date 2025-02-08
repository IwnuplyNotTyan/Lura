package main

import (
	"database/sql"
	"log"
	"math/rand"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func main() {
	db, err := sql.Open("sqlite3", "yasg.db")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS app (id INTEGER PRIMARY KEY, monsterType TEXT, hp INTEGER, damage INTEGER, weaponType TEXT, damage INTEGER)")
	if err != nil {
		log.Fatal(err)
	}
}

func rng() int {
	rand.Seed(time.Now().UnixNano())
	min, max := 1, 6
	randomNumber := rand.Intn(max-min+1) + min
	log.Println(randomNumber)
	return randomNumber
}

func Player(weaponType string, damage int, hp int) {
	stmt, err := db.Prepare("INSERT INTO app (weaponType, damage) VALUES (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(weaponType, damage, hp)
	if err != nil {
		log.Fatal(err)
	}
}

func addMonster(monsterType string, hp int, damage int) {
	stmt, err := db.Prepare("INSERT INTO app (monsterType, hp, damage) VALUES (?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(monsterType, hp, damage)
	if err != nil {
		log.Fatal(err)
	}
}

func infoMonster() {
	rows, err := db.Query("SELECT * FROM app")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var monsterType string
		var hp int
		var damage int
		err = rows.Scan(&id, &monsterType, &hp, &damage)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, monsterType, hp)
	}
}
