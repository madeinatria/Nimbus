package utils

import (
	"crypto/rand"
	"math/big"
	"strconv"

	"github.com/google/uuid"
)

func GenerateOTP() (string, error) {
	nBig, e := rand.Int(rand.Reader, big.NewInt(8999))
	if e != nil {
		return "failed to generate random int", e
	}
	return strconv.FormatInt(nBig.Int64()+1000, 10), nil
}

func GenerateNonce() (string, error) {
	u4, err := uuid.NewRandom()
	UUIDtoken := u4.String()
	if err != nil {
		return "failed to generate uuid", err
	}
	return UUIDtoken, nil
}
