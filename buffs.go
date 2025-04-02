package main

import (
	"fmt"
	"math/rand"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
)

var (
	buff1 string
	buff2 string
	buff3 string
)

var (
	red    = lipgloss.AdaptiveColor{Light: "#FE5F86", Dark: "#FE5F86"}
	indigo = lipgloss.AdaptiveColor{Light: "#5A56E0", Dark: "#7571F9"}
	green  = lipgloss.AdaptiveColor{Light: "#02BA84", Dark: "#02BF87"}
)

const maxWidth = 80

type Styles struct {
	Base,
	HeaderText,
	Status,
	StatusHeader,
	Highlight,
	ErrorHeaderText,
	Help lipgloss.Style
}

func NewStyles(lg *lipgloss.Renderer) *Styles {
	s := Styles{}
	s.Base = lg.NewStyle().
		Padding(1, 4, 0, 1)
	s.HeaderText = lg.NewStyle().
		Foreground(indigo).
		Bold(true).
		Padding(0, 1, 0, 2)
	s.Status = lg.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(indigo).
		PaddingLeft(1).
		MarginTop(1)
	s.StatusHeader = lg.NewStyle().
		Foreground(green).
		Bold(true)
	s.Highlight = lg.NewStyle().
		Foreground(lipgloss.Color("212"))
	s.ErrorHeaderText = s.HeaderText.
		Foreground(red)
	s.Help = lg.NewStyle().
		Foreground(lipgloss.Color("240"))
	return &s
}

type state int

const (
	statusNormal state = iota
	stateDone
)

type Model struct {
	state  state
	lg     *lipgloss.Renderer
	styles *Styles
	form   *huh.Form
	width  int
}

func NewModel(player *Player) Model {
	m := Model{width: maxWidth}
	m.lg = lipgloss.DefaultRenderer()
	m.styles = NewStyles(m.lg)

	buff1 = getRandomBuff(player)
	buff2 = getRandomBuff(player)
	buff3 = getRandomBuff(player)

	m.form = huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Key("class").
				Options(huh.NewOptions(buff1, buff2, buff3)...).
				Title("Choose your class").
				Description("This will determine your department"),
		),
	).
		WithWidth(45).
		WithShowHelp(false).
		WithShowErrors(false)
	return m
}

func (m Model) Init() tea.Cmd {
	return m.form.Init()
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = min(msg.Width, maxWidth) - m.styles.Base.GetHorizontalFrameSize()
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Interrupt
		case "esc", "q":
			return m, tea.Quit
		}
	}

	var cmds []tea.Cmd

	// Process the form
	form, cmd := m.form.Update(msg)
	if f, ok := form.(*huh.Form); ok {
		m.form = f
		cmds = append(cmds, cmd)
	}

	if m.form.State == huh.StateCompleted {
		// Quit when the form is done.
		cmds = append(cmds, tea.Quit)
	}

	return m, tea.Batch(cmds...)
}

