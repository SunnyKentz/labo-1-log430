package tests

import (
	"caisse-app/app/caissier"
	"os"
	"testing"
	"time"
)

func BeforeAll() {
	os.Setenv("DB_PASSWORD", "test_password")
	os.Setenv("DB_USER", "test_user")
	caissier.InitialiserPOS()
	time.Sleep(time.Second * 5)
}

func TestMain(m *testing.M) {
	BeforeAll()
	code := m.Run()
	caissier.FermerPOS()
	os.Exit(code)
}

func TestInitialiserPOS(t *testing.T) {
	if !caissier.InitialiserPOS() {
		t.Error("InitialiserPOS should return true")
	}
	if caissier.Nom() == "" {
		t.Error("Nom should not be empty after initialization")
	}
}

func TestNom(t *testing.T) {
	nom := caissier.Nom()
	if nom == "" {
		t.Error("Nom should not be empty")
	}
}

func TestAfficherProduits(t *testing.T) {
	produits, err := caissier.AfficherProduits()
	errnotnilFatal("there was an error", t, err)
	if len(produits) == 0 {
		t.Error("Expected at least one produit")
	}
}

func TestAfficherTransactions(t *testing.T) {
	transactions := caissier.AfficherTransactions()
	if transactions == nil {
		t.Error("Transactions should not be nil")
	}
}

func TestTrouverProduit(t *testing.T) {
	produits, err := caissier.TrouverProduit("Eau")
	errnotnilFatal("Error finding product", t, err)
	if len(produits) == 0 {
		t.Error("Expected to find at least one product containing 'Eau'")
	}
}

func TestAjouterALaCart(t *testing.T) {
	// First get a product to add
	produits, err := caissier.AfficherProduits()
	errnotnilFatal("Error getting products", t, err)
	if len(produits) == 0 {
		t.Fatal("No products available for test")
	}

	total := caissier.AjouterALaCart(produits[0].ID)
	if total <= 0 {
		t.Error("Total should be greater than 0 after adding a product")
	}
}

func TestTotalDeLACart(t *testing.T) {
	total := caissier.TotalDeLACart()
	if total < 0 {
		t.Error("Total should not be negative")
	}
}

func TestRetirerDeLaCart(t *testing.T) {
	// First get a product to add
	produits, err := caissier.AfficherProduits()
	errnotnilFatal("Error getting products", t, err)
	if len(produits) == 0 {
		t.Fatal("No products available for test")
	}

	// Add a product first
	caissier.AjouterALaCart(produits[0].ID)
	initialTotal := caissier.TotalDeLACart()

	// Then remove it
	newTotal := caissier.RetirerDeLaCart(produits[0].ID)
	if newTotal >= initialTotal {
		t.Error("Total should decrease after removing a product")
	}
}

func TestQuantiteDansLaCart(t *testing.T) {
	// First get a product to add
	produits, err := caissier.AfficherProduits()
	errnotnilFatal("Error getting products", t, err)
	if len(produits) == 0 {
		t.Fatal("No products available for test")
	}
	caissier.ViderLaCart()
	// Add the same product twice
	caissier.AjouterALaCart(produits[0].ID)
	caissier.AjouterALaCart(produits[0].ID)

	quantity := caissier.QuantiteDansLaCart(produits[0].ID)
	if quantity != 2 {
		t.Errorf("Expected quantity to be 2, got %d", quantity)
	}
}

func TestFaireUneVente(t *testing.T) {
	// First get a product to add
	produits, err := caissier.AfficherProduits()
	errnotnilFatal("Error getting products", t, err)
	if len(produits) == 0 {
		t.Fatal("No products available for test")
	}

	// Add a product to cart
	caissier.AjouterALaCart(produits[0].ID)

	// Make the sale
	err = caissier.FaireUneVente()
	errnotnilFatal("Error making sale", t, err)

	// Verify cart is empty
	if caissier.TotalDeLACart() != 0 {
		t.Error("Cart should be empty after sale")
	}
}

func TestFaireUnRetour(t *testing.T) {
	// First make a sale
	produits, err := caissier.AfficherProduits()
	errnotnilFatal("Error getting products", t, err)
	if len(produits) == 0 {
		t.Fatal("No products available for test")
	}

	// Add a product and make a sale
	caissier.AjouterALaCart(produits[0].ID)
	err = caissier.FaireUneVente()
	errnotnilFatal("Error making sale", t, err)

	// Get the transaction
	transactions := caissier.AfficherTransactions()
	if len(transactions) == 0 {
		t.Fatal("No transactions available for test")
	}

	// Make the return
	err = caissier.FaireUnRetour(transactions[0].ID)
	errnotnilFatal("Error making return", t, err)

	// Try to return the same transaction again (should fail)
	err = caissier.FaireUnRetour(transactions[0].ID)
	if err == nil {
		t.Error("Expected error when trying to return the same transaction twice")
	}
}

func errnotnilFatal(s string, t *testing.T, err error) {
	if err != nil {
		t.Fatal(s + ": " + err.Error())
	}
}
