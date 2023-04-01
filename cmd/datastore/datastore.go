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

type KeyPair struct {
	Nonce string `json:"nonce"`
	Key   string `json:"key"  gorm:"index:idx_key,unique"`
	Value string `json:"value"`
}

// type Country struct {
// 	ID   uint32 `gorm:"primary_key;auto_increment"`
// 	Name string `gorm:"size:255;not null;unique"`
// }

// type City struct {
// 	ID        uint32   `gorm:"primary_key;auto_increment"`
// 	Name      string   `gorm:"size:255;not null;unique"`
// 	CountryID uint32   `gorm:"size:255;not null;index"`
// 	Country   Location `gorm:"foreignKey:CountryID;references:ID"`
// }

// rest
type Client struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey"`
	Name    string `json:"name" binding:"required"`
	Phone   string `json:"phone" binding:"required" gorm:"not null;index:idx_phone,unique"`
	Details string `json:"details" binding:"required"`
	// Offers  []Offers `gorm:"foreignKey:UserRefer"`
}

type User struct {
	gorm.Model
	ID      uint   `gorm:"primaryKey"`
	Name    string `json:"name" binding:"required"`
	Phone   string `json:"phone" binding:"required" gorm:"index:idx_phone,unique"`
	Details string `json:"details" binding:"required"`
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
	CardName     Card   `json:"card_name"`
	DiscountRate uint   `json:"discount_rate"`
	ClientId     uint   `gorm:"size:255;not null;index"`
	Client       Client `gorm:"foreignKey:ClientId;references:ID"`
	// Country      Location `gorm:"foreignKey:CountryID;references:ID"`
	// UserRefer uint
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
	// dsn := "test:test@tcp(127.0.0.1:3306)/nimbus?charset=utf8mb4&parseTime=True&loc=Local"
	// Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panicln("failed to establish error db connection")
	}
	log.Println("db conn")

	Db.AutoMigrate(&Client{})
	Db.AutoMigrate(&User{})
	Db.AutoMigrate(&Offers{})
	// Db.AutoMigrate(&Transaction{})
	Db.AutoMigrate(&KeyPair{})
}
