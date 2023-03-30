package datastore

import (
	"fmt"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB
var err error
var dsn string

// rest
type Client struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Details string `json:"details"`
}

type User struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Details string `json:"details"`
}

type Card string

const (
	Basic    string = "Basic"
	Premium  string = "Premium"
	Platinum string = "Platinum"
)

type Offers struct {
	gorm.Model
	ID           uint   `gorm:"primaryKey"`
	CardName     Card   `json:"name"`
	DiscountRate uint   `json:"phone"`
	Client       Client `json:"client"`
}

type Transaction struct {
	gorm.Model
	ID     uint   `gorm:"primaryKey"`
	Offer  Offers `json:"offer"`
	Amount uint   `json:"amount"`
	UserId User   `json:"user"`
}

func init() {

	// dbLoc := fmt.Sprintf("%s/%s.db", config.Config.SERVICE_DB, config.Config.REGION)
	dbLoc := fmt.Sprintf("nimbus.db")
	Db, err = gorm.Open(sqlite.Open(dbLoc), &gorm.Config{})

	if err != nil {
		log.Panicln("failed to establish error db connection")
	}
	log.Println("db conn")

	Db.AutoMigrate(&Client{})
	Db.AutoMigrate(&User{})
	Db.AutoMigrate(&Offers{})
	Db.AutoMigrate(&Transaction{})
}
