package db

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"caisse-app/app/logger"
	"caisse-app/app/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	lg "gorm.io/gorm/logger"
)

type dbProxy struct {
	db       *gorm.DB
	username string
	password string
}

var (
	instance *dbProxy
	once     sync.Once
)

func Init() {
	once.Do(func() {
		instance = &dbProxy{
			username: os.Getenv("DB_USER"),
			password: os.Getenv("DB_PASSWORD"),
		}
		instance.connect()
	})
}

func (d *dbProxy) connect() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		"5432",
		instance.username,
		instance.password,
		"postgres",
	)
	var err error
	instance.db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func GetCaissier(nom string) bool {
	var c models.Caisse
	result := instance.db.Where("nom = ?", nom).First(&c)
	if result.Error != nil {
		logger.Error("Impossible d'ouvrir la caisse: " + result.Error.Error())
	}
	return result.Error == nil && !c.Occupe
}

func SetupLog() {
	instance.db.Logger = lg.New(log.New(logger.GetFile(), "\r\n", log.LstdFlags), lg.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  lg.Warn,
		IgnoreRecordNotFoundError: false,
		Colorful:                  false,
	})
}
func OccuperCaisse(nom string) error {
	return instance.db.Model(&models.Caisse{}).Where("nom = ?", nom).Update("occupe", true).Error
}

func LibererCaisse(nom string) error {
	return instance.db.Model(&models.Caisse{}).Where("nom = ?", nom).Update("occupe", false).Error
}

func ListProduit() ([]models.Produit, error) {
	var produits []models.Produit
	err := instance.db.Find(&produits).Error
	return produits, err
}

func ListTransactions() ([]models.Transaction, error) {
	var transactions []models.Transaction
	err := instance.db.Order("date desc").Find(&transactions).Error
	return transactions, err
}

func GetProduitParID(id int) (*models.Produit, error) {
	var p models.Produit
	err := instance.db.First(&p, id).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func GetTransactionByID(transactionID int) (models.Transaction, error) {
	var t models.Transaction
	err := instance.db.First(&t, transactionID).Error
	return t, err
}

func GetProduitsParNomWildcard(nomWildcard string) ([]models.Produit, error) {
	var produits []models.Produit
	err := instance.db.Where("nom ILIKE ?", "%"+nomWildcard+"%").Find(&produits).Error
	return produits, err
}

func SetTransactionToDejaRetourne(transactionID int) error {
	err := instance.db.Model(&models.Transaction{}).Where("id = ?", transactionID).Update("deja_retourne", true).Error
	return err
}

func EnregistrerTransaction(t *models.Transaction) error {
	return instance.db.Transaction(func(tx *gorm.DB) error {
		// Insert the transaction
		if err := tx.Omit("id").Create(t).Error; err != nil {
			return err
		}

		// Update product quantities
		ids := strings.Split(t.ProduitIDs, ",")
		produitQMap := make(map[int]int)
		for _, v := range ids {
			id, _ := strconv.Atoi(v)
			produitQMap[id]++
		}

		for id, quantity := range produitQMap {
			if t.Type == "RETOUR" {
				quantity = -quantity
			}
			if err := tx.Model(&models.Produit{}).
				Where("id = ?", id).
				UpdateColumn("quantite", gorm.Expr("quantite + ?", -quantity)).Error; err != nil {
				return err
			}
		}
		return nil
	})
}
