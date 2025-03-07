# ✨ ~ Lura *, CLI turn based roguelike*
![Preview](https://gachi.gay/pHLVC)

[All mobs, weapon and effect statistics](https://github.com/IwnuplyNotTyan/Lura/blob/main/STAT.md)

## Build
Note: build dep `go`. To play need only [Nerdfonts](https://www.nerdfonts.com/).

### Makepkg
```sh
makepkg -si
```

### Gnumake

```sh
sudo make install
```
Default binary path: `~/go/bin/`

### Build manually
```sh
go mod download
go build -o lura
```

## Mods
Mod folder in `~/.local/share/Lura/mods/`, all mods must be in lua

### Adding weapon & monster mod:
```lua
local monsterName = {
    en = "test monster",
    ua = "тестовий противник"
}

local weaponName = {
    en = "test weapon",
    ua = "тестова зброя"
}

local lang = lang or "en"

local monsterIdx = Monster.new(monsterName[lang], 200, 30)
Monster.setHP(monsterIdx, 250)

local weaponIdx = Weapon.new(weaponName[lang], 15, 20)
Weapon.setDamage(weaponIdx, 18)
```

### Removing weapon & monster mod:
```lua
local monsterName = {
    en = "Dragon",
    ua = "Дракон"
}

local weaponName = {
    en = "Axe",
    ua = "Сокира"
}

local lang = lang or "en"

local weaponToRemove = weaponName[lang]
local weaponResult = Weapon.removeByName(weaponToRemove)

local monsterToRemove = monsterName[lang]
local monsterResult = Monster.removeByName(monsterToRemove)
```

## Credits
UA translate by [Purple Sky](https://github.com/Osian-linux) and me

### Made with ❤️
