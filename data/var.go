package data

type Monster struct {
	ID          int
	MonsterType string
	HP          int
	Damage      int
	LVL         int
	MaxHP       int
	Score       int
	Coins       int
	Position    int
}


type Player struct {
	WeaponType string
	Damage     int
	HP         int
	MaxHP      int
	Coins      int
	Stamina    int
	MaxStamina int
	Heart      int
	Buffs      int
	Score      int
	Loc        int
	Monster    bool
	Name       string
	Time       int
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
	Vmonsters  []Monster
	Scmonsters []Monster
	Boss   	   []Monster
	Lanter     []Weapon
	Musket     []Weapon
	Weapons    []Weapon
	Crossbow   []Weapon
	Longsword  []Weapon
	Lang       string
	Tmp        int
	Verbose    bool
)
