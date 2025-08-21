# ✨ ~ **Lura**, CLI turn based roguelike

![Preview](https://github.com/IwnuplyNotTyan/Lura/blob/main/assets/preview.png?raw=true)

# Tree
- [Installing](https://github.com/IwnuplyNotTyan/Lura?tab=readme-ov-file#installing)
  - [Windows](https://github.com/iwnuplynottyan/lura?tab=readme-ov-file#windows)
  - [Arch-Based](https://github.com/IwnuplyNotTyan/Lura?tab=readme-ov-file#makepkg)
  - [Nix](https://github.com/iwnuplynottyan/lura?tab=readme-ov-file#nix)
  - [Xbps-install](https://github.com/IwnuplyNotTyan/Lura?tab=readme-ov-file#xbps-install)
  - [Xbps-src](https://github.com/IwnuplyNotTyan/Lura?tab=readme-ov-file#xbps-src)
  - [Gnumake](https://github.com/IwnuplyNotTyan/Lura?tab=readme-ov-file#gnumake)
  - [Just](https://github.com/IwnuplyNotTyan/Lura?tab=readme-ov-file#just)
  - [Manual](https://github.com/IwnuplyNotTyan/Lura?tab=readme-ov-file#gnumake)
  - [Github Action](https://github.com/IwnuplyNotTyan/Lura?tab=readme-ov-file#github-action)
- [Mods](https://github.com/IwnuplyNotTyan/Lura?tab=readme-ov-file#mods)
  - [Adding weapon & monster mod](https://github.com/IwnuplyNotTyan/Lura?tab=readme-ov-file#adding-weapon--monster-mod)
  - [Removing weapon & monster mod](https://github.com/IwnuplyNotTyan/Lura?tab=readme-ov-file#removing-weapon--monster-mod)
  - [Debug console](https://github.com/IwnuplyNotTyan/Lura?tab=readme-ov-file#removing-weapon--monster-mod)
- [Credits](https://github.com/IwnuplyNotTyan/Lura?tab=readme-ov-file#mods)

# Installing
> [!NOTE]
> Note: build dep `go` v1.23.6. To play need only [Nerdfonts](https://www.nerdfonts.com/).

> [!IMPORTANT]
> *Troubleshooting:* downgrade go version in `go.mod` to needed version

### Windows
download .exe from [release](https://github.com/IwnuplyNotTyan/Lura/releases) or build [Manual](https://github.com/IwnuplyNotTyan/Lura?tab=readme-ov-file#gnumake)

### Makepkg
```sh
makepkg -si
```

### Just
*Install*
```sh
just install
```
*Build*
```sh
just build
```
*Run*
```sh
just run
```

### Nix

> [!WARNING]
> UNMAINTAINED

> [!IMPORTANT]
> *Troubleshooting:* update sha256 in `flake.nix` to needed version

```sh
nix build
```

### Xbps-install
From unofficial repo

Write in `/etc/xbps.d/void-extras.conf` this:
```
repository=https://raw.githubusercontent.com/binarylinuxx/void-extras/x86_64/pkgs
```

Or manualy

Download `lura-*_1-x86_64.xbps` from [releases](https://github.com/IwnuplyNotTyan/Lura/releases/tag/v1.0.0)
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
go install
```
or
```sh
go build -o lura
```

### Github Action
Login and check install from [Action](https://github.com/IwnuplyNotTyan/Lura/actions/workflows/build.yml) button

# Mods
> [!NOTE]
> Use `nix develop` to run in dev env, or `nix-shell` to enter dev shell

### Mod folder in `~/.local/share/Lura/mods/`, all mods must be in lua

<details>
  <summary>Adding weapon & monster mod</summary>

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

</details>


<details>
  <summary>Removing weapon & monster mod</summary>

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

</details>

### Debug console, add flag `-debug`, all command in help

# Credits
UA translate by [Purple Sky](https://github.com/Osian-linux) and me

BE translate by [Табунов Артем](https://t.me/perakladadkata)

Void package by [Binarnik](https://github.com/binarylinuxx/)

# Made with ❤️
