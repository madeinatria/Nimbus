package model

import "github.com/madeinatria/Nimbus/cmd/datastore"

type SignUpPayload struct {
	datastore.Client
	Type string `json:"type" binding:"required"`
}

type InitLoginPayload struct {
	Phone string `json:"phone" binding:"required"`
}

type LoginPayload struct {
	Phone string `json:"phone" binding:"required"`
	Nonce string `json:"nonce" binding:"required"`
	OTP   string `json:"otp" binding:"required"`
}

type InitOfferPayload struct {
	Phone string `json:"phone" binding:"required"`
}

type ReedemOfferPayload struct {
	UserId string `json:"user_id" binding:"required"`
	Amount string `json:"amount" binding:"required"`
	OTP    string `json:"otp" binding:"required"`
}
