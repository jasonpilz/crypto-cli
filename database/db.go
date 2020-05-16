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

var (
	dbPath string
	Sqlite *gorm.DB
)

func init() {
	sqlite, err := gorm.Open("sqlite3", getDbPath())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer sqlite.Close()
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
