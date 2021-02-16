package database

import (
	"fmt"
	"os"
	"path"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	homedir "github.com/mitchellh/go-homedir"
)

const app = "crypto"

var dbPath string

type Portfolio struct {
	gorm.Model
	Name         string `gorm:"unique;not null"`
	Description  string
	Transactions []Transaction
}

type Transaction struct {
	gorm.Model
	PortfolioID int
	// BTC
	CoinName string `gorm:"not null"`
	// 0.025896
	CoinAmount float64 `gorm:"not null"`
	// 2.77
	CoinPrice float64 `gorm:"not null"`
	// 3.84
	Fee float64 `gorm:"not null"`
	// 100.00
	Total float64 `gorm:"not null"`
}

func init() {
	GetDb().AutoMigrate(&Portfolio{}, &Transaction{})
}

func GetDb() *gorm.DB {
	db, err := gorm.Open("sqlite3", getDbPath())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return db
}

func getDbPath() string {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	conf := fmt.Sprintf(".%v.db", app)
	return path.Join(home, conf)
}
