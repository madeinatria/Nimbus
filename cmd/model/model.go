package model

import "github.com/madeinatria/Nimbus/cmd/datastore"

type SignUpPayload struct {
	datastore.Client
	Type string `json:"type"`
}

type InitLoginPayload struct {
	Phone string `json:"phone"`
}

type LoginPayload struct {
	Phone string `json:"phone"`
	UUID  string `json:"uuid"`
	OTP   string `json:"otp"`
}

type InitOfferPayload struct {
	User string `json:"user"`
}

type ReedemOfferPayload struct {
	UserId string `json:"user_id"`
	Amount string `json:"amount"`
	OTP    string `json:"otp"`
}
