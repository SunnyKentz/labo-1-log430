package logger

import (
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	"caisse-app/app/models"
)

type logger struct {
	file    *os.File
	posname string
}

var (
	instance *logger
	once     sync.Once
)

func Init(name string) {
	once.Do(func() {
		s := strings.ReplaceAll(name, " ", "")
		logFile, err := os.OpenFile(s+"-pos.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			println("Impossible d'initialiser les logs")
			os.Exit(1)
		}
		instance = &logger{file: logFile}
	})
	setPOSName(name)
}
func setPOSName(name string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	instance.posname = name
	fmt.Fprintf(instance.file, "[%s][POS]: %s initialisé\n", timestamp, name)
}

func Info(info string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(instance.file, "[%s][%s][INFO] %s\n", timestamp, instance.posname, info)
}

func Error(error string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(instance.file, "[%s][%s][ERROR] %s\n", timestamp, instance.posname, error)
}

func Transaction(t *models.Transaction, msg string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(instance.file, "[%s][%s][TRANSACTION] Type: %s : %s, Total: %.2f, Produit: %s\n",
		timestamp, instance.posname, t.Type, msg, t.Montant, t.ProduitIDs)
}

func GetFile() *os.File {
	return instance.file
}
