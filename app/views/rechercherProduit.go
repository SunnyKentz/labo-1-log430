package views

import (
	cs "caisse-app/app/caissier"
	"caisse-app/app/models"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type RechercherProduit struct {
	textInput textinput.Model
	produits  []models.Produit
	err       error
}

func NewRechercherProduit() RechercherProduit {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 30
	produits, err := cs.AfficherProduits()
	return RechercherProduit{
		textInput: ti,
		err:       err,
		produits:  produits,
	}
}

func (r RechercherProduit) Init() tea.Cmd {
	if r.err != nil {
		fmt.Println(r.err)
		return tea.Quit
	}
	return textinput.Blink
}

func (r RechercherProduit) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			return r, tea.Quit
		default:
			if len(r.produits) == 1 {
				return r, nil
			}
			r.textInput, cmd = r.textInput.Update(msg)
			r.produits, _ = cs.TrouverProduit(r.textInput.Value())
		}
	}
	return r, cmd
}

func (r RechercherProduit) View() string {
	if len(r.produits) == 1 {
		return fmt.Sprintf("Quantite : %d \nnom: %s \nCatégorie: %s \nPrix: %.2f$ \n\n(Press esc to quit)",
			r.produits[0].Quantite, r.produits[0].Nom, r.produits[0].Categorie, r.produits[0].Prix)

	}
	s := strings.Builder{}
	s.WriteString("\nProduits trouvé:\n\n")
	for _, prod := range r.produits {
		q := fmt.Sprintf("  [ %v ]", prod.Quantite)
		s.WriteString(fmt.Sprintf("%s %s\n", q, prod.Nom))
	}
	s.WriteString("\n")
	s.WriteString(" Recherche : ")
	s.WriteString(r.textInput.View())
	s.WriteString("\nPress esc to quit")
	return s.String()
}
