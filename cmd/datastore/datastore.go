package datastore

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


type Client struct {
	gorm.model
	ID         uint   `gorm:"primaryKey"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Details    string `json:"details"`
}

type Card string

const (
	Basic Card = iota
	Premium
	Platinum
)

type Offers struct {
	gorm.model
	ID            uint   `gorm:"primaryKey"`
	CardName      Card `json:"name"`
	DiscountRate  uint `json:"phone"`
	Client        Client `json:"client"`
}



