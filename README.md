# ✨ ~ Lura *, CLI turn based roguelike*

![Preview](https://gachi.gay/pHLVC)

# Tree
- [Installing](https://github.com/IwnuplyNotTyan/Lura?tab=readme-ov-file#build)
  - [Arch-Based](https://github.com/IwnuplyNotTyan/Lura?tab=readme-ov-file#makepkg)
  - Void
    - [Xbps-install](https://github.com/IwnuplyNotTyan/Lura?tab=readme-ov-file#xbps-install)
    - [Xbps-src](https://github.com/IwnuplyNotTyan/Lura?tab=readme-ov-file#xbps-src)
  - [Gnumake](https://github.com/IwnuplyNotTyan/Lura?tab=readme-ov-file#gnumake)
  - [Manual](https://github.com/IwnuplyNotTyan/Lura?tab=readme-ov-file#gnumake)
- [Mods](https://github.com/IwnuplyNotTyan/Lura?tab=readme-ov-file#mods)
  - [Adding weapon & monster mod](https://github.com/IwnuplyNotTyan/Lura?tab=readme-ov-file#adding-weapon--monster-mod)
  - [Removing weapon & monster mod](https://github.com/IwnuplyNotTyan/Lura?tab=readme-ov-file#removing-weapon--monster-mod)
  - [Debug console](https://github.com/IwnuplyNotTyan/Lura?tab=readme-ov-file#removing-weapon--monster-mod)
- [Credits](https://github.com/IwnuplyNotTyan/Lura?tab=readme-ov-file#mods)

# Build
Note: build dep `go` v1.23.6. To play need only [Nerdfonts](https://www.nerdfonts.com/).

*Troubleshooting:* downgrade go version in `go.mod` to needed version

### Makepkg
```sh
makepkg -si
```

### Xbps-install
Download `lura-0.1.0-1-x86_64.xbps` from [releases](https://github.com/IwnuplyNotTyan/Lura/releases/tag/v1.0.0)
```sh
xbps-install -R ~/path/to/lura-0.1.0-1-x86_64.xbps
```

### Xbps-src
Install [xbps-src](https://github.com/void-linux/void-packages)
```sh
git clone https://github.com/IwnuplyNotTyan/Lura
mkdir -p ~/path/to/void-packages/srcpkgs/Lura
cp ~/Lura/template ./path/to/void-packages/srcpkgs/Lura/
cd ./path/to/void-packages/
./xbps-src pkg Lura
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

# Mods
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
### Debug console, add flag `-debug`, all command in help

# Credits
UA translate by [Purple Sky](https://github.com/Osian-linux) and me

Void package by [Binarnik](https://github.com/binarylinuxx/)

# Made with ❤️
