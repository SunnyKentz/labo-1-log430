package views

import (
	"fmt"
	"strconv"

	cs "caisse-app/app/caissier"
	"caisse-app/app/models"

	"github.com/charmbracelet/bubbles/table"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color("240"))

type TransactionsModel struct {
	table        table.Model
	textInput    textinput.Model
	transactions []models.Transaction
	err          string
}

func NewTransactionsModel() TransactionsModel {
	columns := []table.Column{
		{Title: "ID", Width: 4},
		{Title: "Date", Width: 20},
		{Title: "Type", Width: 10},
		{Title: "Produits ID", Width: 20},
		{Title: "Montant", Width: 10},
	}
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 30
	transactions := cs.AfficherTransactions()
	var rows []table.Row
	for _, t := range transactions {
		rows = append(rows, table.Row{
			fmt.Sprintf("%d", t.ID),
			t.Date.Format("2006-01-02 15:04:05"),
			t.Type,
			t.ProduitIDs,
			fmt.Sprintf("%.2f$", t.Montant),
		})
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows(rows),
		table.WithFocused(true),
		table.WithHeight(10),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)
	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return TransactionsModel{
		textInput:    ti,
		table:        t,
		transactions: transactions,
	}
}

func (m TransactionsModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m TransactionsModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			return m, tea.Quit
		case "up":
			m.table.MoveUp(1)
		case "down":
			m.table.MoveDown(1)
		case "enter":
			m.err = ""
			if v, err := strconv.Atoi(m.textInput.Value()); err == nil && v > 0 {
				if err := cs.FaireUnRetour(v); err == nil {
					return m, tea.Quit
				}
				m.err = "Impossible D'efectué ce retour"
			}
		}
		m.textInput, _ = m.textInput.Update(msg)
		m.table, cmd = m.table.Update(msg)
	}

	return m, cmd
}

func (m TransactionsModel) View() string {
	return baseStyle.Render(m.table.View()) + "\n\n" +
		"  Rembouser transaction par ID " + m.textInput.View() +
		"\n" + m.err +
		"\n↑/↓: Naviguer • Échap: Retour"
}
