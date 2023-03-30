package datastore

import (
	"gorm.io/gorm"
)

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
