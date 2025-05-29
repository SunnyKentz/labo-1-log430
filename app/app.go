package app

import (
	"caisse-app/app/caissier"
	"caisse-app/app/views"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func App() {
	if !caissier.InitialiserPOS() {
		os.Exit(1)
	}
	p := tea.NewProgram(views.NewMainMenu())
	if _, err := p.Run(); err != nil {
		println(err)
	}
	caissier.FermerPOS()
}
