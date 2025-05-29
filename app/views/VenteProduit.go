package views

import (
	cs "caisse-app/app/caissier"
	"caisse-app/app/models"
	. "caisse-app/app/utils"
	"errors"
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type VenteProduit struct {
	textInput textinput.Model
	produits  []models.Produit
	total     float64
	err       error
}

func NewVenteProduit() VenteProduit {
	ti := textinput.New()
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 30
	produits, err := cs.AfficherProduits()
	return VenteProduit{
		textInput: ti,
		err:       err,
		produits:  produits,
		total:     cs.TotalDeLACart(),
	}
}

func (r VenteProduit) Init() tea.Cmd {
	if r.err != nil {
		fmt.Println(r.err)
		return tea.Quit
	}
	return textinput.Blink
}

func (r VenteProduit) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			return r, tea.Quit
		case "enter":
			if r.total == 0 {
				r.err = fmt.Errorf("ajouter un produit")
				return r, nil
			}
			r.err = nil
			err := cs.FaireUneVente()
			Errnotnil(err)
			return r, tea.Quit
		case "+":
			if len(r.produits) != 1 {
				r.textInput, cmd = r.textInput.Update(msg)
				return r, cmd
			}

			if r.produits[0].Quantite < cs.QuantiteDansLaCart(r.produits[0].ID) {
				r.err = errors.New("quantité insuffisant de produit")
				return r, nil
			}
			r.err = nil
			r.total = cs.AjouterALaCart(r.produits[0].ID)
			return r, nil
		case "-":
			if len(r.produits) != 1 {
				r.textInput, cmd = r.textInput.Update(msg)
				return r, cmd
			}
			r.err = nil
			r.total = cs.RetirerDeLaCart(r.produits[0].ID)
			return r, nil
		default:
			r.textInput, cmd = r.textInput.Update(msg)
			r.produits, _ = cs.TrouverProduit(r.textInput.Value())
		}
	}
	return r, cmd
}

func (r VenteProduit) View() string {

	s := strings.Builder{}
	s.WriteString("\nProduits Disponible:\n\n")
	for _, prod := range r.produits {
		q := fmt.Sprintf("  [ %v ]", prod.Quantite)
		s.WriteString(fmt.Sprintf("%s %s %s %.2f$\n", q, prod.Nom, prod.Categorie, prod.Prix))
	}
	s.WriteString("\n")
	s.WriteString(" Recherche : ")
	s.WriteString(r.textInput.View())
	if r.err != nil {
		s.WriteString("\nError : " + r.err.Error())
	}
	s.WriteString(fmt.Sprintf("\n\n Total : %.2f$", r.total))
	s.WriteString("\nPress + to add")
	s.WriteString("\nPress - to remove")
	s.WriteString("\nPress Enter to Checkout")
	s.WriteString("\nPress esc to quit")
	return s.String()
}
