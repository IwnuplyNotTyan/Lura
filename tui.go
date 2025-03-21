package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205")) // Pink title
	cursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("212")).Bold(true) // Selected option
	choiceStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("245"))            // Unselected option (gray)
	quitStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("240")).Italic(true)
)

type model struct {
	cursor   int
	choices  []string
	selected string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter":
			m.selected = m.choices[m.cursor]
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	s := titleStyle.Render("󱢡  Select a card") + "\n\n"

	for i, choice := range m.choices {
		if i == m.cursor {
			s += cursorStyle.Render("  "+choice) + "\n"
		} else {
			s += choiceStyle.Render("  "+choice) + "\n"
		}
	}
	if lang == "ua" {
		s += "\n" + quitStyle.Render("(Використовуйте ↑/↓ для навігації)") + "\n"
	} else {
		s += "\n" + quitStyle.Render("(Use ↑/↓ to navigate)") + "\n"
	}

	return s
}

func getSelectedLanguage() string {
	p := tea.NewProgram(languageModel())

	m, err := p.Run()
	if err != nil {
		log.Fatalf("Error running language selection: %v", err)
	}

	selectedModel, ok := m.(model)
	if !ok {
		log.Fatalf("Unexpected model type")
	}
	clearScreen()
	switch selectedModel.selected {
	case "Українська":
		return "ua"
	default:
		return "en"
	}
}

func getSelectedAttack() string {
	p := tea.NewProgram(attackModel())

	m, err := p.Run()
	if err != nil {
		log.Fatalf("Error running language selection: %v", err)
	}

	selectedModel, ok := m.(model)
	if !ok {
		log.Fatalf("Unexpected model type")
	}
	clearScreen()
	switch selectedModel.selected {
	case "Attack", "Атакувати":
		return "Attack"
	case "Defend", "Захищатися":
		return "Defend"
	case "Heal", "Лікуватися":
		return "Heal"
	case "Skip", "Пропустити":
		return "Skip"
	default:
		return "Attack"
	}
}

func getSelectedBuff() string {
	p := tea.NewProgram(buffsModel())

	m, err := p.Run()
	if err != nil {
		log.Fatalf("Error running language selection: %v", err)
	}

	selectedModel, ok := m.(model)
	if !ok {
		log.Fatalf("Unexpected model type")
	}
	clearScreen()
	switch selectedModel.selected {
	case "buff1":
		return buff1
	case "buff2":
		return buff2
	case "buff3":
		return buff3
	default:
		return buff1
	}
}
