package main

type Monster struct {
	ID          int
	MonsterType string
	HP          int
	Damage      int
	LVL         int
	maxHP       int
	score       int
}

type Player struct {
	WeaponType string
	Damage     int
	HP         int
	maxHP      int
	Coins      int
	Stamina    int
	maxStamina int
	heart      bool
	buffs      int
	score      int
	loc        int
}

type Weapon struct {
	WeaponType string
	Damage     int
	Stamina    int
}

var (
	vmonsters  []Monster
	scmonsters []Monster
	weapons    []Weapon
	lang       string
)

func seedData() {
	if lang == "en" {
		vmonsters = []Monster{
			{MonsterType: "Dragon", HP: 130, Damage: 25, score: 50},
			{MonsterType: "Human", HP: 60, Damage: 10, score: 30},
			{MonsterType: "Ork", HP: 50, Damage: 15, score: 20},
			{MonsterType: "Goblin", HP: 40, Damage: 10, score: 15},
			{MonsterType: "Troll", HP: 70, Damage: 20, score: 35},
			{MonsterType: "Warrior", HP: 100, Damage: 20, score: 45},
			{MonsterType: "Ogre", HP: 80, Damage: 20, score: 40},
			{MonsterType: "Skeleton", HP: 50, Damage: 10, score: 20},
			{MonsterType: "Zombie", HP: 60, Damage: 15, score: 30},
		}
		weapons = []Weapon{
			{WeaponType: "Sword", Damage: 13, Stamina: 10},
			{WeaponType: "Spear", Damage: 12, Stamina: 7},
			{WeaponType: "Axe", Damage: 16, Stamina: 15},
			{WeaponType: "Longsword", Damage: 15, Stamina: 13},
			{WeaponType: "Dagger", Damage: 11, Stamina: 5},
			{WeaponType: "Crossbow", Damage: 12, Stamina: 11},
			{WeaponType: "Bow", Damage: 11, Stamina: 9},
		}
		scmonsters = []Monster{
			{MonsterType: "Crystal guardian", HP: 100, Damage: 20, score: 60},
			{MonsterType: "Golem", HP: 130, Damage: 10, score: 60},
			{MonsterType: "Miner", HP: 60, Damage: 20, score: 20},
			{MonsterType: "Aetherite titan", HP: 100, Damage: 20, score: 70},
		}
	} else if lang == "be" {
		vmonsters = []Monster{
			{MonsterType: "Цмок", HP: 130, Damage: 25, score: 50},
			{MonsterType: "Чалавек", HP: 60, Damage: 10, score: 30},
			{MonsterType: "Орк", HP: 50, Damage: 15, score: 20},
			{MonsterType: "Гоблін", HP: 40, Damage: 10, score: 15},
			{MonsterType: "Троль", HP: 70, Damage: 20, score: 35},
			{MonsterType: "Воін", HP: 100, Damage: 20, score: 45},
			{MonsterType: "Голем", HP: 150, Damage: 10, score: 60},
			{MonsterType: "Огр", HP: 80, Damage: 20, score: 40},
			{MonsterType: "Шкілет", HP: 50, Damage: 10, score: 20},
			{MonsterType: "Зомбі", HP: 60, Damage: 15, score: 30},
		}
		weapons = []Weapon{
			{WeaponType: "Меч", Damage: 13, Stamina: 10},
			{WeaponType: "Дзіда", Damage: 12, Stamina: 7},
			{WeaponType: "Сякера", Damage: 16, Stamina: 15},
			{WeaponType: "Доўгі меч", Damage: 15, Stamina: 13},
			{WeaponType: "Кінжал", Damage: 11, Stamina: 5},
			{WeaponType: "Арбалет", Damage: 12, Stamina: 11},
			{WeaponType: "Лук", Damage: 11, Stamina: 9},
		}
	} else if lang == "ua" {
		vmonsters = []Monster{
			{MonsterType: "Дракон", HP: 130, Damage: 30, score: 50},
			{MonsterType: "Людина", HP: 60, Damage: 10, score: 30},
			{MonsterType: "Орк", HP: 50, Damage: 15, score: 20},
			{MonsterType: "Гоблін", HP: 40, Damage: 10, score: 15},
			{MonsterType: "Троль", HP: 70, Damage: 20, score: 35},
			{MonsterType: "Воїн", HP: 100, Damage: 20, score: 45},
			{MonsterType: "Голем", HP: 150, Damage: 10, score: 60},
			{MonsterType: "Огр", HP: 90, Damage: 25, score: 40},
			{MonsterType: "Скелет", HP: 50, Damage: 10, score: 20},
			{MonsterType: "Зомбі", HP: 60, Damage: 15, score: 30},
		}
		weapons = []Weapon{
			{WeaponType: "Меч", Damage: 13, Stamina: 10},
			{WeaponType: "Спис", Damage: 12, Stamina: 7},
			{WeaponType: "Сокира", Damage: 16, Stamina: 15},
			{WeaponType: "Довгий Меч", Damage: 15, Stamina: 13},
			{WeaponType: "Кинджал", Damage: 11, Stamina: 5},
			{WeaponType: "Арбалет", Damage: 12, Stamina: 11},
			{WeaponType: "Лук", Damage: 11, Stamina: 9},
		}
	}
}
