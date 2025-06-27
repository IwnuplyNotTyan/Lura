package main

import (
	"bytes"
	"embed"
	"fmt"
	"image/png"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/qeesung/image2ascii/convert"
)

//go:embed assets/*
var assets embed.FS

type Buff struct {
	Name     string
	Desc     string
	ArtFile  string // Path to the art file in assets
}

type model struct {
	buffs    []Buff
	cursor   int
	selected map[int]bool
	done     bool
}

func loadArtFromFile(filename string) []string {
	// Check if it's a PNG file
	if strings.HasSuffix(filename, ".png") {
		return loadPNGAsASCII(filename)
	}
	
	// Load text file as before
	content, err := assets.ReadFile(filename)
	if err != nil {
		// Return a default art if file not found
		return []string{
			"  ERROR   ",
			" LOADING  ",
			"   ART    ",
			"         ",
			"  FILE    ",
			"   NOT    ",
			"  FOUND   ",
		}
	}
	
	lines := strings.Split(string(content), "\n")
	
	// Ensure we have exactly 7 lines for consistent display
	art := make([]string, 7)
	for i := 0; i < 7; i++ {
		if i < len(lines) {
			// Trim and pad/truncate to 10 characters for consistent width
			line := strings.TrimRight(lines[i], "\r\n")
			if len(line) > 10 {
				art[i] = line[:10]
			} else {
				art[i] = fmt.Sprintf("%-10s", line)
			}
		} else {
			art[i] = "          " // 10 spaces
		}
	}
	
	return art
}

func loadPNGAsASCII(filename string) []string {
	// Read PNG file from embedded assets
	data, err := assets.ReadFile(filename)
	if err != nil {
		return []string{
			"  ERROR   ",
			" LOADING  ",
			"   PNG    ",
			"  " + err.Error()[:6] + "  ",
			"  FILE    ",
			"   NOT    ",
			"  FOUND   ",
		}
	}

	// Decode PNG using bytes.Reader instead of strings.Reader
	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return []string{
			"  ERROR   ",
			"DECODING  ",
			"   PNG    ",
			"  " + err.Error()[:6] + "  ",
			"  FILE    ",
			" CORRUPT  ",
			"         ",
		}
	}

	// Convert to ASCII using the correct method for image.Image
	converter := convert.NewImageConverter()
	options := convert.DefaultOptions
	options.FixedWidth = 10
	options.FixedHeight = 7
	options.Colored = false // Set to true if you want colored ASCII
	
	// Use Image2ASCIIString instead of ImageFile2ASCIIString
	asciiArt := converter.Image2ASCIIString(img, &options)
	
	// Debug: if asciiArt is empty, show debug info
	if asciiArt == "" {
		return []string{
			"  DEBUG   ",
			" EMPTY    ",
			"  ASCII   ",
			"  ART     ",
			fmt.Sprintf("W:%d H:%d", img.Bounds().Dx(), img.Bounds().Dy())[:10],
			"         ",
			"         ",
		}
	}
	
	// Split into lines and ensure 7 lines
	lines := strings.Split(asciiArt, "\n")
	art := make([]string, 7)
	for i := 0; i < 7; i++ {
		if i < len(lines) {
			line := lines[i]
			if len(line) > 10 {
				art[i] = line[:10]
			} else {
				art[i] = fmt.Sprintf("%-10s", line)
			}
		} else {
			art[i] = "          "
		}
	}
	
	return art
}

func initialModel() model {
	buffs := []Buff{
		{
			Name:    "Longsword",
			Desc:    "A sharp blade for close combat",
			ArtFile: "assets/buffs/longsword.txt",
		},
		{
			Name:    "Shield",
			Desc:    "Defensive equipment for protection",
			ArtFile: "assets/buffs/shield.txt",
		},
		{
			Name:    "Bow",
			Desc:    "Ranged weapon for distant targets",
			ArtFile: "assets/buffs/bow.txt",
		},
		{
			Name:    "Potion",
			Desc:    "Magical elixir for healing",
			ArtFile: "assets/buffs/potion.txt",
		},
	}

	return model{
		buffs:    buffs,
		selected: make(map[int]bool),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.buffs)-1 {
				m.cursor++
			}

		case " ":
			if m.selected[m.cursor] {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = true
			}

		case "enter":
			m.done = true
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	var s strings.Builder

	// Header
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		MarginBottom(1)
	
	s.WriteString(headerStyle.Render("Select buffs (space to select, enter to continue)"))
	s.WriteString("\n\n")

	// Buffs
	for i, buff := range m.buffs {
		// Create styles
		nameStyle := lipgloss.NewStyle().Bold(true)
		descStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			MarginLeft(3)

		// Cursor indicator
		cursor := " "
		if m.cursor == i {
			cursor = ">"
			nameStyle = nameStyle.Foreground(lipgloss.Color("212"))
			descStyle = descStyle.Foreground(lipgloss.Color("212"))
		}

		// Selection indicator
		checked := " "
		if m.selected[i] {
			checked = "✓"
		}

		// Format buff name and description
		buffName := fmt.Sprintf("%s [%s] %s", cursor, checked, buff.Name)
		buffDesc := fmt.Sprintf("   %s", buff.Desc)

		// Create left side (name + desc)
		leftSide := nameStyle.Render(buffName) + "\n" + descStyle.Render(buffDesc)

		// Only show ASCII art for focused buff
		if m.cursor == i {
			artStyle := lipgloss.NewStyle().
				Foreground(lipgloss.Color("33")).
				MarginLeft(2)

			// Load and create right side (ASCII art)
			artLines := loadArtFromFile(buff.ArtFile)
			var styledArtLines []string
			for _, line := range artLines {
				styledArtLines = append(styledArtLines, artStyle.Render(line))
			}
			rightSide := strings.Join(styledArtLines, "\n")

			// Combine left and right sides
			leftStyle := lipgloss.NewStyle().Width(40)
			rightStyle := lipgloss.NewStyle().Width(15)

			row := lipgloss.JoinHorizontal(
				lipgloss.Top,
				leftStyle.Render(leftSide),
				rightStyle.Render(rightSide),
			)

			s.WriteString(row)
		} else {
			// Just show the buff without art
			s.WriteString(leftSide)
		}

		s.WriteString("\n\n")
	}

	// Footer
	footerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")).
		MarginTop(1)

	footer := "Use ↑↓/jk to move, space to select, enter to continue, q to quit"
	s.WriteString(footerStyle.Render(footer))

	return s.String()
}
