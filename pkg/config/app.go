package config

import(
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

var (
	db * gorm.DB
)

func Connect() {
	databaseUSER := os.Getenv("DBUSERINFO")
	d, err := gorm.Open("mysql", databaseUSER)

	if err  != nil {
		panic(err)
	}

	db = d
}

func GetDB() *gorm.DB{
	return db
}