func (m Model) View() string {
	s := m.styles

	switch m.form.State {
	case huh.StateCompleted:
		title, role := m.getRole()
		title = s.Highlight.Render(title)
		var b strings.Builder
		fmt.Fprintf(&b, "Congratulations, you’re Charm’s newest\n%s!\n\n", title)
		fmt.Fprintf(&b, "Your job description is as follows:\n\n%s\n\nPlease proceed to HR immediately.", role)
		return s.Status.Margin(0, 1).Padding(1, 2).Width(48).Render(b.String()) + "\n\n"
	default:

		var class string
		if m.form.GetString("class") != "" {
			class = "Class: " + m.form.GetString("class")
		}

		// Form (left side)
		v := strings.TrimSuffix(m.form.View(), "\n\n")
		form := m.lg.NewStyle().Margin(1, 0).Render(v)

		// Status (right side)
		var status string
		{
			var (
				buildInfo      = "(None)"
				role           string
				jobDescription string
				level          string
			)

			if m.form.GetString("level") != "" {
				level = "Level: " + m.form.GetString("level")
				role, jobDescription = m.getRole()
				role = "\n\n" + s.StatusHeader.Render("Projected Role") + "\n" + role
				jobDescription = "\n\n" + s.StatusHeader.Render("Duties") + "\n" + jobDescription
			}
			if m.form.GetString("class") != "" {
				buildInfo = fmt.Sprintf("%s\n%s", class, level)
			}

			const statusWidth = 28
			statusMarginLeft := m.width - statusWidth - lipgloss.Width(form) - s.Status.GetMarginRight()
			status = s.Status.
				Height(lipgloss.Height(form)).
				Width(statusWidth).
				MarginLeft(statusMarginLeft).
				Render(s.StatusHeader.Render("Current Build") + "\n" +
					buildInfo +
					role +
					jobDescription)
		}

		errors := m.form.Errors()
		header := m.appBoundaryView("Charm Employment Application")
		if len(errors) > 0 {
			header = m.appErrorBoundaryView(m.errorView())
		}
		body := lipgloss.JoinHorizontal(lipgloss.Left, form, status)

		footer := m.appBoundaryView(m.form.Help().ShortHelpView(m.form.KeyBinds()))
		if len(errors) > 0 {
			footer = m.appErrorBoundaryView("")
		}

		return s.Base.Render(header + "\n" + body + "\n\n" + footer)
	}
}

func (m Model) errorView() string {
	var s string
	for _, err := range m.form.Errors() {
		s += err.Error()
	}
	return s
}

func (m Model) appBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.HeaderText.Render(text),
		lipgloss.WithWhitespaceForeground(indigo),
	)
}

func (m Model) appErrorBoundaryView(text string) string {
	return lipgloss.PlaceHorizontal(
		m.width,
		lipgloss.Left,
		m.styles.ErrorHeaderText.Render(text),
		lipgloss.WithWhitespaceForeground(red),
	)
}

func (m Model) getRole() (string, string) {
	switch m.form.GetString("class") {
	case "Warrior":
		return
	case "Mage":
		return
	case "Rogue":
		return
	default:
		return "", ""
	}
}

func getRandomBuff(player *Player, excludeBuffs ...string) string {
	var buffs []string
	if player.loc == 1 {
		if lang == "en" {
			buffs = []string{
				"Upgrade Weapon",
				"Longsword",
				"Crossbow",
				//"Random Weapon",
				"Broken heart",
				"Turtle scute",
			}
		} else if lang == "be" {
			buffs = []string{
				"Палепшыць зброю",
				//"Выпадковая зброя",
				"Разбітае сэрца",
				"Шчыт чарапахі",
			}
		} else {
			buffs = []string{
				"Покращити зброю",
				//"Випадкова зброя",
				"Розбите серце",
				"Щиток черепахи",
			}
		}
	} else if player.loc == 0 {
		if lang == "en" {
			buffs = []string{
				"Crystal heart",
				"Lotus",
				"Tears",
				//"Amethyst necklace",
				//"Flask with star tears",
			}
		} else if lang == "be" {
			buffs = []string{
				"Кристалічна сэрца",
				"Лотас",
				"Слёзы",
			}
		} else if lang == "ua" {
			buffs = []string{
				"Кристалічне серце",
				"Лотос",
				"Сльози",
			}
		}
	}

	availableBuffs := make([]string, 0, len(buffs))
	for _, buff := range buffs {
		if !contains(excludeBuffs, buff) {
			availableBuffs = append(availableBuffs, buff)
		}
	}

	if len(availableBuffs) == 0 {
		return ""
	}

	return availableBuffs[rand.Intn(len(availableBuffs))]
}

