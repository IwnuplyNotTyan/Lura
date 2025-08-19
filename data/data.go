package data

func SeedData() {
	if Lang == "en" {
		Vmonsters = []Monster{
			{MonsterType: "Dragon", HP: 130, Damage: 25, Score: 50, Coins: 20, ID: 3},
			{MonsterType: "Ork", HP: 50, Damage: 15, Score: 20, Coins: 7, ID: 4},
			{MonsterType: "Goblin", HP: 40, Damage: 10, Score: 15, Coins: 5, ID: 5},
			{MonsterType: "Troll", HP: 70, Damage: 20, Score: 35, Coins: 9, ID: 6},
			{MonsterType: "Warrior", HP: 100, Damage: 20, Score: 45, Coins: 15, ID: 7},
			{MonsterType: "Ogre", HP: 80, Damage: 20, Score: 40, Coins: 14, ID: 8},
			{MonsterType: "Skeleton", HP: 50, Damage: 10, Score: 20, Coins: 6, ID: 9},
			{MonsterType: "Zombie", HP: 60, Damage: 15, Score: 30, Coins: 11, ID: 10},
			{MonsterType: "Musketeer", HP: 80, Damage: 30, Score: 30, Coins: 10, ID: 1},
		}
		Weapons = []Weapon{
			{WeaponType: "Sword", Damage: 13, Stamina: 10, ID: 1},
			{WeaponType: "Spear", Damage: 12, Stamina: 7, ID: 2},
			{WeaponType: "Axe", Damage: 16, Stamina: 15, ID: 3},
			{WeaponType: "Dagger", Damage: 11, Stamina: 5, ID: 4},
			{WeaponType: "Bow", Damage: 11, Stamina: 9, ID: 5},
			{WeaponType: "Longbow", Damage: 11, Stamina: 9, ID: 6},
			//{WeaponType: "Torch", Damage: 6, Stamina: 3},
		}
		Lanter = []Weapon{
			{WeaponType: "Mirror", Damage: 5, Stamina: 5, ID: 7},
		}
		Musket = []Weapon{
			{WeaponType: "Musket", Damage: 30, Stamina: 5, ID: 8},
		}
		Longsword = []Weapon{
			{WeaponType: "Longsword", Damage: 15, Stamina: 13, ID: 9},
		}
		Crossbow = []Weapon{
			{WeaponType: "Crossbow", Damage: 13, Stamina: 11, ID: 10},
		}
		Scmonsters = []Monster{
			{MonsterType: "Crystal guardian", HP: 100, Damage: 20, Score: 60, Coins: 13, ID: 11},
			{MonsterType: "Golem", HP: 130, Damage: 10, Score: 60, Coins: 17, ID: 12},
			{MonsterType: "Miner", HP: 60, Damage: 20, Score: 20, Coins: 14, ID: 13},
			{MonsterType: "Aetherite titan", HP: 100, Damage: 20, Score: 70, Coins: 14},
		//	{MonsterType: "Lanter keeper", HP: 70, Damage: 10, score: 10, coins: 14, ID: 2},
			{MonsterType: "Ghost", HP: 40, Damage: 5, Score: 4, Coins: 10, ID: 15},
		}
		Boss = []Monster{
			//{MonsterType: "Colosus", HP: 200, Damage: 30, score: 100, coins: 50, ID: 16},
			{MonsterType: "Gato", HP: 300, Damage: 40, Score: 150, Coins: 100, ID: 17},
		}
	} else if Lang == "be" {
		Vmonsters = []Monster{
			{MonsterType: "Цмок", HP: 130, Damage: 25, Score: 50, Coins: 20, ID: 3},
			{MonsterType: "Орк", HP: 50, Damage: 15, Score: 20, Coins: 7, ID: 4},
			{MonsterType: "Гоблін", HP: 40, Damage: 10, Score: 15, Coins: 5, ID: 5},
			{MonsterType: "Троль", HP: 70, Damage: 20, Score: 35, Coins: 9, ID: 6},
			{MonsterType: "Воін", HP: 100, Damage: 20, Score: 45, Coins: 15, ID: 7},
			{MonsterType: "Огр", HP: 80, Damage: 20, Score: 40, Coins: 14, ID: 8},
			{MonsterType: "Шкілет", HP: 50, Damage: 10, Score: 20, Coins: 6, ID: 9},
			{MonsterType: "Зомбі", HP: 60, Damage: 15, Score: 30, Coins: 11, ID: 10},
			{MonsterType: "Мушкетэр", HP: 80, Damage: 30, Score: 30, Coins: 10, ID: 1},
		}
		Scmonsters = []Monster{
			{MonsterType: "Крыштальны ахоўнік", HP: 100, Damage: 20, Score: 60, Coins: 13, ID: 11},
			{MonsterType: "Голем", HP: 130, Damage: 10, Score: 60, Coins: 17, ID: 12},
			{MonsterType: "Шахцёр", HP: 60, Damage: 20, Score: 20, Coins: 14, ID: 13},
			{MonsterType: "Эфірны тытан", HP: 100, Damage: 20, Score: 70, Coins: 14, ID: 14},
		//	{MonsterType: "Ахоўца ліхтара", HP: 70, Damage: 10, score: 10, coins: 14, ID: 2},
			{MonsterType: "Прывід", HP: 40, Damage: 5, Score: 4, Coins: 10, ID: 15},
		}
		Weapons = []Weapon{
			{WeaponType: "Меч", Damage: 13, Stamina: 10, ID: 1},
			{WeaponType: "Дзіда", Damage: 12, Stamina: 7, ID: 2},
			{WeaponType: "Сякера", Damage: 16, Stamina: 15, ID: 3},
			{WeaponType: "Кінжал", Damage: 11, Stamina: 5, ID: 4},
			{WeaponType: "Лук", Damage: 11, Stamina: 9, ID: 5},
			{WeaponType: "Доўгі лук", Damage: 11, Stamina: 9, ID: 6},
		}
		Musket = []Weapon{
			{WeaponType: "Мушкет", Damage: 30, Stamina: 5, ID: 8},
		}
		Crossbow = []Weapon{
			{WeaponType: "Арбалет", Damage: 13, Stamina: 11, ID: 10},
		}
		Lanter = []Weapon{
			{WeaponType: "Люстэрка", Damage: 5, Stamina: 5, ID: 7},
		}
		Longsword = []Weapon{
			{WeaponType: "Доўгі меч", Damage: 15, Stamina: 13, ID: 9},
		}
		Boss = []Monster{
			{MonsterType: "Гато", HP: 300, Damage: 40, Score: 150, Coins: 100, ID: 17},
		}
	} else if Lang == "ua" {
		Vmonsters = []Monster{
			{MonsterType: "Дракон", HP: 130, Damage: 30, Score: 50, Coins: 20, ID: 3},
			{MonsterType: "Орк", HP: 50, Damage: 15, Score: 20, Coins: 7, ID: 4},
			{MonsterType: "Гоблін", HP: 40, Damage: 10, Score: 15, Coins: 5, ID: 5},
			{MonsterType: "Троль", HP: 70, Damage: 20, Score: 35, Coins: 9, ID: 6},
			{MonsterType: "Воїн", HP: 100, Damage: 20, Score: 45, Coins: 15, ID: 7},
			{MonsterType: "Огр", HP: 90, Damage: 25, Score: 40, Coins: 14, ID: 8},
			{MonsterType: "Скелет", HP: 50, Damage: 10, Score: 20, Coins: 6, ID: 9},
			{MonsterType: "Зомбі", HP: 60, Damage: 15, Score: 30, Coins: 11, ID: 10},
			{MonsterType: "Мушкетер", HP: 80, Damage: 30, Score: 30, Coins: 10, ID: 1},
		}
		Scmonsters = []Monster{
			{MonsterType: "Кристальний охоронець", HP: 100, Damage: 20, Score: 60, Coins: 13, ID: 11},
			{MonsterType: "Голем", HP: 130, Damage: 10, Score: 60, Coins: 17, ID: 12},
			{MonsterType: "Шахтар", HP: 60, Damage: 20, Score: 20, Coins: 14, ID: 13},
			{MonsterType: "Аетеритний титан", HP: 100, Damage: 20, Score: 70, Coins: 14, ID: 14},
		//	{MonsterType: "Охоронець ліхтаря", HP: 70, Damage: 10, score: 10, coins: 14, ID: 2},
			{MonsterType: "Привид", HP: 40, Damage: 5, Score: 4, Coins: 10, ID: 15},
		}
		Boss = []Monster{
			{MonsterType: "Гато", HP: 300, Damage: 40, Score: 150, Coins: 100, ID: 17},
		}
		Weapons = []Weapon{
			{WeaponType: "Меч", Damage: 13, Stamina: 10, ID: 1},
			{WeaponType: "Спис", Damage: 12, Stamina: 7, ID: 2},
			{WeaponType: "Сокира", Damage: 16, Stamina: 15, ID: 3},
			{WeaponType: "Кинджал", Damage: 11, Stamina: 5, ID: 4},
			{WeaponType: "Лук", Damage: 11, Stamina: 9, ID: 5},
			{WeaponType: "Довгий лук", Damage: 11, Stamina: 9, ID: 6},
		}
		Musket = []Weapon{
			{WeaponType: "Мушкет", Damage: 30, Stamina: 5, ID: 8},
		}
		Crossbow = []Weapon{
			{WeaponType: "Арбалет", Damage: 13, Stamina: 11, ID: 10},
		}
		Lanter = []Weapon{
			{WeaponType: "Зеркало", Damage: 5, Stamina: 5, ID: 7},
		}
		Longsword = []Weapon{
			{WeaponType: "Довгий меч", Damage: 15, Stamina: 13, ID: 9},
		}
	}
}
