package main

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
	Coins      int
	Stamina    int
	maxStamina int
}

type Weapon struct {
	WeaponType string
	Damage     int
	Stamina    int
}

var (
	vmonsters  []Monster
	smonsters  []Monster
	ccmonsters []Monster
	weapons    []Weapon
	lang       string
)

func seedData() {
	if lang == "en" {
		vmonsters = []Monster{
			{MonsterType: "Dragon", HP: 150, Damage: 25},
			{MonsterType: "Human", HP: 70, Damage: 10},
			{MonsterType: "Ork", HP: 60, Damage: 15},
			{MonsterType: "Goblin", HP: 50, Damage: 10},
			{MonsterType: "Troll", HP: 80, Damage: 20},
			{MonsterType: "Warrior", HP: 100, Damage: 20},
			{MonsterType: "Golem", HP: 200, Damage: 10},
			{MonsterType: "Ogre", HP: 100, Damage: 20},
			{MonsterType: "Skeleton", HP: 60, Damage: 10},
			{MonsterType: "Zombie", HP: 70, Damage: 15},
		}
		ccmonsters = []Monster{
			{MonsterType: "CCPlaceholder", HP: 1, Damage: 1},
		}
		smonsters = []Monster{
			{MonsterType: "CCPlaceholder", HP: 1, Damage: 1},
		}
		weapons = []Weapon{
			{WeaponType: "Sword", Damage: 10, Stamina: 10},
			{WeaponType: "Spear", Damage: 9, Stamina: 7},
			{WeaponType: "Axe", Damage: 12, Stamina: 15},
			{WeaponType: "Longsword", Damage: 11, Stamina: 13},
			{WeaponType: "Dagger", Damage: 8, Stamina: 5},
			{WeaponType: "Crossbow", Damage: 9, Stamina: 11},
			{WeaponType: "Bow", Damage: 8, Stamina: 9},
		}
	} else if lang == "ua" {
		vmonsters = []Monster{
			{MonsterType: "Дракон", HP: 150, Damage: 30},
			{MonsterType: "Людина", HP: 70, Damage: 10},
			{MonsterType: "Орк", HP: 60, Damage: 15},
			{MonsterType: "Гоблін", HP: 50, Damage: 10},
			{MonsterType: "Троль", HP: 80, Damage: 20},
			{MonsterType: "Воїн", HP: 100, Damage: 20},
			{MonsterType: "Голем", HP: 200, Damage: 10},
			{MonsterType: "Огр", HP: 110, Damage: 25},
			{MonsterType: "Скелет", HP: 60, Damage: 10},
			{MonsterType: "Зомбі", HP: 70, Damage: 15},
		}
		ccmonsters = []Monster{
			{MonsterType: "CCPlaceholder", HP: 1, Damage: 1},
		}
		smonsters = []Monster{
			{MonsterType: "CCPlaceholder", HP: 1, Damage: 1},
		}
		weapons = []Weapon{
			{WeaponType: "Меч", Damage: 10, Stamina: 10},
			{WeaponType: "Спис", Damage: 9, Stamina: 7},
			{WeaponType: "Сокира", Damage: 12, Stamina: 15},
			{WeaponType: "Довгий Меч", Damage: 11, Stamina: 13},
			{WeaponType: "Кинджал", Damage: 8, Stamina: 5},
			{WeaponType: "Арбалет", Damage: 9, Stamina: 11},
			{WeaponType: "Лук", Damage: 8, Stamina: 9},
		}
	}
}
