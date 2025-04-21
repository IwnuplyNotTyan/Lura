package main

type Monster struct {
	ID          int
	MonsterType string
	HP          int
	Damage      int
	LVL         int
	maxHP       int
	score       int
	coins       int
	Position   int
}

type Player struct {
	WeaponType string
	Damage     int
	HP         int
	maxHP      int
	Coins      int
	Stamina    int
	maxStamina int
	heart      int
	buffs      int
	score      int
	loc        int
	monster    bool
	name       string
	time       int
	Inventory  Inventory
	Position   int
	Tmp        int
}

type Weapon struct {
	WeaponType string
	Damage     int
	Stamina    int
}

type Item struct {
	ID       int
	Name     string
	Quantity int
	Effect   string
	Value    int
	Price    int
}

type Inventory struct {
	Items  []Item
	NextID int
}

var (
	vmonsters  []Monster
	scmonsters []Monster
	boss   	   []Monster
	lanter     []Weapon
	musket     []Weapon
	weapons    []Weapon
	crossbow   []Weapon
	longsword  []Weapon
	lang       string
)

func seedData() {
	if lang == "en" {
		vmonsters = []Monster{
			{MonsterType: "Dragon", HP: 130, Damage: 25, score: 50, coins: 20},
			{MonsterType: "Ork", HP: 50, Damage: 15, score: 20, coins: 7},
			{MonsterType: "Goblin", HP: 40, Damage: 10, score: 15, coins: 5},
			{MonsterType: "Troll", HP: 70, Damage: 20, score: 35, coins: 9},
			{MonsterType: "Warrior", HP: 100, Damage: 20, score: 45, coins: 15},
			{MonsterType: "Ogre", HP: 80, Damage: 20, score: 40, coins: 14},
			{MonsterType: "Skeleton", HP: 50, Damage: 10, score: 20, coins: 6},
			{MonsterType: "Zombie", HP: 60, Damage: 15, score: 30, coins: 11},
			{MonsterType: "Musketeer", HP: 80, Damage: 30, score: 30, coins: 10},
		}
		weapons = []Weapon{
			{WeaponType: "Sword", Damage: 13, Stamina: 10},
			{WeaponType: "Spear", Damage: 12, Stamina: 7},
			{WeaponType: "Axe", Damage: 16, Stamina: 15},
			{WeaponType: "Dagger", Damage: 11, Stamina: 5},
			{WeaponType: "Bow", Damage: 11, Stamina: 9},
			{WeaponType: "Longbow", Damage: 11, Stamina: 9},
			{WeaponType: "Torch", Damage: 6, Stamina: 3},
		}
		lanter = []Weapon{
			{WeaponType: "Lanter of the soul", Damage: 5, Stamina: 5},
		}
		musket = []Weapon{
			{WeaponType: "Musket", Damage: 30, Stamina: 5},
		}
		longsword = []Weapon{
			{WeaponType: "Longsword", Damage: 15, Stamina: 13},
		}
		crossbow = []Weapon{
			{WeaponType: "Crossbow", Damage: 13, Stamina: 11},
		}
		scmonsters = []Monster{
			{MonsterType: "Crystal guardian", HP: 100, Damage: 20, score: 60, coins: 13},
			{MonsterType: "Golem", HP: 130, Damage: 10, score: 60, coins: 17},
			{MonsterType: "Miner", HP: 60, Damage: 20, score: 20, coins: 14},
			{MonsterType: "Aetherite titan", HP: 100, Damage: 20, score: 70, coins: 14},
			{MonsterType: "Lanter keeper", HP: 70, Damage: 10, score: 10, coins: 14},
			{MonsterType: "Ghost", HP: 40, Damage: 5, score: 4, coins: 10},
		}
		boss = []Monster{
			{MonsterType: "Colosus", HP: 200, Damage: 30, score: 100, coins: 50},
		}
	} else if lang == "be" {
		vmonsters = []Monster{
			{MonsterType: "Цмок", HP: 130, Damage: 25, score: 50, coins: 20},
			{MonsterType: "Чалавек", HP: 60, Damage: 10, score: 30, coins: 10},
			{MonsterType: "Орк", HP: 50, Damage: 15, score: 20, coins: 7},
			{MonsterType: "Гоблін", HP: 40, Damage: 10, score: 15, coins: 5},
			{MonsterType: "Троль", HP: 70, Damage: 20, score: 35, coins: 9},
			{MonsterType: "Воін", HP: 100, Damage: 20, score: 45, coins: 15},
			{MonsterType: "Огр", HP: 80, Damage: 20, score: 40, coins: 14},
			{MonsterType: "Шкілет", HP: 50, Damage: 10, score: 20, coins: 6},
			{MonsterType: "Зомбі", HP: 60, Damage: 15, score: 30, coins: 11},
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
			{MonsterType: "Дракон", HP: 130, Damage: 30, score: 50, coins: 20},
			{MonsterType: "Людина", HP: 60, Damage: 10, score: 30, coins: 10},
			{MonsterType: "Орк", HP: 50, Damage: 15, score: 20, coins: 7},
			{MonsterType: "Гоблін", HP: 40, Damage: 10, score: 15, coins: 5},
			{MonsterType: "Троль", HP: 70, Damage: 20, score: 35, coins: 9},
			{MonsterType: "Воїн", HP: 100, Damage: 20, score: 45, coins: 15},
			{MonsterType: "Огр", HP: 90, Damage: 25, score: 40, coins: 14},
			{MonsterType: "Скелет", HP: 50, Damage: 10, score: 20, coins: 6},
			{MonsterType: "Зомбі", HP: 60, Damage: 15, score: 30, coins: 11},
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