func buffsAction(player *Player) {
	currentCoins(player)

	buff1 = getRandomBuff(player)
	buff2 = getRandomBuff(player)
	buff3 = getRandomBuff(player)

	selectedBuffs := selectBuff(player)
	if len(selectedBuffs) == 0 {
		noBuffDialog()
		return
	}

	for _, buff := range selectedBuffs {
		switch buff {

		//case "Amethyst necklace":
		//	if player.Coins > 20 {
		//		player.amenuck = true
		//	} else {
		//		noBuffDialog()
		//	}

		//case "Flask with star tears":
		//	if player.Coins > 100 {
		//		maxh := player.maxHP
		//		d := player.Damage
		//		player.Damage *= 2
		//		player.maxHP += 2
		//		player.monster = true
		//		fmt.Println(termenv.String(fmt.Sprintf("  %d  %d", maxh, player.maxHP)).Foreground(termenv.ANSIGreen))
		//		fmt.Println(termenv.String(fmt.Sprintf("  %d  %d", d, player.Damage)).Foreground(termenv.ANSIGreen))
		//	}

		case "Longsword":
			if player.Coins > 20 {
				w := player.WeaponType
				getLongsword(player)
				fmt.Println(termenv.String(fmt.Sprintf("  %s  %s", w, player.WeaponType)).Foreground(termenv.ANSIGreen))
			} else {
				noBuffDialog()
			}

		case "Crossbow":
			if player.Coins > 20 {
				w := player.WeaponType
				getCrossbow(player)
				fmt.Println(termenv.String(fmt.Sprintf("  %s  %s", w, player.WeaponType)).Foreground(termenv.ANSIGreen))
			} else {
				noBuffDialog()
			}

		case "Щиток черепахи", "Turtle scute", "Шчыт чарапахі":
			if player.Coins > 20 {
				player.Coins -= 20
				currentHp := player.HP
				player.HP += 50
				fmt.Println(termenv.String(fmt.Sprintf("  %d  %d", currentHp, player.HP)).Foreground(termenv.ANSIGreen))
			} else {
				noBuffDialog()
			}

		case "Crystal heart":
			if player.Coins > 50 {
				player.Coins -= 50
				player.heart = 2
				fmt.Println(termenv.String(fmt.Sprintf("  Your heart regenerate new power")))
			} else {
				noBuffDialog()
			}

		case "Lotus", "Лотус", "Лотас":
			if player.Coins > 10 {
				player.Coins -= 10
				currentMaxStamina := player.maxStamina
				player.maxStamina += 10
				fmt.Println(termenv.String(fmt.Sprintf("  %d  %d", currentMaxStamina, player.maxStamina)).Foreground(termenv.ANSIGreen))
			} else {
				noBuffDialog()
			}

		case "Tears", "Сльози", "Слёзы":
			if player.Coins > 5 {
				player.Coins -= 5
				currentMaxHP := player.maxHP
				player.maxHP += 10
				fmt.Println(termenv.String(fmt.Sprintf("  %d  %d", currentMaxHP, player.maxHP)).Foreground(termenv.ANSIGreen))
			} else {
				noBuffDialog()
			}

		case "Broken heart", "Розбите серце", "Разбітае сэрца":
			if player.Coins > 50 {
				player.heart = 0
				player.Coins -= 50
				fmt.Println(termenv.String(fmt.Sprintf("  No heart now.")))
			} else {
				noBuffDialog()
			}

		case "Upgrade Weapon", "Покращити зброю", "Палепшыць зброю":
			if player.Coins > 30 {
				player.Coins -= 30
				CurrentDamage := player.Damage
				player.Damage += 10
				fmt.Println(termenv.String(fmt.Sprintf("  %d  %d", CurrentDamage, player.Damage)).Foreground(termenv.ANSIGreen))
			} else {
				noBuffDialog()
			}

		default:
			noBuffDialog()
		}
	}
}

func selectBuff(player *Player) []string {
	var selectedBuffs []string

	buff1 := getRandomBuff(player)
	buff2 := getRandomBuff(player, buff1)
	buff3 := getRandomBuff(player, buff1, buff2)

	f := huh.NewForm(
		huh.NewGroup(
			huh.NewMultiSelect[string]().
				Title(" Select card").
				Options(
					huh.NewOption(buff1, buff1),
					huh.NewOption(buff2, buff2),
					huh.NewOption(buff3, buff3),
				).
				Value(&selectedBuffs),
		),
	)

	if err := f.Run(); err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return selectedBuffs
}

func contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
