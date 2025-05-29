package views

import (
	"caisse-app/app/caissier"
	. "caisse-app/app/utils"
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type MainMenu struct {
	choices  []string
	selected map[int]struct{}
	cursor   int
	err      error
}

func NewMainMenu() MainMenu {
	return MainMenu{
		choices: []string{
			"Rechercher un produit",
			"Vente de produit",
			"Afficher les transactions",
			"Quitter",
		},
		selected: make(map[int]struct{}),
		cursor:   0,
		err:      nil,
	}
}

func (m MainMenu) Init() tea.Cmd {
	m.selected[m.cursor] = struct{}{}
	return tea.ClearScreen
}

func (m MainMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "down":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			} else if m.cursor > len(m.choices)-1 {
				m.cursor = len(m.choices) - 1
			}
			clear(m.selected)
			m.selected[m.cursor] = struct{}{}
		case "up":
			if m.cursor > 0 {
				m.cursor--
			} else if m.cursor < 0 {
				m.cursor = 0
			}
			clear(m.selected)
			m.selected[m.cursor] = struct{}{}
		case "enter":
			return m, switchToWindow(m.cursor)
		case "ctrl+c", "esc":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m MainMenu) View() string {
	if m.err != nil {
		return m.err.Error()
	}
	s := strings.Builder{}
	s.WriteString("\nMenu principal " + caissier.Nom() + "\n\n")
	for i, choice := range m.choices {
		checked := "[    ]"
		if _, ok := m.selected[i]; ok {
			checked = "[ -> ]"
		}
		s.WriteString(fmt.Sprintf("%s %s\n", checked, choice))
	}
	s.WriteString("\n")
	s.WriteString("Press esc to quit")
	return s.String()
}

func switchToWindow(index int) tea.Cmd {
	switch index {
	case 0: // Rechercher un produit
		p := tea.NewProgram(NewRechercherProduit())
		_, err := p.Run() // blocking
		Errnotnil(err)
		return tea.ClearScreen
	case 1: // Vente de produit
		p := tea.NewProgram(NewVenteProduit())
		_, err := p.Run() // blocking
		Errnotnil(err)
		return tea.ClearScreen
	case 2: // Afficher les transactions
		p := tea.NewProgram(NewTransactionsModel())
		_, err := p.Run() // blocking
		Errnotnil(err)
		return tea.ClearScreen
	case 3: // Quitter
		return tea.Quit
	}
	return tea.Quit
}
