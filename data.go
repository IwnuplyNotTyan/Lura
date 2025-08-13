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
	Position    int
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
	WeaponID   int
}

type Weapon struct {
	WeaponType string
	Damage     int
	Stamina    int
	ID         int
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
	tmp        int
)

func seedData() {
	if lang == "en" {
		vmonsters = []Monster{
			{MonsterType: "Dragon", HP: 130, Damage: 25, score: 50, coins: 20, ID: 3},
			{MonsterType: "Ork", HP: 50, Damage: 15, score: 20, coins: 7, ID: 4},
			{MonsterType: "Goblin", HP: 40, Damage: 10, score: 15, coins: 5, ID: 5},
			{MonsterType: "Troll", HP: 70, Damage: 20, score: 35, coins: 9, ID: 6},
			{MonsterType: "Warrior", HP: 100, Damage: 20, score: 45, coins: 15, ID: 7},
			{MonsterType: "Ogre", HP: 80, Damage: 20, score: 40, coins: 14, ID: 8},
			{MonsterType: "Skeleton", HP: 50, Damage: 10, score: 20, coins: 6, ID: 9},
			{MonsterType: "Zombie", HP: 60, Damage: 15, score: 30, coins: 11, ID: 10},
			{MonsterType: "Musketeer", HP: 80, Damage: 30, score: 30, coins: 10, ID: 1},
		}
		weapons = []Weapon{
			{WeaponType: "Sword", Damage: 13, Stamina: 10, ID: 1},
			{WeaponType: "Spear", Damage: 12, Stamina: 7, ID: 2},
			{WeaponType: "Axe", Damage: 16, Stamina: 15, ID: 3},
			{WeaponType: "Dagger", Damage: 11, Stamina: 5, ID: 4},
			{WeaponType: "Bow", Damage: 11, Stamina: 9, ID: 5},
			{WeaponType: "Longbow", Damage: 11, Stamina: 9, ID: 6},
			//{WeaponType: "Torch", Damage: 6, Stamina: 3},
		}
		lanter = []Weapon{
			{WeaponType: "Mirror", Damage: 5, Stamina: 5, ID: 7},
		}
		musket = []Weapon{
			{WeaponType: "Musket", Damage: 30, Stamina: 5, ID: 8},
		}
		longsword = []Weapon{
			{WeaponType: "Longsword", Damage: 15, Stamina: 13, ID: 9},
		}
		crossbow = []Weapon{
			{WeaponType: "Crossbow", Damage: 13, Stamina: 11, ID: 10},
		}
		scmonsters = []Monster{
			{MonsterType: "Crystal guardian", HP: 100, Damage: 20, score: 60, coins: 13, ID: 11},
			{MonsterType: "Golem", HP: 130, Damage: 10, score: 60, coins: 17, ID: 12},
			{MonsterType: "Miner", HP: 60, Damage: 20, score: 20, coins: 14, ID: 13},
			{MonsterType: "Aetherite titan", HP: 100, Damage: 20, score: 70, coins: 14},
		//	{MonsterType: "Lanter keeper", HP: 70, Damage: 10, score: 10, coins: 14, ID: 2},
			{MonsterType: "Ghost", HP: 40, Damage: 5, score: 4, coins: 10, ID: 15},
		}
		boss = []Monster{
			//{MonsterType: "Colosus", HP: 200, Damage: 30, score: 100, coins: 50, ID: 16},
			{MonsterType: "Gato", HP: 300, Damage: 40, score: 150, coins: 100, ID: 17},
		}
	} else if lang == "be" {
		vmonsters = []Monster{
			{MonsterType: "Цмок", HP: 130, Damage: 25, score: 50, coins: 20, ID: 3},
			{MonsterType: "Орк", HP: 50, Damage: 15, score: 20, coins: 7, ID: 4},
			{MonsterType: "Гоблін", HP: 40, Damage: 10, score: 15, coins: 5, ID: 5},
			{MonsterType: "Троль", HP: 70, Damage: 20, score: 35, coins: 9, ID: 6},
			{MonsterType: "Воін", HP: 100, Damage: 20, score: 45, coins: 15, ID: 7},
			{MonsterType: "Огр", HP: 80, Damage: 20, score: 40, coins: 14, ID: 8},
			{MonsterType: "Шкілет", HP: 50, Damage: 10, score: 20, coins: 6, ID: 9},
			{MonsterType: "Зомбі", HP: 60, Damage: 15, score: 30, coins: 11, ID: 10},
			{MonsterType: "Мушкетэр", HP: 80, Damage: 30, score: 30, coins: 10, ID: 1},
		}
		scmonsters = []Monster{
			{MonsterType: "Крыштальны ахоўнік", HP: 100, Damage: 20, score: 60, coins: 13, ID: 11},
			{MonsterType: "Голем", HP: 130, Damage: 10, score: 60, coins: 17, ID: 12},
			{MonsterType: "Шахцёр", HP: 60, Damage: 20, score: 20, coins: 14, ID: 13},
			{MonsterType: "Эфірны тытан", HP: 100, Damage: 20, score: 70, coins: 14, ID: 14},
		//	{MonsterType: "Ахоўца ліхтара", HP: 70, Damage: 10, score: 10, coins: 14, ID: 2},
			{MonsterType: "Прывід", HP: 40, Damage: 5, score: 4, coins: 10, ID: 15},
		}
		weapons = []Weapon{
			{WeaponType: "Меч", Damage: 13, Stamina: 10, ID: 1},
			{WeaponType: "Дзіда", Damage: 12, Stamina: 7, ID: 2},
			{WeaponType: "Сякера", Damage: 16, Stamina: 15, ID: 3},
			{WeaponType: "Кінжал", Damage: 11, Stamina: 5, ID: 4},
			{WeaponType: "Лук", Damage: 11, Stamina: 9, ID: 5},
			{WeaponType: "Доўгі лук", Damage: 11, Stamina: 9, ID: 6},
		}
		musket = []Weapon{
			{WeaponType: "Мушкет", Damage: 30, Stamina: 5, ID: 8},
		}
		crossbow = []Weapon{
			{WeaponType: "Арбалет", Damage: 13, Stamina: 11, ID: 10},
		}
		lanter = []Weapon{
			{WeaponType: "Люстэрка", Damage: 5, Stamina: 5, ID: 7},
		}
		longsword = []Weapon{
			{WeaponType: "Доўгі меч", Damage: 15, Stamina: 13, ID: 9},
		}
		boss = []Monster{
			{MonsterType: "Гато", HP: 300, Damage: 40, score: 150, coins: 100, ID: 17},
		}
	} else if lang == "ua" {
		vmonsters = []Monster{
			{MonsterType: "Дракон", HP: 130, Damage: 30, score: 50, coins: 20, ID: 3},
			{MonsterType: "Орк", HP: 50, Damage: 15, score: 20, coins: 7, ID: 4},
			{MonsterType: "Гоблін", HP: 40, Damage: 10, score: 15, coins: 5, ID: 5},
			{MonsterType: "Троль", HP: 70, Damage: 20, score: 35, coins: 9, ID: 6},
			{MonsterType: "Воїн", HP: 100, Damage: 20, score: 45, coins: 15, ID: 7},
			{MonsterType: "Огр", HP: 90, Damage: 25, score: 40, coins: 14, ID: 8},
			{MonsterType: "Скелет", HP: 50, Damage: 10, score: 20, coins: 6, ID: 9},
			{MonsterType: "Зомбі", HP: 60, Damage: 15, score: 30, coins: 11, ID: 10},
			{MonsterType: "Мушкетер", HP: 80, Damage: 30, score: 30, coins: 10, ID: 1},
		}
		scmonsters = []Monster{
			{MonsterType: "Кристальний охоронець", HP: 100, Damage: 20, score: 60, coins: 13, ID: 11},
			{MonsterType: "Голем", HP: 130, Damage: 10, score: 60, coins: 17, ID: 12},
			{MonsterType: "Шахтар", HP: 60, Damage: 20, score: 20, coins: 14, ID: 13},
			{MonsterType: "Аетеритний титан", HP: 100, Damage: 20, score: 70, coins: 14, ID: 14},
		//	{MonsterType: "Охоронець ліхтаря", HP: 70, Damage: 10, score: 10, coins: 14, ID: 2},
			{MonsterType: "Привид", HP: 40, Damage: 5, score: 4, coins: 10, ID: 15},
		}
		boss = []Monster{
			{MonsterType: "Гато", HP: 300, Damage: 40, score: 150, coins: 100, ID: 17},
		}
		weapons = []Weapon{
			{WeaponType: "Меч", Damage: 13, Stamina: 10, ID: 1},
			{WeaponType: "Спис", Damage: 12, Stamina: 7, ID: 2},
			{WeaponType: "Сокира", Damage: 16, Stamina: 15, ID: 3},
			{WeaponType: "Кинджал", Damage: 11, Stamina: 5, ID: 4},
			{WeaponType: "Лук", Damage: 11, Stamina: 9, ID: 5},
			{WeaponType: "Довгий лук", Damage: 11, Stamina: 9, ID: 6},
		}
		musket = []Weapon{
			{WeaponType: "Мушкет", Damage: 30, Stamina: 5, ID: 8},
		}
		crossbow = []Weapon{
			{WeaponType: "Арбалет", Damage: 13, Stamina: 11, ID: 10},
		}
		lanter = []Weapon{
			{WeaponType: "Зеркало", Damage: 5, Stamina: 5, ID: 7},
		}
		longsword = []Weapon{
			{WeaponType: "Довгий меч", Damage: 15, Stamina: 13, ID: 9},
		}
	}
}
