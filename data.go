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
	heart      bool
}

type Weapon struct {
	WeaponType string
	Damage     int
	Stamina    int
}

var (
	vmonsters []Monster
	//smonsters  []Monster
	//ccmonsters []Monster
	weapons []Weapon
	lang    string
)

func seedData() {
	if lang == "en" {
		vmonsters = []Monster{
			{MonsterType: "Dragon", HP: 130, Damage: 25},
			{MonsterType: "Human", HP: 60, Damage: 10},
			{MonsterType: "Ork", HP: 50, Damage: 15},
			{MonsterType: "Goblin", HP: 40, Damage: 10},
			{MonsterType: "Troll", HP: 70, Damage: 20},
			{MonsterType: "Warrior", HP: 100, Damage: 20},
			{MonsterType: "Golem", HP: 150, Damage: 10},
			{MonsterType: "Ogre", HP: 80, Damage: 20},
			{MonsterType: "Skeleton", HP: 50, Damage: 10},
			{MonsterType: "Zombie", HP: 60, Damage: 15},
		}
		weapons = []Weapon{
			{WeaponType: "Sword", Damage: 12, Stamina: 10},
			{WeaponType: "Spear", Damage: 11, Stamina: 7},
			{WeaponType: "Axe", Damage: 15, Stamina: 15},
			{WeaponType: "Longsword", Damage: 14, Stamina: 13},
			{WeaponType: "Dagger", Damage: 10, Stamina: 5},
			{WeaponType: "Crossbow", Damage: 11, Stamina: 11},
			{WeaponType: "Bow", Damage: 10, Stamina: 9},
		}
	} else if lang == "ua" {
		vmonsters = []Monster{
			{MonsterType: "Дракон", HP: 130, Damage: 30},
			{MonsterType: "Людина", HP: 60, Damage: 10},
			{MonsterType: "Орк", HP: 50, Damage: 15},
			{MonsterType: "Гоблін", HP: 40, Damage: 10},
			{MonsterType: "Троль", HP: 70, Damage: 20},
			{MonsterType: "Воїн", HP: 100, Damage: 20},
			{MonsterType: "Голем", HP: 150, Damage: 10},
			{MonsterType: "Огр", HP: 90, Damage: 25},
			{MonsterType: "Скелет", HP: 50, Damage: 10},
			{MonsterType: "Зомбі", HP: 60, Damage: 15},
		}
		weapons = []Weapon{
			{WeaponType: "Меч", Damage: 12, Stamina: 10},
			{WeaponType: "Спис", Damage: 11, Stamina: 7},
			{WeaponType: "Сокира", Damage: 15, Stamina: 15},
			{WeaponType: "Довгий Меч", Damage: 14, Stamina: 13},
			{WeaponType: "Кинджал", Damage: 10, Stamina: 5},
			{WeaponType: "Арбалет", Damage: 11, Stamina: 11},
			{WeaponType: "Лук", Damage: 10, Stamina: 9},
		}
	}
}
