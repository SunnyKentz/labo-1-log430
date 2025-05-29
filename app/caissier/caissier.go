package caissier

import (
	"caisse-app/app/db"
	"caisse-app/app/logger"
	"caisse-app/app/models"
	. "caisse-app/app/utils"
	"fmt"
	"log"
	"sync"
	"time"
)

type caissier struct {
	Nom  string
	cart []models.Produit
}

var instance *caissier
var once sync.Once

func InitialiserPOS() bool {
	once.Do(func() {
		nom := "New Caisse"
		db.Init()
		caisses := []string{"Caisse 1", "Caisse 2", "Caisse 3"}
		for _, v := range caisses {
			available := db.GetCaissier(v)
			if !available {
				log.Print("Erreur: la Caisse n'existe pas ou est occupé :" + v)
				continue
			}
			nom = v
			logger.Init(v)
			break
		}
		if nom == "New Caisse" {
			logger.Error("Erreur aucune caisse disponible")
			return
		}
		logger.Init(nom)
		db.SetupLog()
		// Occuper la caisse
		if err := db.OccuperCaisse(nom); err != nil {
			logger.Error("Erreur lors de l'occupation de la caisse: " + err.Error())
			return
		}

		instance = &caissier{
			Nom:  nom,
			cart: make([]models.Produit, 0),
		}
	})
	return instance != nil
}
func Nom() string {
	return instance.Nom
}
func FermerPOS() {
	if instance != nil {
		err := db.LibererCaisse(instance.Nom)
		Errnotnil(err)
	}
}

func AfficherProduits() ([]models.Produit, error) {
	return db.ListProduit()
}

func AfficherTransactions() []models.Transaction {
	//recuperer les transactions
	transactions, err := db.ListTransactions()
	if err != nil {
		logger.Error("Erreur lors de la recupration des transactions: " + err.Error())
		return nil
	}
	return transactions
}

func TrouverProduit(nomPartiel string) ([]models.Produit, error) {
	return db.GetProduitsParNomWildcard(nomPartiel)
}

func AjouterALaCart(produitID int) float64 {
	produit, err := db.GetProduitParID(produitID)
	if err != nil {
		logger.Error(err.Error())
		return 0.0
	}
	instance.cart = append(instance.cart, *produit)
	total := 0.0
	for _, p := range instance.cart {
		total += p.Prix
	}
	return total
}

func TotalDeLACart() float64 {
	total := 0.0
	for _, p := range instance.cart {
		total += p.Prix
	}
	return total
}

func RetirerDeLaCart(produitID int) float64 {
	total := 0.0
	for i, p := range instance.cart {
		if p.ID == produitID {
			instance.cart = append(instance.cart[:i], instance.cart[i+1:]...)
			break
		}
	}
	for _, p := range instance.cart {
		total += p.Prix
	}
	return total
}

func ViderLaCart() {
	instance.cart = make([]models.Produit, 0)
}
func QuantiteDansLaCart(produitID int) int {
	count := 0
	for _, p := range instance.cart {
		if p.ID == produitID {
			count++
		}
	}
	return count
}

func FaireUneVente() error {
	produitIDs := ""
	total := 0.0
	for i, p := range instance.cart {
		if i > 0 {
			produitIDs += ","
		}
		total += p.Prix
		produitIDs += fmt.Sprintf("%d", p.ID)
	}

	transaction := &models.Transaction{
		Caisse:     instance.Nom,
		Type:       "VENTE",
		ProduitIDs: produitIDs,
		Montant:    total,
		Date:       time.Now(),
	}

	if err := db.EnregistrerTransaction(transaction); err != nil {
		logger.Error("Erreur lors de la vente: " + err.Error())
		return err
	}
	logger.Transaction(transaction, "Vente effectuée")
	instance.cart = make([]models.Produit, 0)
	return nil
}

func FaireUnRetour(transactionID int) error {
	transaction, err := db.GetTransactionByID(transactionID)
	if err != nil {
		logger.Error(err.Error())
		return err
	}
	if transaction.Type == "RETOUR" || transaction.Deja_retourne {
		return fmt.Errorf("cette transaction ne peut etre retourne")
	}
	transaction.Montant *= -1
	transaction.Type = "RETOUR"
	transaction.Date = time.Now()
	if err := db.EnregistrerTransaction(&transaction); err != nil {
		logger.Error("Erreur lors du retour: " + err.Error())
		return err
	}
	err = db.SetTransactionToDejaRetourne(transactionID)
	Errnotnil(err)
	logger.Transaction(&transaction, "Retour effectué")
	return nil
}
